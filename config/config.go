package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	AppPort string `mapstructure:"APP_PORT"`

	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`

	JWTAccessSecret   string `mapstructure:"JWT_ACCESS_SECRET"`
	JWTRefreshSecret  string `mapstructure:"JWT_REFRESH_SECRET"`
	JWTAccessExpired  int64  `mapstructure:"JWT_ACCESS_EXPIRED"`
	JWTRefreshExpired int64  `mapstructure:"JWT_REFRESH_EXPIRED"`
}

var App Config

func Load() {

	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		panic("Failed load .env: " + err.Error())
	}

	viper.AutomaticEnv()
	if err := viper.Unmarshal(&App); err != nil {
		panic(err)
	}

}
