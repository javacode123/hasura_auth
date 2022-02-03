package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"time"
)

var GlobalConfig Config

type Config struct {
	AppName     string
	JWTSecret   string
	JWTExpireIn time.Duration
	HasuraHost  string
	Address     string
}

func InitConfig() {
	config := viper.New()
	config.AddConfigPath(".")
	err := config.ReadInConfig()
	if err != nil {
		panic(err)
	}

	GlobalConfig = Config{
		AppName:     config.Get("app-name").(string),
		JWTSecret:   config.Get("auth.jwt-secret").(string),
		JWTExpireIn: config.GetDuration("auth.jwt-expire-in"),
		Address:     config.Get("address").(string),
		HasuraHost:  config.Get("hasura.host").(string),
	}
	logrus.Infof("init config: %+v", GlobalConfig)
}
