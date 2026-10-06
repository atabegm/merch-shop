package apiserver

import (
	"os"
	"strconv"
)

// Config object create.

// Config create.
type Config struct {
	BindAddr     string `yaml:"bind_addr"`
	LogLevel     string `yaml:"log_level"`
	KafkaTopic   string `yaml:"kafka_topic"`
	KafkaGroupID string `yaml:"kafka_group_id"`

	KafkaBroker string

	SMTPHost string
	SMTPPort string
	SMTPFrom string
}

// DBConfig create.
type DBConfig struct {
	PgUser     string `env:"PGUSER"`
	PgPassword string `env:"PGPASSWORD"`
	PgHost     string `env:"PGHOST"`
	PgPort     uint16 `env:"PGPORT"`
	PgDatabase string `env:"PGDATABASEURL"`
	PgSSLMode  string `env:"PGSSLMODE"`
}

// NewDBConfig create.
func NewDBConfig() *DBConfig {
	port, err := strconv.ParseUint(
		os.Getenv("PGPORT"),
		10,
		16,
	)
	if err != nil {
		port = 5433
	}

	return &DBConfig{
		PgUser:     os.Getenv("PGUSER"),
		PgPassword: os.Getenv("PGPASSWORD"),
		PgHost:     os.Getenv("PGHOST"),
		PgPort:     uint16(port),
		PgDatabase: os.Getenv("PGDATABASE"),
		PgSSLMode:  os.Getenv("PGSSLMODE"),
	}
}

// NewConfig constructor create.
func NewConfig() *Config {
	return &Config{
		KafkaBroker: os.Getenv("KAFKA_BROKER"),
		SMTPHost:    os.Getenv("SMTP_HOST"),
		SMTPPort:    os.Getenv("SMTP_PORT"),
		SMTPFrom:    os.Getenv("SMTP_FROM"),
	}
}
