package config

import (
	"fmt"
	"log"
	"os"

	configs "codeup.aliyun.com/6145b2b428003bdc3daa97c8/go-simba/go-simba-pkg.git/config"
	"github.com/Bifang-Bird/simbapkg/pkg/dbconfig"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		configs.App         `yaml:"app"`
		configs.HTTP        `yaml:"http"`
		configs.Log         `yaml:"logger"`
		dbconfig.DataSource `yaml:"datasource"`
		dbconfig.RabbitMQ   `yaml:"rabbitmq" ,env-required:"false" ,env:"RABBITMQ"`
	}
)

type MqConfig struct {
	ExchangeName    string
	BindingKey      string
	MessageTypeName string
}

var cfg *Config

func NewConfig() (*Config, error) {
	cfg = &Config{}
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	err = cleanenv.ReadConfig(dir+"/config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}
	err = cleanenv.ReadEnv(cfg)

	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func GetConfig() *Config {
	return cfg
}
