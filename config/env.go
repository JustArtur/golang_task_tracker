package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	JWTSecret string
	DbName    string
	DbHost    string
	DbPort    string
	DbUser    string
	DbPass    string
	DbSSLMode string
}

var Envs Config

func InitEnvs() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev"
	}

	err := godotenv.Load("../.env." + env)

	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("env: load env success")
	}

	Envs = loadEnvs()
}

func loadEnvs() Config {
	return Config{
		JWTSecret: os.Getenv("JWT_SECRET"),
		DbName:    os.Getenv("DB_NAME"),
		DbHost:    os.Getenv("DB_HOST"),
		DbPort:    os.Getenv("DB_PORT"),
		DbUser:    os.Getenv("DB_USER"),
		DbPass:    os.Getenv("DB_PASS"),
		DbSSLMode: os.Getenv("DB_SSL_MODE"),
	}
}
