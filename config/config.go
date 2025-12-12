package config

import (
    "os"
    "gopkg.in/yaml.v3"
)

type PostgresConfig struct {
    DSN string `yaml:"dsn"`
}

type RedisConfig struct {
    Addr string `yaml:"addr"`
}

type KafkaConfig struct {
    Brokers  []string `yaml:"brokers"`
    GroupID  string   `yaml:"group_id"`
    Topic    string   `yaml:"topic"`
}

type HTTPConfig struct {
    Addr string `yaml:"addr"`
}

type Config struct {
    Postgres   PostgresConfig `yaml:"postgres"`
    Redis      RedisConfig    `yaml:"redis"`
    Kafka      KafkaConfig    `yaml:"kafka"`
    HTTP       HTTPConfig     `yaml:"http"`
    Migrations struct {
        Path string `yaml:"path"`
    } `yaml:"migrations"`
}

func Load(path string) (*Config, error) {
    b, err := os.ReadFile(path)
    if err != nil {
        return nil, err
    }
    var cfg Config
    if err := yaml.Unmarshal(b, &cfg); err != nil {
        return nil, err
    }
    return &cfg, nil
}
