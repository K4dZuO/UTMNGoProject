package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type PostgresConfig struct {
	DSN string `yaml:"dsn"`
}

type KafkaConfig struct {
	Brokers []string `yaml:"brokers"`
	GroupID string   `yaml:"group_id"`
	Topic   string   `yaml:"topic"`
}

type RedisConfig struct {
	Addr string `yaml:"addr"`
}

type HTTPConfig struct {
	Addr string `yaml:"addr"`
}

type Config struct {
	Postgres PostgresConfig `yaml:"postgres"`
	Kafka    KafkaConfig    `yaml:"kafka"`
	Redis    RedisConfig    `yaml:"redis"`
	HTTP     HTTPConfig     `yaml:"http"`
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
