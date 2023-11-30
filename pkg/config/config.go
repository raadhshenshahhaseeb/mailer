package configuration

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config interface {
	GetConfig() *config
}

func (c *config) GetConfig() *config {
	return c
}

type config struct {
	Logger struct {
		Level int    `mapstructure:"LEVEL" json:"level"`
		Env   string `mapstructure:"ENV" json:"env"`
	} `mapstructure:"LOGGER" json:"logger"`
	Server struct {
		Host    string `mapstructure:"HOST" json:"host"`
		PORT    string `mapstructure:"PORT" json:"port"`
		CORSAGE int    `mapstructure:"CORS_AGE" json:"cors"`
	} `mapstructure:"SERVER" json:"server"`
}

func Init() (Config, error) {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.BindEnv("LOGGER.LEVEL")
	viper.BindEnv("LOGGER.ENV")

	viper.BindEnv("SERVER.HOST")
	viper.BindEnv("SERVER.PORT")
	viper.BindEnv("SERVER.CORS_AGE")

	conf := config{}

	err := viper.Unmarshal(&conf)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling config: %w", err)
	}

	return &conf, nil
}
