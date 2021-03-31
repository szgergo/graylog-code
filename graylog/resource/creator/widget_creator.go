package creator

import "com/github/graylog-code/graylog/api"

func CreateAllMessagesWidget(widget *api.WidgetApi) api.WidgetApi {
	widget.Streams = make([]string,0)
	var widgetConfig api.WidgetConfigApi
	widget.Config = &widgetConfig
	widgetConfig.Fields = &[]string{"timestamp","source"}
	widgetConfig.ShowMessageRow = true
	decorators := make([]string,0)
	widgetConfig.Decorators = &decorators
	sort := api.SortApi{}
	widgetConfig.Sort = make([]*api.SortApi, 1)
	widgetConfig.Sort[0] = &sort
	sort.Type = "pivot"
	sort.Field = "timestamp"
	sort.Direction = "Descending"
	return *widget
}

func CreateAggregationWidget(widget *api.WidgetApi) api.WidgetApi {
	return *widget
}
