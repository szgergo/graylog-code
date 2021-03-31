package manager

import (
	"com/github/graylog-code/config/template"
	resourceCreator "com/github/graylog-code/graylog/resource/creator"
	resourceDeleter "com/github/graylog-code/graylog/resource/deleter"
	resourceReader "com/github/graylog-code/graylog/resource/reader"
	log "github.com/sirupsen/logrus"
)

func ManageStreams(config template.GraylogConfigTemplate) {
	streamsFromGraylog := resourceReader.GetStreams(config)
	log.Debug("Inputs from graylog: ", streamsFromGraylog)
	streamsFromConfig := config.Streams
	log.Debug("Inputs from config: ", streamsFromConfig)

	needsToBeDeleted := template.StreamComplement(streamsFromGraylog, streamsFromConfig)
	needsToBeAdded := template.StreamComplement(streamsFromConfig, streamsFromGraylog)

	if len(needsToBeDeleted) == 0 && len(needsToBeAdded) == 0 {
		log.Info("All streams are up to date.")
	} else {
		resourceCreator.CreateStreams(config, needsToBeAdded)
		resourceDeleter.DeleteStreams(config, needsToBeDeleted)
	}
}
