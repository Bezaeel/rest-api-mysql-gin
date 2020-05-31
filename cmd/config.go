package main

import "fmt"

type Config struct {
	Env      string  `envconfig:"ENVIRONMENT" default:"development"`
	Port     int     `envconfig:"PORT" default:"8000"`
	Host     string	 `envconfig:"HOST" default:"localhost"`
	User     string  `envconfig:"USER" default:"root"`
	DSN      string  `envconfig:"DSN" default:"ContactManager"`
	Password string  `envconfig:"PASSWORD"`
}

func (cfg *Config) toURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DSN,
	)
}
