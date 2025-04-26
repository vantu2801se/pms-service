package config

import (
	"os"
	"time"

	"github.com/BurntSushi/toml"
)

type Config struct {
	AppName   string       `toml:"app_name"`
	Env       string       `toml:"env"`
	Port      string       `toml:"server_port"`
	LogLevel  string       `toml:"log_level"`
	LogFolder string       `toml:"log_folder"`
	Version   string       `toml:"version"`
	JWTSecret string       `toml:"jwt_secret_key"`
	Redis     *RedisConfig `toml:"redis"`
	RDS       *RDSConfig   `toml:"rds"`
}

func NewConfig(filePath string) (*Config, error) {
	var cfg Config

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	_, err = toml.NewDecoder(f).Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

type RedisConfig struct {
	Endpoint string `toml:"endpoint"`
	Port     int    `toml:"port"`
	Password string `toml:"password"`
}

type RDSConfig struct {
	DSN         string        `toml:"dsn"`
	MaxIdle     int           `toml:"max_idle"`
	MaxConn     int           `toml:"max_conn"`
	MaxLifetime time.Duration `toml:"max_lifetime"`
}
