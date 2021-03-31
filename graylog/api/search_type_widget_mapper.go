package api

//This class maps the search_type field to the type field of WidgetApi type (in view_creator.go)

var SearchTypeWidget = map[string]string {
	"pivot": "aggregation",
	"messages" : "messages",
}
