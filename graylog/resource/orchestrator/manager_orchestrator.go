package orchestrator

import (
	"com/github/graylog-code/config/template"
	"com/github/graylog-code/graylog/resource/manager"
	log "github.com/sirupsen/logrus"
	"sync"
)

var managers = map[string]func(template.GraylogConfigTemplate) {
	"Inputs manager": manager.ManageInputs,
	"Saved searches manager": manager.ManageSavedSearches,
	"Streams manager": manager.ManageStreams,
	"Notifications manager": manager.ManageNotifications,
}

func StartManagers(config template.GraylogConfigTemplate) {
	var wg sync.WaitGroup
	wg.Add(len(managers))
	for stepName, step := range managers {
		go managerWrapper(config, step, &wg, stepName)
	}
	wg.Wait()
}

func managerWrapper(config template.GraylogConfigTemplate,
	step func(template.GraylogConfigTemplate),
	group *sync.WaitGroup,
	stepName string) {
	defer group.Done()
	log.Infof("Starting manager: %s", stepName)
	step(config)
	log.Infof("Finished manager: %s", stepName)
}
