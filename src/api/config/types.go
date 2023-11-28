package config

import (
	"github.com/spf13/viper"
)

type Configuration struct {
	Mongo MongoConfiguration `yaml:"mongo"`
}

type MongoConfiguration struct {
	ConnectionString   string `yaml:"connectionString"`
	Database           string `yaml:"database"`
	CompanyCollection  string `yaml:"companyCollection"`
	LocationCollection string `yaml:"locationCollection"`
	UserName           string `yaml:"userName"`
	Password           string `yaml:"password"`
}

func ProvideConfiguration() (*Configuration, error) {
	config := &Configuration{}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") // Path to the directory where the config file is located

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
