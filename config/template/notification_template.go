package template

import "com/github/graylog-code/util"

type LegacyNotificationCallbackConfigurationTemplate struct {
	Url string `yaml:"url,omitempty"`
	Sender string `yaml:"sender,omitempty"`
	Subject string `yaml:"subject,omitempty"`
	Body string `yaml:"body,omitempty"`
	EmailRecipients *[]string `yaml:"email_recipients,omitempty"`
	UserRecipients *[]string `yaml:"user_recipients,omitempty"`
}

func (lncct LegacyNotificationCallbackConfigurationTemplate) Equals(other LegacyNotificationCallbackConfigurationTemplate) bool {
	return lncct.Url == other.Url &&
		lncct.Sender == other.Sender &&
		lncct.Subject == other.Subject &&
		lncct.Body == other.Body &&
		stringArrayPointerEquals(lncct.EmailRecipients,other.EmailRecipients) &&
		stringArrayPointerEquals(lncct.UserRecipients,other.UserRecipients)
}

type UrlConfigTemplate struct {
	Id string
	Url string `yaml:"url,omitempty"`
	Type string `yaml:"type,omitempty"`
	Title string `yaml:"title,omitempty"`
}

func (ut UrlConfigTemplate) Equals(other UrlConfigTemplate) bool {
	return ut.Url == other.Url &&
		ut.Type == other.Type &&
		ut.Title == other.Title
}

type NotificationConfigTemplate struct {
	Type string `yaml:"type,omitempty"`
	Sender string `yaml:"sender,omitempty"`
	Subject string `yaml:"subject,omitempty"`
	BodyTemplate string `yaml:"body_template,omitempty"`
	EmailRecipients *[]string `yaml:"email_recipients,omitempty"`
	UserRecipients *[]string `yaml:"user_recipients,omitempty"`
	Url *UrlConfigTemplate `yaml:"url,omitempty"`
	LegacyCallbackConfiguration *LegacyNotificationCallbackConfigurationTemplate `yaml:"legacy_callback_configuration,omitempty"`
}

func (nct NotificationConfigTemplate) Equals(other NotificationConfigTemplate) bool {
	return nct.Type == other.Type &&
		nct.Sender == other.Sender &&
		nct.Subject == other.Subject &&
		nct.BodyTemplate == other.BodyTemplate &&
		stringArrayPointerEquals(nct.EmailRecipients,other.EmailRecipients) &&
		stringArrayPointerEquals(nct.UserRecipients,other.UserRecipients) &&
		urlConfigPointerEquals(nct.Url,other.Url) &&
		legacyNotificationCallbackConfigurationEquals(nct.LegacyCallbackConfiguration,other.LegacyCallbackConfiguration)
}

func legacyNotificationCallbackConfigurationEquals(a *LegacyNotificationCallbackConfigurationTemplate,
	b *LegacyNotificationCallbackConfigurationTemplate) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Equals(*b)
}

func urlConfigPointerEquals(a *UrlConfigTemplate, b *UrlConfigTemplate) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Equals(*b)
}

type NotificationTemplate struct {
	Id string
	Title string
	Description string
	Config *NotificationConfigTemplate
}

func (nt NotificationTemplate) Equals(other NotificationTemplate) bool {
	return nt.Title == other.Title &&
		nt.Description == other.Description &&
		notificationConfigPointerEquals(nt.Config,other.Config)

}

func notificationConfigPointerEquals(a *NotificationConfigTemplate, b *NotificationConfigTemplate) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return a.Equals(*b)
}

func stringArrayPointerEquals(a *[]string, b *[]string) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return util.StringArrayEquals(*a,*b)
}

/*
From: http://web.mnstate.edu/peil/MDEV102/U1/S6/Complement3.htm

Complement of a Set of SavedSearches: The complement of a set of SavedSearches, denoted diff,
is the set of all elements in the given universal set a that are not in b.
diff = {x ∈ a : x ∉ b}.
*/
func NotificationComplement(a, b []NotificationTemplate) (diff []NotificationTemplate) {

	if len(b) == 0 {
		return a
	}

	for i := range a {
		for j := range b {
			if a[i].Equals(b[j]) {
				break
			} else if j + 1 == len(b) {
				diff = append(diff,a[i])
			}
		}
	}
	return diff
}