package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type AppConfig struct {
	Database DatabaseConfig `mapstructure:",squash"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"postgres_host"`
	Port     string `mapstructure:"postgres_port"`
	User     string `mapstructure:"postgres_user"`
	Database string `mapstructure:"postgres_db"`
	Password string `mapstructure:"postgres_password"`
}

func Init() *AppConfig {
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		viper.SetDefault(pair[0], "")
	}

	viper.AutomaticEnv() // read in environment variables that match

	var c AppConfig

	if err := viper.Unmarshal(&c); err != nil {
		fmt.Printf("unable to decode into struct, %v", err)
	}

	return &c
}
