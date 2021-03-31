package manager

import (
	"com/github/graylog-code/config/template"
	resourceCreator "com/github/graylog-code/graylog/resource/creator"
	resourceDeleter "com/github/graylog-code/graylog/resource/deleter"
	resourceReader "com/github/graylog-code/graylog/resource/reader"
	log "github.com/sirupsen/logrus"
)

func ManageNotifications(config template.GraylogConfigTemplate) {
	streamsFromGraylog := resourceReader.GetNotifications(config)
	log.Debug("Notifications from graylog: ", streamsFromGraylog)
	streamsFromConfig := config.Notifications
	log.Debug("Notifications from config: ", streamsFromConfig)

	needsToBeDeleted := template.NotificationComplement(streamsFromGraylog, streamsFromConfig)
	needsToBeAdded := template.NotificationComplement(streamsFromConfig, streamsFromGraylog)

	if len(needsToBeDeleted) == 0 && len(needsToBeAdded) == 0 {
		log.Info("All notifications are up to date.")
	} else {
		resourceCreator.CreateNotification(config, needsToBeAdded)
		resourceDeleter.DeleteNotification(config, needsToBeDeleted)
	}
}
