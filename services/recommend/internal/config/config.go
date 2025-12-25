package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type HTTPConfig struct {
	Addr string `yaml:"addr"`
}

type RedisConfig struct {
	Addr string `yaml:"addr"`
}

type RatingConfig struct {
	BaseURL string `yaml:"base_url"`
}

type Config struct {
	HTTP   HTTPConfig   `yaml:"http"`
	Redis  RedisConfig  `yaml:"redis"`
	Rating RatingConfig `yaml:"rating"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
