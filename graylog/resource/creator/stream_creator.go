package creator

import (
	"com/github/graylog-code/config/template"
	"com/github/graylog-code/graylog/api"
	"com/github/graylog-code/graylog/client"
	log "github.com/sirupsen/logrus"
)

type indexSetApi struct {
	Id string `json:"id"`
	Title string `json:"title"`
}

type indexSetsApi struct {
	Total int `json:"total"`
	IndexSets *[]indexSetApi `json:"index_sets"`
}

func CreateStreams(template template.GraylogConfigTemplate, streams []template.StreamTemplate) {
	var streamCreationEndpoint = "/streams"
	for i,stream := range streams {
		log.Debugf("Start creating %d. stream with name: %s",i+1, stream.Title)
		var requestBody api.StreamCreationApi
		requestBody.Title = stream.Title
		requestBody.RemoveMatchesFromDefaultStream = stream.RemoveMatchesFromDefaultStream
		requestBody.Description = stream.Description
		requestBody.MatchingType = stream.MatchingType
		requestBody.IndexSetId = getIndexSetIdByTitle(template,stream.IndexSetName)
		if requestBody.IndexSetId == "" {
			log.Errorf("Unable to set IndexSet for stream %s, skipping",stream.Title)
			continue
		}
		requestBody.ContentPack = stream.ContentPack

		var response api.StreamCreationResponseApi
		client.PostToGraylog(template,requestBody,streamCreationEndpoint,&response)
		if stream.AutoStart && response.Id != "" {
			log.Debugf("Starting stream: %s", stream.Title)
			client.PostToGraylog(template,nil,"/streams/" + response.Id + "/resume",nil)
		}

		if stream.Rules != nil {
			var streamRuleCreationEndpoint = "/streams/" + response.Id + "/rules"
			for i, rule := range *stream.Rules {
				log.Debugf("Saving %d. rule", i + 1)
				var streamRuleRequest api.StreamRuleCreationApi
				streamRuleRequest.Value = rule.Value
				streamRuleRequest.Inverted = rule.Inverted
				streamRuleRequest.Type = getStreamRuleTypeIdByType(template,rule.Type,response.Id)
				if streamRuleRequest.Type == -1 {
					log.Errorf("Couldn't gather stream rule type id for rule, skipping")
					continue
				}
				streamRuleRequest.Field = rule.Field
				streamRuleRequest.Description = rule.Description
				client.PostToGraylog(template,&streamRuleRequest,streamRuleCreationEndpoint,nil)
			}
		}
	}
}

func getStreamRuleTypeIdByType(template template.GraylogConfigTemplate,
	streamRuleName string,
	streamId string) int {
	var endpoint = "/streams/" + streamId + "/rules/types"
	var streamRuleTypes []api.StreamRuleType
	client.GetFromGraylog(template,endpoint,&streamRuleTypes)

	if len(streamRuleTypes) != 0 {
		for i, streamRuleType := range streamRuleTypes {
			log.Debugf("Checking %d. stream rule type, with name: ", i+ 1, streamRuleType.Name)
			if streamRuleType.ShortDescription == streamRuleName {
				return streamRuleType.Id
			}
		}
	}
	return -1
}

func getIndexSetIdByTitle(template template.GraylogConfigTemplate,
	title string) string {
	var response indexSetsApi
	var endpoint = "/system/indices/index_sets?stats=false"
	client.GetFromGraylog(template,endpoint,&response)

	if response.IndexSets != nil {
		for i, stream := range *response.IndexSets {
			log.Debugf("Checking %d stream against desired name: %s", i+1,title)
			if stream.Title == title {
				return stream.Id
			}
		}
	}
	return ""
}

func GetIndexSetTitleById(configTemplate template.GraylogConfigTemplate,id string) string {
	var response indexSetsApi
	var endpoint = "/system/indices/index_sets?stats=false"
	client.GetFromGraylog(configTemplate,endpoint,&response)

	if response.IndexSets != nil {
		for i, stream := range *response.IndexSets {
			log.Debugf("Checking %d stream against desired id: %s", i+1,id)
			if stream.Id == id {
				return stream.Title
			}
		}
	}
	return ""
}