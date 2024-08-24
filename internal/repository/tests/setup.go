package tests

import (
	"context"
	"errors"
	"fmt"
	"github.com/docker/go-connections/nat"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func setupPostgresContainer(ctx context.Context) (*sqlx.DB, func(), error) {
	req := testcontainers.ContainerRequest{
		Image:        "postgres:16",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_DB":       "test",
			"POSTGRES_USER":     "test",
			"POSTGRES_PASSWORD": "test123",
		},
		WaitingFor: wait.ForSQL("5432/tcp", "postgres", func(host string, port nat.Port) string {
			return fmt.Sprintf("postgres://test:test123@%s:%s/test?sslmode=disable", host, port.Port())
		}).WithStartupTimeout(60 * time.Second),
	}
	postgresC, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		panic(err)
	}

	// Функция для остановки контейнера после теста
	cleanup := func() {
		if err := postgresC.Terminate(ctx); err != nil {
			fmt.Printf("Could not terminate postgres container: %v", err)
		}
	}

	// Получаем порт, на котором запущен контейнер
	mappedPort, err := postgresC.MappedPort(ctx, "5432")
	if err != nil {
		cleanup()
		panic(err)
	}

	// Получаем хост, на котором запущен контейнер
	host, err := postgresC.Host(ctx)
	if err != nil {
		cleanup()
		panic(err)
	}

	// Формируем строку подключения к БД
	dsn := fmt.Sprintf("postgres://test:test123@%s:%s/test?sslmode=disable", host, mappedPort.Port())

	// Подключаемся к БД
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		cleanup()
		panic(err)
	}

	// Получаем текущий рабочий каталог
	wd, err := os.Getwd()
	if err != nil {
		cleanup()
		panic(err)
	}

	// Построение абсолютного пути к директории с миграциями
	migrationsPath := filepath.Join(wd, "../../../migrations")
	// Преобразование пути для Windows, заменяем обратные слеши на прямые
	migrationsPath = strings.ReplaceAll(migrationsPath, `\`, "/") // Преобразование для Windows
	// Не используем url.PathEscape для всего пути, так как это некорректно для file:// URL
	migrationsURL := fmt.Sprintf("file://%s", migrationsPath) // Формирование корректного URL

	// Применение миграций
	m, err := migrate.New(migrationsURL, dsn)
	if err != nil {
		cleanup()
		panic(err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		cleanup()
		panic(err)
	}

	return db, cleanup, nil
}
