package reader

import (
	"com/github/graylog-code/config/template"
	"com/github/graylog-code/graylog/api"
	"com/github/graylog-code/graylog/client"
)

type notificationsApi struct {
	Notifications *[]api.NotificationApi
}

func GetNotifications(config template.GraylogConfigTemplate) []template.NotificationTemplate {
	endpoint := "/events/notifications?page=1&per_page=100"
	var notificationsFromGraylog notificationsApi
	var toReturn []template.NotificationTemplate
	client.GetFromGraylog(config,endpoint,&notificationsFromGraylog)

	for _, fromGraylog := range *notificationsFromGraylog.Notifications {
		var notification template.NotificationTemplate
		notification.Title = fromGraylog.Title
		notification.Description = fromGraylog.Description
		notification.Id = fromGraylog.Id
		notification.Config = &template.NotificationConfigTemplate{}
		notification.Config.Type, _ = api.GetNotificationTemplateTypeByNotificationTypeFromGraylog(fromGraylog.Config.Type)
		if fromGraylog.Config.Type == "http-notification-v1" {
			notification.Config.Url = getUrlInformation(config,fromGraylog.Config.Url)
		}
		notification.Config.EmailRecipients = fromGraylog.Config.EmailRecipients
		notification.Config.UserRecipients = fromGraylog.Config.UserRecipients
		notification.Config.Sender = fromGraylog.Config.Sender
		notification.Config.Subject = fromGraylog.Config.Subject
		notification.Config.BodyTemplate = fromGraylog.Config.BodyTemplate
		notification.Config.LegacyCallbackConfiguration = createLegacyCallbackConfig(fromGraylog.Config.Configuration)
		toReturn = append(toReturn,notification)
	}
	return toReturn
}

func createLegacyCallbackConfig(configuration *api.LegacyCallbackConfigurationApi) *template.LegacyNotificationCallbackConfigurationTemplate {
	var toReturn template.LegacyNotificationCallbackConfigurationTemplate

	if configuration == nil {
		return nil
	}
	toReturn.Subject = configuration.Subject
	toReturn.Sender = configuration.Sender
	toReturn.UserRecipients = configuration.UserRecipients
	toReturn.EmailRecipients = configuration.EmailRecipients
	toReturn.Url = configuration.Url
	toReturn.Body = configuration.Body

	return &toReturn
}

func getUrlInformation(config template.GraylogConfigTemplate,
	url string) *template.UrlConfigTemplate {
	var urlConfig template.UrlConfigTemplate
	var whiteListedUrls api.WhiteListApi
	endpoint := "/system/urlwhitelist"
	client.GetFromGraylog(config,endpoint, &whiteListedUrls)

	for _, whiteListedUrl := range *whiteListedUrls.Entries {
		if whiteListedUrl.Value == url {
			urlConfig.Url = whiteListedUrl.Value
			urlConfig.Type = whiteListedUrl.Type
			urlConfig.Title = whiteListedUrl.Title
			urlConfig.Id = whiteListedUrl.Id
			return &urlConfig
		}
	}
return nil
}
