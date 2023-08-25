package config

import (
	"fmt"
	"log"
	"os"

	configs "codeup.aliyun.com/6145b2b428003bdc3daa97c8/go-simba/go-simba-pkg.git/config"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		configs.App  `yaml:"app"`
		configs.HTTP `yaml:"http"`
		configs.Log  `yaml:"logger"`
		DataSource   `yaml:"datasource"`
	}
	DataSource struct {
		Type  string `env-required:"true" yaml:"type" env:"TYPE"`
		Mysql Mysql  `env-required:"true" yaml:"mysql" env:"MYSQL"`
		PG    PG     `env-required:"true" yaml:"postgres" env:"POSTGRES"`
	}
	PG struct {
		PoolMax int                  `env-required:"true" yaml:"pool_max" env:"PG_POOL_MAX"`
		DsnURL  configs.DBConnString `env-required:"true" yaml:"dsn_url" env:"PG_DSN_URL"`
	}

	Mysql struct {
		MaxOpenConns int                  `env-required:"true" yaml:"max_open_conns" env:"MAX_OPEN_CONNS"`
		MaxIdleConns int                  `env-required:"true" yaml:"max_idle_conns" env:"MAX_IDLE_CONNS"`
		URL          configs.DBConnString `env-required:"true" yaml:"url" env:"URL"`
	}
)

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
