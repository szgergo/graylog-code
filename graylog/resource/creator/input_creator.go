package creator

import (
	"com/github/graylog-code/config/template"
	"com/github/graylog-code/graylog/api"
	"com/github/graylog-code/graylog/client"
	log "github.com/sirupsen/logrus"
)

func CreateInputs(config template.GraylogConfigTemplate,inputConfigs []template.InputTemplate) {
	for i,inputConfig := range inputConfigs {
		log.Debugf("Creating %d. input, named: %s", i + 1, inputConfig.Title)
		var inputApi api.InputApi

		inputApi.Title = inputConfig.Title
		inputApi.Global = inputConfig.Global
		inputApi.Node = inputConfig.Node
		inputApi.Type = getInputTypeByName(config,inputConfig.Name)
		var configuration api.InputConfigurationApi
		inputApi.Configuration = &configuration

		configuration.BindAddress = inputConfig.Configuration.BindAddress
		configuration.DecompressSizeLimit = inputConfig.Configuration.DecompressSizeLimit
		configuration.NumberWorkerThreads = inputConfig.Configuration.NumberWorkerThreads
		configuration.OverrideSource = inputConfig.Configuration.OverrideSource
		configuration.Port = inputConfig.Configuration.Port
		configuration.RecvBufferSize = inputConfig.Configuration.RecvBufferSize

		var resource api.ResourceCreationResponse
		responseCode:= client.PostToGraylog(config,inputApi,"/system/inputs",&resource)
		log.Debugf("Resource with id: %s created: %d",resource.Id,responseCode)

		if inputConfig.AutoStart {
			log.Debugf("Starting input: %s", inputConfig.Title)
			client.PutToGraylog(config,nil,"/system/inputstates/" + resource.Id)
		}
	}
}

func getInputTypeByName(config template.GraylogConfigTemplate, inputName string) string {
	inputTypes := getInputTypes(config)
	for key, element := range inputTypes.Types {
		if element == inputName {
			return key
		}
	}
	return ""
}

func getInputTypes(config template.GraylogConfigTemplate) api.InputTypesApi {
	inputTypesEndpoint := "/system/inputs/types"
	var inputTypes api.InputTypesApi
	client.GetFromGraylog(config,inputTypesEndpoint,&inputTypes)
	return inputTypes
}
