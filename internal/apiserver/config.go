package apiserver

// Config object create.

// Config create.
type Config struct {
	BindAddr    string `yaml:"bind_addr"`
	LogLevel    string `yaml:"log_level"`
	DatabaseURL string `yaml:"database_url"`
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
	dbCfg := &DBConfig{
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
