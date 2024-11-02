package infra

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Environment    string
	ServerPort     string
	DBHost         string
	DBPort         string
	DBUser         string
	DBPassword     string
	DBDatabase     string
	DBSSLMode      string
	DBDriver       string
	QueueServer    string
	QueueGroup     string
	QueueTopic     string
	SignatureToken string
	DBDatabaseSP   string
	AWSRegion      string
}

func NewConfig() Config {
	if os.Getenv("ENVIRONMENT") == "" {
		if err := godotenv.Load(".env"); err != nil {
			panic("Error loading env file")
		}
	}

	return Config{
		Environment:    os.Getenv("ENVIRONMENT"),
		ServerPort:     os.Getenv("SERVER_PORT"),
		DBHost:         os.Getenv("DB_HOST"),
		DBPort:         os.Getenv("DB_PORT"),
		DBUser:         os.Getenv("DB_USER"),
		DBPassword:     os.Getenv("DB_PASSWORD"),
		DBDatabase:     os.Getenv("DB_DATABASE"),
		DBSSLMode:      os.Getenv("DB_SSL_MODE"),
		DBDriver:       os.Getenv("DB_DRIVER"),
		QueueServer:    os.Getenv("QUEUE_SERVER"),
		QueueGroup:     os.Getenv("QUEUE_GROUP"),
		QueueTopic:     os.Getenv("QUEUE_TOPIC"),
		SignatureToken: os.Getenv("SIGNATURE_STRING"),
		DBDatabaseSP:   os.Getenv("DB_DATABASE_SP"),
		AWSRegion:      os.Getenv("AWS_REGION"),
	}
}
