package manager

import (
	"com/github/graylog-code/config/template"
	resourceCreator "com/github/graylog-code/graylog/resource/creator"
	resourceDeleter "com/github/graylog-code/graylog/resource/deleter"
	resourceReader "com/github/graylog-code/graylog/resource/reader"
	log "github.com/sirupsen/logrus"
)

func ManageInputs(config template.GraylogConfigTemplate) {
	inputsFromGraylog := resourceReader.GetInputs(config)
	log.Debug("Inputs from graylog: ",inputsFromGraylog)
	inputsFromConfig := config.Inputs
	log.Debug("Inputs from config: ", inputsFromConfig)

	needsToBeDeleted := template.InputComplement(inputsFromGraylog, inputsFromConfig)
	needsToBeAdded := template.InputComplement(inputsFromConfig, inputsFromGraylog)

	if len(needsToBeDeleted) == 0 && len(needsToBeAdded) == 0 {
		log.Info("All inputs are up to date.")
	} else {
		resourceCreator.CreateInputs(config, needsToBeAdded)
		resourceDeleter.DeleteInputs(config, needsToBeDeleted)
	}
}
