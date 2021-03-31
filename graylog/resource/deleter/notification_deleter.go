package deleter

import (
	"com/github/graylog-code/config/template"
	"com/github/graylog-code/graylog/api"
	"com/github/graylog-code/graylog/client"
)

func DeleteNotification(config template.GraylogConfigTemplate,notifications []template.NotificationTemplate) {
	endpoint := "/events/notifications/"
	for _, notification := range notifications {
		client.DeleteFromGraylog(config,endpoint + notification.Id)
	}
}

func isHttpNotification(notification template.NotificationTemplate) bool {
	return api.NotificationTypeMapping[notification.Config.Type].NotificationType() == "http-notification-v1"
}

func removeElementFromArray(s []api.WhiteListedUrl, index int) []api.WhiteListedUrl {
	s[len(s)-1], s[index] = s[index], s[len(s)-1]
	return s[:len(s)-1]
}

func checkIfUrlsAreTheSame(a api.WhiteListedUrl, b template.UrlConfigTemplate) bool {
	return a.Type == b.Type &&
		a.Title == b.Title &&
		a.Value == b.Url
}