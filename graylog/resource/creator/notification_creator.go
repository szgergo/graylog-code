package creator

import (
	"com/github/graylog-code/config/template"
	"com/github/graylog-code/graylog/api"
	"com/github/graylog-code/graylog/client"
)

func CreateNotification(config template.GraylogConfigTemplate,
	notificationTemplates []template.NotificationTemplate) {
	endpoint := "/events/notifications"
	for _, notification := range notificationTemplates {
		var notificationRequest api.NotificationApi
		notificationRequest.Title = notification.Title
		notificationRequest.Description = notification.Description

		notificationConfig := notification.Config
		if notificationConfig != nil {
			var notificationConfigRequest api.NotificationConfigApi
			notificationRequest.Config = &notificationConfigRequest
			notificationType := api.NotificationTypeMapping[notificationConfig.Type]
			notificationConfigRequest.Type = notificationType.NotificationType()

			url := notificationConfig.Url
			if url != nil && notificationConfig.Type == "HTTP notification"{
				WhiteListUrl(config,url.Url,url.Type,url.Title)
				notificationConfigRequest.Url = url.Url
			}
			notificationConfigRequest.BodyTemplate = notificationConfig.BodyTemplate
			notificationConfigRequest.CallbackType =  notificationType.CallBackType()
			notificationConfigRequest.EmailRecipients = notificationConfig.EmailRecipients
			notificationConfigRequest.UserRecipients = notificationConfig.UserRecipients
			notificationConfigRequest.Sender = notificationConfig.Sender
			notificationConfigRequest.Subject = notificationConfig.Subject

			legacyCallbackConfiguration := notificationConfig.LegacyCallbackConfiguration
			if legacyCallbackConfiguration != nil {
				var legacyCallbackConfigurationRequest api.LegacyCallbackConfigurationApi
				notificationConfigRequest.Configuration = &legacyCallbackConfigurationRequest
				legacyCallbackConfigurationRequest.Subject = legacyCallbackConfiguration.Subject
				legacyCallbackConfigurationRequest.Sender = legacyCallbackConfiguration.Sender
				legacyCallbackConfigurationRequest.UserRecipients = legacyCallbackConfiguration.UserRecipients
				legacyCallbackConfigurationRequest.EmailRecipients = legacyCallbackConfiguration.EmailRecipients
				legacyCallbackConfigurationRequest.Body = legacyCallbackConfiguration.Body
				legacyCallbackConfigurationRequest.Url = legacyCallbackConfiguration.Url
			}
		}
		client.PostToGraylog(config,&notificationRequest,endpoint,nil)
	}
}
