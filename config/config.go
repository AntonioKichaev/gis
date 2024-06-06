package config

import (
	"fmt"
	"os"
)

type Config struct {
	host string
	port string
}

func NewConfig() *Config {
	host := os.Getenv("HOST")

	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	return &Config{
		host: host,
		port: port,
	}
}

func (c *Config) GetAddr() string {
	return fmt.Sprintf("%s:%s", c.host, c.port)
}
