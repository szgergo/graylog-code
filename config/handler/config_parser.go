package handler

import (
	"com/github/graylog-code/config/template"
	"com/github/graylog-code/error"
	"gopkg.in/yaml.v2"
)

func Parse(data []byte) template.GraylogConfigTemplate {
	var config template.GraylogConfigTemplate
	
	err := yaml.Unmarshal(data, &config)
	error.Check(err)
	checkConfiguration(config)
	applyDefaultValues(config)
	return config
}

func applyDefaultValues(config template.GraylogConfigTemplate) {
	if config.ApiPrefix == "" {
		config.ApiPrefix = "/api"
	}
}

func checkConfiguration(config template.GraylogConfigTemplate) {
	isInvalid := config.GraylogServer == ""
	if isInvalid {
		 error.ThrowError("graylog_server can't be null")
	 }
}


