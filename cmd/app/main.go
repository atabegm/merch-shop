package main

import (
	"avito/internal/apiserver"
	"context"
	"log"
	"os"

	"github.com/goccy/go-yaml"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := apiserver.NewConfig()
	dbCfg := apiserver.NewDBConfig()

	data, err := os.ReadFile("configs/config.yaml")
	if err != nil {
		log.Fatal("failed to read config file:", err)
	}

	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		log.Fatal("failed to unmarshal config:", err)
	}

	logger := logrus.New()

	err = apiserver.Start(
		context.Background(),
		cfg,
		dbCfg,
		logger,
	)
	if err != nil {
		log.Fatal("failed to start:", err)
	}
}
