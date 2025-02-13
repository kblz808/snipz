package utils

import "os"

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
