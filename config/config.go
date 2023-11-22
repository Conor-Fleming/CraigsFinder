package config

import (
	"io"
	"os"

	"github.com/Conor-Fleming/CraigsFinder/craigslist"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Searches []craigslist.Search `yaml:"searches"`
	APIURL   string              `yaml:"api_url"`
}

func LoadConfig(filename string) (Config, error) {
	//declaring empty Config object
	cfg := Config{}
	yamlFile, err := os.Open(filename)
	if err != nil {
		return cfg, err
	}
	defer yamlFile.Close()

	yamlBytes, err := io.ReadAll(yamlFile)
	if err != nil {
		return cfg, err
	}

	//unmarshal data from yaml file into our struct
	err = yaml.Unmarshal(yamlBytes, &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
