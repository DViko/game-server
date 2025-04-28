package helpers

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Route struct {
	Name string `yaml:"name"`
	URL  string `yaml:"url"`
}

type GRouts struct {
	GRoutes []Route `yaml:"GRoutes"`
}

func ReadConfig(filePath string) *GRouts {

	file, _ := os.Open(filePath)

	defer file.Close()

	var config GRouts

	yaml.NewDecoder(file).Decode(&config)

	return &config
}
