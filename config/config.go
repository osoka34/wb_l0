package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Postgres PostgresConfig `mapstructure:"postgres"`
	Logger   Logger         `mapstructure:"logger"`
	Deploy   Deploy         `mapstructure:"deploy"`
	Nats     NatsConfig     `mapstructure:"nats"`
}

type NatsConfig struct {
	Host    string `json:"host"`
	Port    string `json:"port"`
	Cluster string `json:"cluster"`
	Client  string `json:"client"`
	Topic   string `json:"topic"`
}

type Logger struct {
	Level             string `yaml:"level"`
	InFile            string `yaml:"inFile"`
	Development       bool   `yaml:"development"`
	DisableCaller     bool   `yaml:"disableCaller"`
	DisableStacktrace bool   `yaml:"disableStacktrace"`
	Encoding          string `yaml:"encoding"`
}

type PostgresConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"-"`
	DBName   string `json:"DBName"`
	PgDriver string `json:"pgDriver"`
}

type Deploy struct {
	IsTest bool   `json:"is_test"`
	Port   string `json:"port"`
}

type ServerConfig struct {
	AppVersion string `json:"appVersion"`
	Host       string `json:"host" validate:"required"`
	Port       string `json:"port" validate:"required"`
}

func LoadConfig() (*viper.Viper, error) {

	viperInstance := viper.New()

	viperInstance.AddConfigPath("./config")
	viperInstance.SetConfigName("config")
	viperInstance.SetConfigType("yaml")

	err := viperInstance.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return viperInstance, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {

	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Fatalf("unable to decode config into struct, %v", err)
		return nil, err
	}
	return &c, nil
}
