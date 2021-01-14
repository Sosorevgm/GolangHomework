package main

import (
	"go/homework/homeworkseven/yaml_hw/configuration"
	"log"
)

func main() {

	yamlConfig, err := configuration.ReadYamlConfig("/Users/gregorysosorev/GolandProjects/GolangHomework/homeworkseven/yaml_hw/conf.yaml")
	if err != nil {
		log.Fatal(err)
	}

	log.Print(yamlConfig)
}
