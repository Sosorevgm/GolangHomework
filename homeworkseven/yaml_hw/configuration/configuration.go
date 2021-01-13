package configuration

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

type YamlConfig struct {
	Port        int32  `yaml:"port"`
	DbUrl       string `yaml:"db_url"`
	JaegerUrl   string `yaml:"jaeger_url"`
	SentryUrl   string `yaml:"sentry_url"`
	KafkaBroker string `yaml:"kafka_broker"`
	AppId       string `yaml:"some_app_id"`
	AppKey      string `yaml:"some_app_key"`
}

const urlPrefix = "http://"

func ReadYamlConfig(fileName string) (error, *YamlConfig) {
	yamlConfig := YamlConfig{}

	fileConfig, fileReadErr := ioutil.ReadFile(fileName)
	if fileReadErr != nil {
		return fileReadErr, nil
	}

	unmarshalErr := yaml.Unmarshal(fileConfig, &yamlConfig)
	if unmarshalErr != nil {
		return fileReadErr, nil
	}

	urlErr := validateUrl([]string{yamlConfig.JaegerUrl, yamlConfig.SentryUrl})
	if urlErr != nil {
		return urlErr, nil
	}

	return nil, &yamlConfig
}

func validateUrl(urlSlice []string) error {
	for _, url := range urlSlice {
		if strings.HasPrefix(url, urlPrefix) != true {
			return errors.New(url + " has wrong url")
		}
	}
	return nil
}
