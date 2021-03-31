package manager

import (
	"com/github/graylog-code/config/template"
	resourceCreator "com/github/graylog-code/graylog/resource/creator"
	resourceDeleter "com/github/graylog-code/graylog/resource/deleter"
	resourceReader "com/github/graylog-code/graylog/resource/reader"
	log "github.com/sirupsen/logrus"
)

func ManageSavedSearches(config template.GraylogConfigTemplate) {
	savedSearchesFromGraylog := resourceReader.GetSavedSearches(config)
	log.Debug(savedSearchesFromGraylog)
	savedSearchesFromConfig := config.SavedSearches
	log.Debug(savedSearchesFromConfig)

	needsToBeDeleted := template.SavedSearchComplement(savedSearchesFromGraylog, savedSearchesFromConfig)
	needsToBeAdded := template.SavedSearchComplement(savedSearchesFromConfig, savedSearchesFromGraylog)

	if len(needsToBeDeleted) == 0 && len(needsToBeAdded) == 0 {
		log.Info("All saved searches are up to date.")
	} else {
		resourceCreator.CreateSavedSearches(config, needsToBeAdded)
		resourceDeleter.DeleteSavedSearches(config, needsToBeDeleted)
	}
}