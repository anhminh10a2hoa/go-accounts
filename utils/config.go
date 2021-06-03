package utils

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBDRIVER   string `mapstructure:"DB_DRIVER"`
	DBADDRESS  string `mapstructure:"DB_ADDRESS"`
	DBPORT     string `mapstructure:"DB_PORT"`
	DBUSER     string `mapstructure:"DB_USER"`
	DBNAME     string `mapstructure:"DB_NAME"`
	DBPASSWORD string `mapstructure:"DB_PASSWORD"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
