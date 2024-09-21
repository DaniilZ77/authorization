package config

import (
	"flag"

	"github.com/DaniilZ77/authorization/internal/lib/logger"
)

type (
	Config struct {
		HTTP
		Log
		DB
	}

	HTTP struct {
		Port string
	}

	Log struct {
		Level string
	}

	DB struct {
		URL string
	}
)

func NewConfig() (*Config, error) {
	port := flag.String("port", "8080", "HTTP port")
	logLevel := flag.String("log_level", string(logger.InfoLevel), "logger level")
	dbURL := flag.String("db_url", "", "url for connection to database")

	flag.Parse()

	cfg := &Config{
		HTTP: HTTP{
			Port: *port,
		},
		Log: Log{
			Level: *logLevel,
		},
		DB: DB{
			URL: *dbURL,
		},
	}

	return cfg, nil
}
