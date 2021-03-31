package deleter

import (
	"com/github/graylog-code/config/template"
	"com/github/graylog-code/graylog/client"
	log "github.com/sirupsen/logrus"
)

func DeleteStreams(template template.GraylogConfigTemplate,streams []template.StreamTemplate) {
	for i, stream := range streams {
		log.Debugf("Deleting %d. stream with name: %s", i+1, stream.Title)
		if stream.Id == "" {
			log.Errorf("Stream %s does not have id, skipping deletion", stream.Title)
			continue
		}
		responseCode := client.DeleteFromGraylog(template,"/streams/" + stream.Id)
		if client.ResponseIsFailure(responseCode) {
			log.Errorf("Couldn't delete stream with name: %s",stream.Title)
		}
	}
}
