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
		RabbitMQ            `yaml:"rabbitmq" ,env-required:"false" ,env:"RABBITMQ"`
	}

	RabbitMQ struct {
		URL      string            `yaml:"url" ,env-required:"false" ,env:"RABBITMQ_URL"`
		Publish  []PublishProfile  `yaml:"publish" ,env-required:"false" ,env:"PUBLISH"`
		Consumer []ConsumerProfile `yaml:"consumer" ,env-required:"false" ,env:"CONSUMER"`
	}

	PublishProfile struct {
		Type string      `env-required:"true" yaml:"type" env:"TYPE"`
		Body PublishBody `env-required:"true" yaml:"body" env:"BODY"`
	}

	PublishBody struct {
		ExchangeName    string `env-required:"true" yaml:"exchangeName" env:"EXCHANGE_NAME"`
		BindingKey      string `env-required:"true" yaml:"bindingKey" env:"BINDING_KEY"`
		MessageTypeName string `env-required:"true" yaml:"messageTypeName" env:"MESSAGE_TYPE_NAME"`
	}
	ConsumerProfile struct {
		Type string       `env-required:"true" yaml:"type" env:"TYPE"`
		Body ConsumerBody `env-required:"true" yaml:"body" env:"BODY"`
	}

	ConsumerBody struct {
		ExchangeName string `env-required:"true" yaml:"exchangeName" env:"EXCHANGE_NAME"`
		BindingKey   string `env-required:"true" yaml:"bindingKey" env:"BINDING_KEY"`
		ConsumerTag  string `env-required:"true" yaml:"consumerTag" env:"CONSUMER_TAG"`
		QueueName    string `env-required:"true" yaml:"queueName" env:"QUEUE_NAME"`
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
