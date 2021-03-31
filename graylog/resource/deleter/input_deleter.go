package deleter

import (
	"com/github/graylog-code/config/template"
	"com/github/graylog-code/graylog/client"
	log "github.com/sirupsen/logrus"
)

func DeleteInputs(config template.GraylogConfigTemplate,inputConfigs []template.InputTemplate) {
	for i,inputConfig := range inputConfigs {
		id := inputConfig.Id
		if id == "" {
			log.Error("Can't delete input, id is null")
		}
		log.Debugf("Deleting %d. input, named: %s", i + 1, inputConfig.Title)
		responseCode := client.DeleteFromGraylog(config,"/system/inputs/"+id)
		if responseCode != 204 {
			log.Errorf("Couldn't delete input with id: %s", id)
		}
	}
}
