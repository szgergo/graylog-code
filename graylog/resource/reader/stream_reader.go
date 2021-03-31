package reader

import (
	"com/github/graylog-code/config/template"
	"com/github/graylog-code/graylog/api"
	"com/github/graylog-code/graylog/client"
	"com/github/graylog-code/graylog/resource/creator"
	"com/github/graylog-code/util"
	log "github.com/sirupsen/logrus"
)

var systemStreams = []string{"All events","All messages","All system events"}

func GetStreams(config template.GraylogConfigTemplate) []template.StreamTemplate {
	var streamsFromGraylog api.StreamsApi
	var streamTemplates []template.StreamTemplate
	client.GetFromGraylog(config,"/streams",&streamsFromGraylog)

	for i, streamFromGraylog := range *streamsFromGraylog.Streams {
		log.Debugf("Checking %d. stream", i + 1)
		if streamFromGraylog.User != config.User {
			log.Debugf("Stream %s's owner is not %s, skipping", streamFromGraylog.Title,config.User)
			continue
		}
		if util.StringInSlice(streamFromGraylog.Title,systemStreams) {
			log.Debugf("Stream %s is a system stream, skipping", streamFromGraylog.Title)
			continue
		}
		var streamTemplate template.StreamTemplate
		streamTemplate.Title = streamFromGraylog.Title
		streamTemplate.ContentPack = streamFromGraylog.ContentPack
		streamTemplate.MatchingType = streamFromGraylog.MatchingType
		streamTemplate.Description = streamFromGraylog.Description
		streamTemplate.RemoveMatchesFromDefaultStream = streamFromGraylog.RemoveMatchesFromDefaultStream
		streamTemplate.IndexSetName = creator.GetIndexSetTitleById(config,streamFromGraylog.IndexSetId)
		streamTemplate.Id = streamFromGraylog.Id
		if streamFromGraylog.Rules != nil {
			var streamRules []template.StreamRuleTemplate
			for i, streamRuleFromGraylog := range *streamFromGraylog.Rules {
				log.Debugf("Checking %d. stream rule for stream: %s", i+1, streamFromGraylog.Title)
				var actualStreamRule template.StreamRuleTemplate
				actualStreamRule.Type = getStreamRuleTypeNameById(config,
					streamFromGraylog.Id,
					streamRuleFromGraylog.Type)
				actualStreamRule.Description = streamRuleFromGraylog.Description
				actualStreamRule.Field = streamRuleFromGraylog.Field
				actualStreamRule.Inverted = streamRuleFromGraylog.Inverted
				actualStreamRule.Value = streamRuleFromGraylog.Value
				streamRules = append(streamRules,actualStreamRule)
			}
			streamTemplate.Rules = &streamRules
		}
		streamTemplates = append(streamTemplates,streamTemplate)
	}
	return streamTemplates
}

func getStreamRuleTypeNameById(config template.GraylogConfigTemplate,
	streamId string,
	desiredStreamRuleType int) string {
	var endpoint = "/streams/" + streamId + "/rules/types"
	var streamRuleTypeResponse *[]api.StreamRuleType
	client.GetFromGraylog(config,endpoint,&streamRuleTypeResponse)

	for i, streamRuleType := range *streamRuleTypeResponse {
		log.Debugf("Checking %d. stream rule type", i + 1)
		if streamRuleType.Id == desiredStreamRuleType {
			return streamRuleType.ShortDescription
		}
	}

	return ""
}

