package template

type GraylogConfigTemplate struct {
	GraylogServer          string                 `yaml:"graylog_server"`
	ApiPrefix 			   string                 `yaml:"api_prefix"`
	AccessToken            string                 `yaml:"access_token"`
	User string `yaml:"user"`
	RequiredGraylogVersion string                 `yaml:"required_graylog_version"`
	SavedSearches          []SavedSearchTemplate  `yaml:"saved_seaches,flow"`
	Inputs 				   []InputTemplate        `yaml:"inputs,flow"`
	Streams 			   []StreamTemplate 	  `yaml:"streams,flow"`
	Notifications 	       []NotificationTemplate `yaml:"notifications,flow"`
}

