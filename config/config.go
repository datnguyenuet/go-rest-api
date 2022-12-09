package config

import (
	"github.com/spf13/viper"
	"log"
	"time"
)

// Config struct
type Config struct {
	Logger  Logger
	Server  ServerConfig
	Session Session
	Cookie  Cookie
}

// Logger struct
type Logger struct {
	Level    string
	Encoding string
}

type ServerConfig struct {
	Mode         string
	Port         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	PprofPort    string
}

type Session struct {
	Prefix string
	Name   string
	Expire int
}

type Cookie struct {
	Name     string
	MaxAge   int
	Secure   bool
	HTTPOnly bool
}

func LoadConfig(configEnv string) (*viper.Viper, error) {
	v := viper.New()

	filePath := "./config/config"

	if configEnv == "LOCAL" {
		filePath = "./config/config"
	}

	v.SetConfigName(filePath)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("config file not found")
		}
		return nil, err
	}

	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("Unable to decode into struct. Error: %v", err)
		return nil, err
	}

	return &c, nil
}
