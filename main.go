package main

import (
	"github.com/joho/godotenv"
	"go-rest-api/config"
	"go-rest-api/pkg/logger"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	cfgFile, err := config.LoadConfig(os.Getenv("CONFIG"))
	cfg, err := config.ParseConfig(cfgFile)
	if err != nil {
		log.Fatalf("ParseConfig: %v", err)
	}
	appLogger := logger.NewApiLogger(cfg)
	appLogger.ConfigLogger()

	appLogger.Info("run logging in first time")
}
