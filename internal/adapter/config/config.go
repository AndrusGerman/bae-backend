package config

import (
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

type (
	Container struct {
		fx.Out

		App  *App
		DB   *DB
		HTTP *HTTP
	}
	// App contains all the environment variables for the application
	App struct {
		Name string
		Env  string
	}

	// Database contains all the environment variables for the database
	DB struct {
		Connection string
		Host       string
		Port       string
		User       string
		Password   string
		Name       string
	}
	// HTTP contains all the environment variables for the http server
	HTTP struct {
		Port string
		URL  string
	}
)

func New() (Container, error) {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			return Container{}, err
		}
	}

	app := &App{
		Name: os.Getenv("APP_NAME"),
		Env:  os.Getenv("APP_ENV"),
	}

	db := &DB{
		Connection: os.Getenv("DB_CONNECTION"),
		Host:       os.Getenv("DB_HOST"),
		Port:       os.Getenv("DB_PORT"),
		User:       os.Getenv("DB_USER"),
		Password:   os.Getenv("DB_PASSWORD"),
		Name:       os.Getenv("DB_NAME"),
	}

	http := &HTTP{
		Port: os.Getenv("HTTP_PORT"),
		URL:  os.Getenv("HTTP_URL"),
	}

	return Container{
		App:  app,
		DB:   db,
		HTTP: http,
	}, nil
}
