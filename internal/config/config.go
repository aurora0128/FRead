package config

import (
	"errors"
	"github.com/spf13/viper"
	"log"
)

var Cfg Config

type Config struct {
	//解析用
	Storage struct {
		MarkdownPath  string `mapstructure:"markdown_path"`
		MarkdownTPath string `mapstructure:"markdown_Testpath"`
	}
}

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			// Config file not found; ignore error if desired
			log.Fatal("No config file found")
		} else {
			log.Fatal(err)
		}
	}
	if err := viper.Unmarshal(&Cfg); err != nil {
		log.Fatal(err)
	}

}
