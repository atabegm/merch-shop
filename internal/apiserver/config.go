package apiserver

import (
	"os"
	"strconv"
)

type Config struct {
	BindAddr string `yaml:"bind_addr"`
	LogLevel string `yaml:"log_level"`

	KafkaBroker  string `yaml:"-"`
	KafkaTopic   string `yaml:"kafka_topic"`
	KafkaGroupID string `yaml:"kafka_group_id"`

	SMTPHost string `yaml:"-"`
	SMTPPort string `yaml:"-"`
	SMTPFrom string `yaml:"-"`
}

type DBConfig struct {
	PgUser     string 
	PgPassword string
	PgHost     string
	PgPort     uint16
	PgDatabase string
	PgSSLMode  string
}

func NewConfig() *Config {
	return &Config{
		KafkaBroker: os.Getenv("KAFKA_BROKERS"),
		SMTPHost:    os.Getenv("SMTP_HOST"),
		SMTPPort:    os.Getenv("SMTP_PORT"),
		SMTPFrom:    os.Getenv("SMTP_FROM"),
	}
}

func NewDBConfig() *DBConfig {
	port, err := strconv.ParseUint(
		os.Getenv("DATABASE_PORT"),
		10,
		16,
	)
	if err != nil {
		port = 5432
	}

	return &DBConfig{
		PgUser:     os.Getenv("DATABASE_USER"),
		PgPassword: os.Getenv("DATABASE_PASSWORD"),
		PgHost:     os.Getenv("DATABASE_HOST"),
		PgPort:     uint16(port),
		PgDatabase: os.Getenv("DATABASE_NAME"),
		PgSSLMode:  os.Getenv("DATABASE_SSLMODE"),
	}
}
