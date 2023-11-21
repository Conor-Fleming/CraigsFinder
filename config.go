package main

import (
	"io"
	"os"

	"github.com/Conor-Fleming/CraigsFinder/craigslist"
	"gopkg.in/yaml.v2"
)

func loadConfig(filename string) (craigslist.Config, error) {
	var cfg craigslist.Config
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
