package helpers

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

const configPath = "config.yml"

type Cfg struct {
	DB          string `yaml:"db"`
}

var AppConfig Cfg

func ReadConfig() {
	f, err := os.Open(configPath)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&AppConfig)

	if err != nil {
		fmt.Println(err)
	}
}

