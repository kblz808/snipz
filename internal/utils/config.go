package utils

import (
	"os"

	"github.com/joho/godotenv"
)

type DB struct {
	Connection string
	Host       string
	Port       string
	User       string
	Password   string
	Name       string
}

type Container struct {
	DB *DB
}

func New() (*Container, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	db := &DB{
		Connection: os.Getenv("DB_CONNECTION"),
		Host:       os.Getenv("DB_HOST"),
		Port:       os.Getenv("DB_PORT"),
		User:       os.Getenv("DB_USER"),
		Password:   os.Getenv("DB_PASSWORD"),
		Name:       os.Getenv("DB_NAME"),
	}

	return &Container{db}, nil

}
