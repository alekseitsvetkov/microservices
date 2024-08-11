package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	HTTP struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"http"`

	GRPC struct {
		Client struct {
			Product struct {
				Host string `yaml:"host"`
				Port int    `yaml:"port"`
			} `yaml:"product"`
		} `yaml:"client"`
	} `yaml:"grpc"`
}

func New() (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadConfig("configs/config.yaml", &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
