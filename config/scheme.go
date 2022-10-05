package config

import "github.com/spf13/viper"

type Scheme struct {
	Port int
}

func init() {
	// environment - could be "local", "prod", "dev"
	viper.SetDefault("port", 9092)
}
