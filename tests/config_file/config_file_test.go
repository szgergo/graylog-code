package config_file

import (
	"com/github/graylog-code/config/handler"
	"com/github/graylog-code/config/template"
	"com/github/graylog-code/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraylogConfigFieldsWithRelativeTimeRange(t *testing.T) {

	config :=
		handler.Parse(util.LoadConfigurationFile("assets/test_basic_one_saved_search.yml"))

	savedSearch := assertAndGetBasicData(t, config)
	var timeRange = savedSearch.TimeRange
	assert.Equal(t, "relative", timeRange.Type)
	assert.Equal(t, 300, timeRange.Range)
	assert.Nil(t, timeRange.To)
	assert.Nil(t, timeRange.From)
}

func TestGraylogConfigFieldsWithAbsoluteTimeRange(t *testing.T) {
	config :=
		handler.Parse(util.LoadConfigurationFile("assets/test_basic_one_saved_search_absolute.yml"))

	savedSearch := assertAndGetBasicData(t, config)
	var timeRange = savedSearch.TimeRange
	assert.Equal(t, "absolute", timeRange.Type)
	assert.Equal(t, 0,timeRange.Range)
	assert.NotNil(t, timeRange.To)
	assert.NotNil(t, timeRange.From)

	var from =  *timeRange.From
	var to = *timeRange.To

	assert.Equal(t, "2021-04-01 13:02:13 +0000 UTC", from.String())
	assert.Equal(t, "2021-04-02 13:02:13 +0000 UTC", to.String())
}

func TestGraylogConfigFieldsWithUsers(t *testing.T) {
	config :=
		handler.Parse(util.LoadConfigurationFile("assets/test_basic_one_saved_search_with_users.yml"))

	savedSearch := assertAndGetBasicData(t, config)
	var timeRange = savedSearch.TimeRange
	assert.Equal(t, "relative", timeRange.Type)
	assert.Equal(t, 300, timeRange.Range)
	assert.Nil(t, timeRange.To)
	assert.Nil(t, timeRange.From)
	assert.Equal(t, 2,len(*savedSearch.Users))
	var users = *savedSearch.Users
	assert.ElementsMatch(t, []string{"test1","test2"},users)
}

func TestGraylogConfigFieldsWithFilters(t *testing.T) {
	config :=
		handler.Parse(util.LoadConfigurationFile("assets/test_basic_one_saved_search_with_filter.yml"))

	savedSearch := assertAndGetBasicData(t, config)
	var timeRange = savedSearch.TimeRange
	assert.Equal(t, "relative", timeRange.Type)
	assert.Equal(t, 300, timeRange.Range)
	assert.Nil(t, timeRange.To)
	assert.Nil(t, timeRange.From)
	assert.Nil(t, savedSearch.Users)
	assert.ElementsMatch(t, []string{"All events"}, *savedSearch.Filters)
}

func assertAndGetBasicData(t *testing.T, config template.GraylogConfigTemplate) template.SavedSearchTemplate {
	assert.Equal(t, "http://test.com:9000", config.GraylogServer)
	assert.Equal(t, "TEST_ACCESS_TOKEN", config.AccessToken)
	assert.Equal(t, "/api", config.ApiPrefix)
	assert.Equal(t, "4.0.0", config.RequiredGraylogVersion)
	assert.Equal(t, 1, len(config.SavedSearches))
	var savedSearch = config.SavedSearches[0]
	assert.NotNil(t, savedSearch)
	assert.Equal(t, "Test", savedSearch.Name)
	assert.Equal(t, "Test", savedSearch.SearchQuery)
	assert.Equal(t, "Test", savedSearch.Description)
	assert.Equal(t, "Test", savedSearch.Summary)
	assert.NotNil(t, savedSearch.TimeRange)
	return savedSearch
}