package config

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	TemplateDirectory string
}

type Template struct {
	Directory string
	Handler   string
}

func InitConfig() (*Config, error) {
	// Find home directory.
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	// Search config in home directory with name ".gotisan" (without extension).
	viper.AddConfigPath(home)
	viper.SetConfigName(".gotisan")
	viper.SetConfigType("yml")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func NewDefaultConfig() *Config {
	return &Config{
		TemplateDirectory: "./templates",
	}
}
