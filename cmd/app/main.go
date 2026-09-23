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
		log.Println("fail to read config file", err)
	}

	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		log.Println("fail to unmarshal", err)
	}

	logger := logrus.New()
	err = apiserver.Start(context.TODO(), cfg, dbCfg, logger)
	if err != nil {
		log.Println("failed to start", err)
	}
}
