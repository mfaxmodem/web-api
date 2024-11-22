package config

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server         ServerConfig
	Postgres       PostgresConfig
	Redis          RedisConfig
	PasswordConfig PasswordConfig
}

type ServerConfig struct {
	Port    string
	RunMode string
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	SSlMode  string
}

type RedisConfig struct {
	Host               string
	Port               string
	Password           string
	DB                 string
	DialTimeout        time.Duration
	ReadTimeout        time.Duration
	WriteTimeout       time.Duration
	IdleCheckFrequency time.Duration
	PoolSize           int
	PoolTimeout        time.Duration
}
type PasswordConfig struct {
	IncludeChars     bool
	IncludeDigits    bool
	MinLength        int
	MaxLength        int
	IncludeUppercase bool
	IncludeLowercase bool
}

func GetConfig() (*Config, error) {
	cfgPath := getConfigPath(os.Getenv("APP_ENV"))
	v, err := LoadConfig(cfgPath, "yaml")
	if err != nil {
		return nil, fmt.Errorf("error loading config: %w", err)
	}

	cfg, err := ParseConfig(v)
	if err != nil {
		return nil, fmt.Errorf("error parsing config: %w", err)
	}
	return cfg, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}
	return &config, nil
}

func LoadConfig(fileName string, fileType string) (*viper.Viper, error) {
	config := viper.New()
	config.SetConfigType(fileType)
	config.SetConfigName(fileName)
	config.AddConfigPath(".")
	config.AutomaticEnv()

	config.SetDefault("Server.Port", "5005")
	config.SetDefault("Redis.PoolSize", 10)

	if err := config.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}
	return config, nil
}

func getConfigPath(env string) string {
	switch env {
	case "docker":
		return "/config/config-docker"
	case "production":
		return "/config/config-production"
	default:
		return "../config/config-development"
	}
}
