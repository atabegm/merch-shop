package main

import (
	"avito/internal/apiserver"
	"context"
	"log"
	"os"

	"github.com/goccy/go-yaml"
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

	err = apiserver.Start(context.TODO(), dbCfg, cfg)
	if err != nil {
		log.Println("failed to start", err)
	}
}
