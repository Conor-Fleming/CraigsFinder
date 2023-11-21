package main

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

func loadConfig(filename string) (Config, error) {
	var cfg Config
	yamlFile, err := os.Open(filename)
	if err != nil {
		return cfg, err
	}
	defer yamlFile.Close()

	yamlBytes, err := io.ReadAll(yamlFile)
	if err != nil {
		return cfg, err
	}

	err = yaml.Unmarshal(yamlBytes, &cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}
