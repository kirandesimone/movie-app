package internal

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

const configPath = "config.yaml"

type Cfg struct {
	Database    DatabaseCfg    `yaml:"database"`
	Application ApplicationCfg `yaml:"application"`
}

type DatabaseCfg struct {
	Uri  string `yaml:"uri"`
	Name string `yaml:"name"`
}

type ApplicationCfg struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func ReadConfig(cfg *Cfg) {
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(cfg)
	if err != nil {
		log.Fatal(err)
	}
}
