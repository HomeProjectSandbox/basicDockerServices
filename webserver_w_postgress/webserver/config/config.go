package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type AppConfig struct {
	App struct {
		DbName string `mapstructure:"dbName"`
		DbPort int    `mapstructure:"dbPort"`
		DbUser string `mapstructure:"dbUser"`
		DbPw   string `mapstructure:"dbPw"`
	} `mapstructure:"app"`
}

func GetConfig(configPath string, configName string) *AppConfig {
	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configPath)

	viper.AutomaticEnv()
	viper.SetEnvPrefix("env")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // this is useful e.g. want to use . in Get() calls, but environmental variables to use _ delimiters (e.g. app.port -> APP_PORT)

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	var config AppConfig
	// Unmarshal the config file into the AppConfig struct
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	return &config

}
