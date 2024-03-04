package config

import "time"

type ServiceConfig struct {
	Service struct {
		JWT struct {
			DB struct {
				Mongo `mapstructure:"mongodb"`
			} `mapstructure:"db"`
			Network `mapstructure:"network"`
		} `mapstructure:"jwt"`
	} `mapstructure:"service"`
}

type Network struct {
	GRPC `mapstructure:"grpc"`
	HTTP `mapstructure:"http"`
}

type GRPC struct {
	Host      string    `mapstructure:"host"`
	Port      string    `mapstructure:"port"`
	Secret    string    `mapstructure:"secret"`
	Keepalive keepalive `mapstructure:"keepalive"`
}
type keepalive struct {
	KeepaliveParams `mapstructure:"KeepaliveParams"`
	KeepalivePolicy `mapstructure:"KeepalivePolicy"`
}
type KeepaliveParams struct {
	ServerParameters `mapstructure:"ServerParameters"`
}
type KeepalivePolicy struct {
	EnforcementPolicy `mapstructure:"EnforcementPolicy"`
}

type EnforcementPolicy struct {
	MinTime             time.Duration `mapstructure:"MinTime"`
	PermitWithoutStream bool          `mapstructure:"PermitWithoutStream"`
}
type ServerParameters struct {
	MaxConnectionIdle     time.Duration `mapstructure:"MaxConnectionIdle"`
	MaxConnectionAge      time.Duration `mapstructure:"MaxConnectionAge"`
	MaxConnectionAgeGrace time.Duration `mapstructure:"MaxConnectionAgeGrace"`
	Time                  time.Duration `mapstructure:"Time"`
	Timeout               time.Duration `mapstructure:"Timeout"`
}

type Mongo struct {
	DB          string   `mapstructure:"dbname"`
	Host        string   `mapstructure:"host"`
	Port        string   `mapstructure:"port"`
	URI         string   `mapstructure:"uri"`
	Username    string   `mapstructure:"username"`
	Password    string   `mapstructure:"password"`
	Collection  string   `mapstructure:"collection"`
	Collections []string `mapstructure:"collections"`
}

type HTTP struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}
