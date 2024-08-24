package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

const (
	DBName         = "DB_NAME"
	DBUser         = "DB_USER"
	DBPassword     = "DB_PASSWORD"
	DBPort         = "DB_PORT"
	DBHost         = "DB_HOST"
	DBResponseTime = "DB_RESPONSE_TIME"

	NATSPort = "NATS_PORT"
	NATSHost = "NATS_HOST"
)

func InitConfig() {
	envPath, _ := os.Getwd()
	envPath = filepath.Join(envPath, "..") // workdir is cmd

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(envPath)

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Failed to init config from file. Error:%v", err.Error())
	}
}
