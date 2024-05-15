package pg_adapter

import "os"

type PostgreSQLConfig struct {
	Url string
	Port string
	User string
	Pwd string
}

var WithEnvConfig = PostgreSQLConfig{
	Url: os.Getenv("POSTGRES_URL"),
	Port: os.Getenv("POSTGRES_PORT"),
	User: os.Getenv("POSTGRES_USER"),
	Pwd: os.Getenv("POSTGRES_PWD"),
}