package main

import (
	"log"
	"yamlMod/yaml_hw/configuration"
)

func main() {

	yamlConfig, err := configuration.ReadYamlConfig("conf.yaml")
	if err != nil {
		log.Fatal(err)
	}

	log.Print(yamlConfig)
}
