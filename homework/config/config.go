package config

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Network string `yaml:"network"`
	Address string `yaml:"address"`
}

func Load(path string) (*Config, error) {
	configInfo, err := os.ReadFile(path)
	if err != nil {
		return nil, errors.New("error when trying to read a httpConfig file or load data from this file" + err.Error())
	}

	var cfg Config
	err = yaml.Unmarshal(configInfo, &cfg)
	if err != nil {
		return nil, errors.New("Failed to unmarshal date from httpConfig file: " + err.Error())
	}

	return &cfg, nil

}
