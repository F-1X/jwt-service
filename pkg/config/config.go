package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type ServiceConfig struct {
	Service struct {
		JWT struct {
			DB struct {
				Mongo `mapstructure:"mongodb"`
			} `mapstructure:"db"`
		} `mapstructure:"jwt"`
	} `mapstructure:"service"`
}

type Mongo struct {
	DB         string `mapstructure:"dbname"`
	Host       string `mapstructure:"host"`
	Port       string `mapstructure:"port"`
	Collection string `mapstructure:"collection"`
	URI        string
}

func ReadConfig() ServiceConfig {

	viper.SetConfigName("mongo")
	viper.AddConfigPath("./pkg/config/")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s\n", err)

	}

	var cfg ServiceConfig
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Error unmarshaling config: %s\n", err)
	}


	cfg.Service.JWT.DB.Mongo.URI = fmt.Sprintf("mongodb://%s:%s", cfg.Service.JWT.DB.Host, cfg.Service.JWT.DB.Port)

	return cfg
}
