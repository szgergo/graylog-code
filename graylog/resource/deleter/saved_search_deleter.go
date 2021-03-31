package deleter

import (
	"com/github/graylog-code/config/template"
	"com/github/graylog-code/graylog/client"
	log "github.com/sirupsen/logrus"
)

func DeleteSavedSearches(config template.GraylogConfigTemplate, search []template.SavedSearchTemplate) {
	for i, savedSearch := range search {
		if savedSearch.Id == "" {
			log.Error("Can't delete saved search, no Id was found")
			continue
		}
		log.Debugf("Deleting %d. saved search with id: %s",i + 1, savedSearch.Id)
		client.DeleteFromGraylog(config,"/views/" + savedSearch.Id)

	}
}
