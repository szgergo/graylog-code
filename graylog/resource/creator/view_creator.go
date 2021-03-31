package creator

import (
	"com/github/graylog-code/config/template"
	error2 "com/github/graylog-code/error"
	"com/github/graylog-code/graylog/api"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

func CreateView(search api.SearchApi,
	savedSearchConfig template.SavedSearchTemplate) api.ViewApi {
	var view api.ViewApi

	view.Type = "SEARCH"
	view.Title = savedSearchConfig.Name
	view.Summary = ""
	view.Description = ""
	view.SearchId = search.Id
	view.Properties = make([]string, 0)

	stateMap := make(map[string]*api.WidgetSettingsApi)
	var widgetSettings api.WidgetSettingsApi
	searchQuery := getQueryFromSearch(search)
	stateMap[searchQuery.Id] = &widgetSettings
	view.State = stateMap
	view.Summary = savedSearchConfig.Summary
	view.Description = savedSearchConfig.Description

	var widgetList = api.WidgetListApi{}
	widgetList.Widget = make(map[string]string)
	var numberOfWidgets = len(search.Queries[0].SearchTypes)
	widgetSettings.Widgets = make([]*api.WidgetApi, numberOfWidgets)
	widgetSettings.Titles = &widgetList
	widgetMapping := make(map[string][]string)
	widgetSettings.WidgetMapping = &widgetMapping
	widgetSettings.Positions = make(map[string]*api.WidgetPositionApi)
	for i, searchType := range searchQuery.SearchTypes {
		log.Debugf("Creating %d. widget with name %s", i + 1,searchType.Name)
		widgetId := uuid.NewString()
		widgetList.Widget[widgetId] = searchType.Name
		var widget = api.WidgetApi{}
		widget.Id = widgetId
		widget.Type = api.SearchTypeWidget[searchType.Type]
		widgetSettings.Widgets[i] = &widget
		switch widget.Type {
		case "messages":
			CreateAllMessagesWidget(&widget)
		case "aggregation":
			CreateAggregationWidget(&widget)
		}
		searchTypeWidgetIdArray := make([]string, 1)
		searchTypeWidgetIdArray[0] = searchType.Id
		widgetMapping[widgetId] = searchTypeWidgetIdArray
		widgetPosition := api.WidgetPositionApi{}
		widgetSettings.Positions[widgetId] = &widgetPosition
		widgetPosition.Col = 1
		widgetPosition.Row = 1
		widgetPosition.Height = 2
		widgetPosition.Width = "Infinity"
	}
	return view
}


func getQueryFromSearch(search api.SearchApi) *api.OuterQueryApi {
	// Currently we are using only one query.
	numberOfQueries := len(search.Queries)
	if len(search.Queries) != 1 {
		error2.ThrowError("Insufficient number of query was found in search object, number: ",  numberOfQueries)
	}
	return search.Queries[0]
}