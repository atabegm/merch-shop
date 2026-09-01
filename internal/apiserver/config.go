package apiserver

// Config object create.

type Config struct {
	BindAddr    string `yaml:"bind_addr"`
	LogLevel    string `yaml:"log_level"`
	DatabaseURL string `yaml:"database_url"`
}

type DbConfig struct {
	PgUser     string `env:"PGUSER"`
	PgPassword string `env:"PGPASSWORD"`
	PgHost     string `env:"PGHOST"`
	PgPort     uint16 `env:"PGPORT"`
	PgDatabase string `env:"PGDATABASEURL"`
	PgSSLMode  string `env:"PGSSLMODE"`
}

func NewDBConfig() *DbConfig {
	dbCfg := &DbConfig{
		PgUser:     "postgres",
		PgPassword: "password",
		PgHost:     "localhost",
		PgPort:     5433,
		PgDatabase: "shop",
		PgSSLMode:  "disable",
	}

	return dbCfg
}

// NewConfig constructor create.
func NewConfig() *Config {
	cfg := &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}

	return cfg
}
