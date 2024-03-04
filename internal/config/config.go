package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func ReadConfig() ServiceConfig {

	viper.SetConfigName("config")
	viper.AddConfigPath("./internal/config/")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)

	}
	var cfg ServiceConfig
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal(err)
	}

	if cfg.Service.JWT.DB.Mongo.URI == "" {
		log.Println("empty uri")
		cfg.Service.JWT.DB.Mongo.URI = fmt.Sprintf("mongodb://%s:%s", cfg.Service.JWT.DB.Host, cfg.Service.JWT.DB.Port)
	}


	return cfg
}
