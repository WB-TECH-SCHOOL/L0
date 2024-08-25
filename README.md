# Задание L0
## Зависимости
- Docker
- NATS CLI
## Установка и запуск
Клонирование репозитория:
```bash
$ git clone git@github.com:WB-TECH-SCHOOL/L0.git
```
Перейдите в папку `L0` и создайте там файл `.env`

Пример заполнения файла `.env` вы можете увидеть в файле `.env.example`

Запустите работу сервиса
```bash
$ docker compose up --build
```

После запуска вы можете взаимодействовать с сервисом через:
- API: `localhost:80/api/`
- SWAGGER: `localhost:80/swagger/index.html#/`
- NATS: `localhost:4222`
