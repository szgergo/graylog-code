package reader

import (
	"com/github/graylog-code/config/template"
	"com/github/graylog-code/graylog/api"
	"com/github/graylog-code/graylog/client"
	log "github.com/sirupsen/logrus"
	"strings"
)

type savedSearchApi struct {
	Views *[]api.ViewApi
}

type grant struct {
	Target string
	GranteeTitle string `json:"grantee_title,omitempty"`
}

type grantsOverview struct {
	Grants *[]grant
}

func GetSavedSearches(config template.GraylogConfigTemplate) []template.SavedSearchTemplate {

	var savedSearches = make([]template.SavedSearchTemplate,0)
	var savedSearchesFromGraylog = savedSearchApi{}
	var endpoint = "/search/saved?page=1&per_page=50&sort=title&order=asc"
	client.GetFromGraylog(config,endpoint,&savedSearchesFromGraylog)

	for i, savedSearch := range *savedSearchesFromGraylog.Views {
		if savedSearch.Owner != config.User {
			log.Debugf("User of saved search %s is not %s,skipping", savedSearch.Title,config.User)
			continue
		}
		var search = api.SearchApi{}
		log.Debugf("Checking %d. saved search",i + 1)
		var searchEndpoint = "/views/search/" + savedSearch.SearchId
		client.GetFromGraylog(config,searchEndpoint,&search)
		var actualSavedSearch = template.SavedSearchTemplate{}
		actualSavedSearch.Id = savedSearch.Id
		actualSavedSearch.Name = savedSearch.Title
		actualSavedSearch.Summary = savedSearch.Summary
		actualSavedSearch.Description = savedSearch.Description
		actualSavedSearch.SearchQuery = search.Queries[0].Query.QueryString
		if search.Queries[0].Filter != nil && len(*search.Queries[0].Filter.Filters) != 0 {
			var actualFilter = make([]string, 0)
			for j, innerFilter := range *search.Queries[0].Filter.Filters {
				log.Debugf("Checking %d. innerFilter", j+1)
				actualFilter = append(actualFilter, getTitleForStreamById(config,innerFilter.Id))
			}
			actualSavedSearch.Filters = &actualFilter
		}

		from := search.Queries[0].Timerange.From
		to := search.Queries[0].Timerange.To
		actualSavedSearch.TimeRange = &template.TimeRangeTemplate{}
		if from != nil && to != nil {

			actualSavedSearch.TimeRange.From = from
			actualSavedSearch.TimeRange.To = to
		}

		actualSavedSearch.TimeRange.Range = search.Queries[0].Timerange.Range
		actualSavedSearch.TimeRange.Type = search.Queries[0].Timerange.Type

		actualSavedSearch.Users = getUsersSearchSharedWith(config,savedSearch.Id)

		savedSearches = append(savedSearches,actualSavedSearch)
	}
	return savedSearches
}

func getUsersSearchSharedWith(config template.GraylogConfigTemplate,viewId string) *[]string {
	var users = make([]string, 0)
	var endpoint = "/authz/grants-overview"
	var grantsOverview grantsOverview
	client.GetFromGraylog(config,endpoint,&grantsOverview)
	userToken := ": "
	for i, grant := range *grantsOverview.Grants {
		log.Debugf("Checking %d. user",i + 1)
		isTarget := grant.Target == "grn::::search:" + viewId
		isUserShare := strings.Contains(grant.GranteeTitle, "user: ")
		if isTarget && isUserShare {
			users = append(users, strings.Split(grant.GranteeTitle, userToken)[1])
		}
	}
	return &users
}

func getTitleForStreamById(config template.GraylogConfigTemplate, streamId string) string {
	var streamEndpoint = "/streams/" + streamId
	var stream api.StreamApi
	client.GetFromGraylog(config,streamEndpoint,&stream)
	return stream.Title
}
