package api

type NotificationType struct {
	notificationType string
	callBackType string
}

func (nt NotificationType) NotificationType() string{
	return nt.notificationType
}

func (nt NotificationType) CallBackType() string {
	return nt.callBackType
}

var NotificationTypeMapping = map[string]NotificationType{
	"HTTP notification" : { "http-notification-v1", ""},
	"Email notification" : {"email-notification-v1",""},
	"Legacy HTTP Callback" : {"legacy-alarm-callback-notification-v1", "org.graylog2.alarmcallbacks.HTTPAlarmCallback" },
	"Legacy Email Callback" : {"legacy-alarm-callback-notification-v1", "org.graylog2.alarmcallbacks.EmailAlarmCallback" },
}

func GetNotificationTemplateTypeByNotificationTypeFromGraylog(value string) (key string, ok bool) {
	for k, v := range NotificationTypeMapping {
		if v.NotificationType() == value {
			key = k
			ok = true
			return
		}
	}
	return
}