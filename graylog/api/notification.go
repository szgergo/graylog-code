package api

type LegacyCallbackConfigurationApi struct {
	Url string `json:"url,omitempty"`
	Sender string `json:"sender,omitempty"`
	Subject string `json:"subject,omitempty"`
	Body string `json:"body,omitempty"`
	EmailRecipients *[]string `json:"email_recipients,omitempty"`
	UserRecipients *[]string `json:"user_recipients,omitempty"`
}

type NotificationConfigApi struct {
	Type string `json:"type,omitempty"`
	CallbackType string `json:"callback_type,omitempty"`
	Sender string `json:"sender,omitempty"`
	Subject string `json:"subject,omitempty"`
	BodyTemplate string `json:"body_template,omitempty"`
	EmailRecipients *[]string `json:"email_recipients,omitempty"`
	UserRecipients *[]string `json:"user_recipients,omitempty"`
	Url string `json:"url,omitempty"`
	Configuration *LegacyCallbackConfigurationApi `json:"configuration,omitempty"`
}

type NotificationApi struct {
	Id string `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Config *NotificationConfigApi `json:"config,omitempty"`
}
