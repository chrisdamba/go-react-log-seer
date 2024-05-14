package config

import (
	"log"
	"os"
)

type Config struct {
	PostgresHost     string
	PostgresUser     string
	PostgresPassword string
	PostgresDBName   string
	PostgresPort     string
	ElasticsearchURL string
}

var AppConfig *Config

func LoadConfig() {
	AppConfig = &Config{
		PostgresHost:     os.Getenv("POSTGRES_HOST"),
		PostgresUser:     os.Getenv("POSTGRES_USER"),
		PostgresPassword: os.Getenv("POSTGRES_PASSWORD"),
		PostgresDBName:   os.Getenv("POSTGRES_DBNAME"),
		PostgresPort:     os.Getenv("POSTGRES_PORT"),
		ElasticsearchURL: os.Getenv("ELASTICSEARCH_URL"),
	}

	missingEnvVars := []string{}
	if AppConfig.PostgresHost == "" {
		missingEnvVars = append(missingEnvVars, "POSTGRES_HOST")
	}
	if AppConfig.PostgresUser == "" {
		missingEnvVars = append(missingEnvVars, "POSTGRES_USER")
	}
	if AppConfig.PostgresPassword == "" {
		missingEnvVars = append(missingEnvVars, "POSTGRES_PASSWORD")
	}
	if AppConfig.PostgresDBName == "" {
		missingEnvVars = append(missingEnvVars, "POSTGRES_DBNAME")
	}
	if AppConfig.PostgresPort == "" {
		missingEnvVars = append(missingEnvVars, "POSTGRES_PORT")
	}
	if AppConfig.ElasticsearchURL == "" {
		missingEnvVars = append(missingEnvVars, "ELASTICSEARCH_URL")
	}

	if len(missingEnvVars) > 0 {
		log.Fatalf("Missing required environment variables: %v", missingEnvVars)
	}
}
