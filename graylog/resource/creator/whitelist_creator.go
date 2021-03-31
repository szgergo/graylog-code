package creator

import (
	"com/github/graylog-code/config/template"
	"com/github/graylog-code/graylog/api"
	"com/github/graylog-code/graylog/client"
	"encoding/json"
	"github.com/google/uuid"
	"strings"
)

type urlWhiteListCheckResponseApi struct {
	Url string `json:"url"`
	IsWhiteListed bool `json:"is_whitelisted"`
}

func IsUrlWhiteListed(template template.GraylogConfigTemplate, url string) bool {
	endpoint := "/system/urlwhitelist/check"
	//This is needed as /api/system/whitelist/check endpoint needs json data with quotes intact.
	//So we can't use normal Unmarshalling, but the json.RawMessage type
	Source := json.RawMessage(`{"url":"` + url + `"}`)
	var response urlWhiteListCheckResponseApi
	client.PostToGraylog(template,Source,endpoint,&response)

	if response.Url != "" {
		return response.IsWhiteListed
	}
	return false
}

func WhiteListUrl(configTemplate template.GraylogConfigTemplate,
	url string,
	_type string,
	title string) bool {
	var endpoint = "/system/urlwhitelist"
	var whiteListedUrls api.WhiteListApi
	if IsUrlWhiteListed(configTemplate,url) {
		return true
	}
	client.GetFromGraylog(configTemplate,endpoint,&whiteListedUrls)
	var urlToWhiteList api.WhiteListedUrl
	urlToWhiteList.Id = strings.ToLower(uuid.NewString())
	urlToWhiteList.Type = _type
	urlToWhiteList.Value = url
	urlToWhiteList.Title = title

	*whiteListedUrls.Entries = append(*whiteListedUrls.Entries,urlToWhiteList)
	responseCode, _ := client.PutToGraylog(configTemplate,&whiteListedUrls,endpoint)

	if client.ResponseIsFailure(responseCode) {
		return false
	}
	return true
}
