package reader

import (
	"com/github/graylog-code/config/template"
	"com/github/graylog-code/graylog/client"
	log "github.com/sirupsen/logrus"
)

type inputReadApi struct {
	Inputs *[]input `json:"inputs"`
}

type inputAttributes struct {
	RecvBufferSize int `json:"recv_buffer_size"`
	Port int `json:"port"`
	NumberWorkerThreads int `json:"number_worker_threads"`
	OverrideSource string `json:"override_source,omitempty"`
	BindAddress string `json:"bind_address"`
	DecompressSizeLimit int `json:"decompress_size_limit"`
}

type input struct {
	Id string                   `json:"id"`
	User string `json:"creator_user_id,omitempty"`
	Title string                `json:"title"`
	Name  string                `json:"name"`
	Global bool                 `json:"global"`
	Type string                 `json:"type"`
	Attributes *inputAttributes `json:"attributes"`
	Node string                 `json:"node"`
}


func GetInputs(config template.GraylogConfigTemplate) []template.InputTemplate {
	var inputConfigs []template.InputTemplate

	var inputReadApi inputReadApi

	client.GetFromGraylog(config,"/system/inputs",&inputReadApi)

	for i, input := range *inputReadApi.Inputs {
		if input.User != config.User {
			log.Debugf("Input %s's owner is not %s, skipping",input.Title,config.User)
		}
		log.Debugf("Checking %d. input from Graylog", i+1)
		var inputConfig template.InputTemplate
		inputConfig.Global = input.Global
		inputConfig.Title = input.Title
		inputConfig.Name = input.Name
		inputConfig.Node = input.Node
		inputConfig.Id = input.Id
		var inputConfiguration = template.InputConfigurationTemplate{}
		inputConfig.Configuration = &inputConfiguration

		inputConfiguration.Port = input.Attributes.Port
		inputConfiguration.RecvBufferSize = input.Attributes.RecvBufferSize
		inputConfiguration.NumberWorkerThreads = input.Attributes.NumberWorkerThreads
		inputConfiguration.OverrideSource = input.Attributes.OverrideSource
		inputConfiguration.DecompressSizeLimit = input.Attributes.DecompressSizeLimit
		inputConfiguration.BindAddress = input.Attributes.BindAddress

		inputConfigs = append(inputConfigs,inputConfig)
	}

	return inputConfigs
}
