package main

import (
	"io"
	"os"

	"gopkg.in/yaml.v2"
)

type Search struct {
	Area     string `yaml:"area"`
	Category string `yaml:"category"`
	Term     string `yaml:"term"`
}

type Config struct {
	Searches []Search `yaml:"searches"`
	APIURL   string   `yaml:"api_url"`
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
