package apiserver

// Config object create.
type Config struct {
	BindAddr    string `yaml:"bind_addr"`
	LogLevel    string `yaml:"log_level"`
	DatabaseURL string `yaml:"database_url"`
}

// New config constructor create.
func NewConfig() *Config {
	cfg := &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}

	return cfg
}
