package creator

import (
	"com/github/graylog-code/config/template"
	"com/github/graylog-code/graylog/api"
	"com/github/graylog-code/graylog/client"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

func createSearch(config template.GraylogConfigTemplate,savedSearch template.SavedSearchTemplate) api.SearchApi {
	var searchToSave api.SearchApi

	searchToSave.Id = bson.NewObjectId().Hex()
	searchToSave.Queries = make([]*api.OuterQueryApi, 1)

	var outerQuery api.OuterQueryApi
	searchToSave.Queries[0] = &outerQuery

	outerQuery.Id = strings.ToLower(uuid.NewString())

	streamsToAddAsFilter := savedSearch.Filters
	numberOfStreamsToAddAsFilter := len(*streamsToAddAsFilter)
	if numberOfStreamsToAddAsFilter != 0 {
		filter := api.FilterApi{}
		filter.Type = "or"
		outerQuery.Filter = &filter
		innerFilters := make([]api.InnerFilterApi, numberOfStreamsToAddAsFilter)
		filter.Filters = &innerFilters
		for i, stream := range *streamsToAddAsFilter {
			log.Debugf("Adding %d. stream (name: %s) as filter to saved search: %s ",
				i + 1,
				stream,
				savedSearch.Name)

			innerFilter := api.InnerFilterApi{}
			innerFilter.Type = "stream"
			innerFilter.Id = getGraylogStreamIdByName(config,stream)
			innerFilters[i] = innerFilter
		}
	}

	var outerQueryTimeRange api.TimeRangeApi
	outerQuery.Timerange = &outerQueryTimeRange

	outerQueryTimeRange.Type = savedSearch.TimeRange.Type
	outerQueryTimeRange.Range = savedSearch.TimeRange.Range
	outerQueryTimeRange.From = savedSearch.TimeRange.From
	outerQueryTimeRange.To = savedSearch.TimeRange.To

	var innerQuery api.InnerQueryApi
	outerQuery.Query = &innerQuery

	innerQuery.Type = "elasticsearch"
	innerQuery.QueryString = savedSearch.SearchQuery

	outerQuery.SearchTypes = make([]*api.SearchTypeApi, 1)

	var allMessagesSearchType api.SearchTypeApi
	outerQuery.SearchTypes[0] = &allMessagesSearchType

	allMessagesSearchType.Id = strings.ToLower(uuid.NewString())
	allMessagesSearchType.Streams = make([]string, 0)
	allMessagesSearchType.Name = "All Messages"
	allMessagesSearchType.Limit = 150
	allMessagesSearchType.Offset = 0

	var sort api.SortApi
	sort.Field = "timestamp"
	sort.Order = "DESC"

	allMessagesSearchType.Sort = make([]*api.SortApi, 1)
	allMessagesSearchType.Sort[0] = &sort
	allMessagesSearchType.Decorators = make([]string, 0)
	allMessagesSearchType.Type = "messages"

	return searchToSave
}

func getGraylogStreamIdByName(config template.GraylogConfigTemplate,name string) string {
	var streams api.StreamsApi
	client.GetFromGraylog(config,"/streams",&streams)
	for i, stream := range *streams.Streams {
		log.Debugf("Checking %d. stream",i)
		if strings.EqualFold(stream.Title, name) {
			log.Debugf("Found id(%s) for stream: %s ",stream.Id,name)
			return stream.Id
		}
	}
	return ""
}
