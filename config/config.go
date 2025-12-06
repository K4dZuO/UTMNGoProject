package config


import (
"os"
"gopkg.in/yaml.v3"
)

type PostgresConfig struct {
    DSN string `yaml:"dsn"`
}



type Config struct {
	Postgres PostgresConfig `yaml:"postgres"`
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
