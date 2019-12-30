package config

import (
	"path"

	"github.com/mitchellh/go-homedir"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Initialize() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	home, hderr := homedir.Dir()
	if hderr == nil {

		p := path.Join(home, ".portal")

		viper.AddConfigPath(p)
	}

	err := viper.ReadInConfig()
	if err != nil {
		logrus.Errorf("error while reading in the configuration: %s", err)
	}
}

func GetConfig() (*Config, error) {
	c := &Config{}
	err := viper.Unmarshal(c)

	return c, err
}
