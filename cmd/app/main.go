package main

import (
	"avito/internal/apiserver"
	"flag"
	"log"
	"os"

	"github.com/goccy/go-yaml"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "configs/config.yaml", "path to config file")
}

func main() {
	flag.Parse()
	cfg := apiserver.NewConfig()
	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Println("fail to read config file", err)
	}

	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		log.Println("fail to unmarshal", err)
	}
}
