package creator

import (
	"com/github/graylog-code/config/template"
	"com/github/graylog-code/graylog/api"
	"com/github/graylog-code/graylog/client"
	log "github.com/sirupsen/logrus"
)

func CreateSavedSearches(config template.GraylogConfigTemplate, searches []template.SavedSearchTemplate) {
	for i, savedSearch := range searches {
		log.Debugf("Saving %d. saved searches: %s", i + 1, savedSearch.Name)

		search := createSearch(config,savedSearch)
		var createdSearchId api.ResourceCreationResponse
		searchSaveResult := client.PostToGraylog(config,search,"/views/searches",&createdSearchId)
		log.Debugf("Search save result: %d with id: %s", searchSaveResult,createdSearchId)


		view := CreateView(search,savedSearch)
		var createdViewId api.ResourceCreationResponse
		viewSaveResult := client.PostToGraylog(config,view,"/views",&createdViewId)
		log.Debugf("View save result: %d with id: %s", viewSaveResult,createdViewId)

		shareSavedSearchBaseEndpoint := "/authz/shares/entities/"
		if len(*savedSearch.Users) != 0 {
			share := CreateShareApiObjectWithUserNames(config,*savedSearch.Users)
			preparedViewId := PrepareViewIdForShare(createdViewId.Id)
			var result api.ResourceCreationResponse
			shareResult := client.PostToGraylog(config,share, shareSavedSearchBaseEndpoint + preparedViewId,&result)
			log.Debugf("View Share result: %d with id: %s", shareResult,result)
		}
	}
}
