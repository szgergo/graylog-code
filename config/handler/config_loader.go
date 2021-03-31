package handler

import (
	"com/github/graylog-code/config/template"
	"com/github/graylog-code/error"
	errorTypes "com/github/graylog-code/error/types"
	"com/github/graylog-code/util"
)

func LoadConfiguration(configFilePath *string, accessToken *string) template.GraylogConfigTemplate {
	if len(*configFilePath) <= 0 {
		error.ThrowError(errorTypes.BadUsageError{})
	}
	config := Parse(util.LoadConfigurationFile(*configFilePath))
	if len(config.AccessToken) == 0 && len(*accessToken) > 0 {
		config.AccessToken = *accessToken
	}

	if len(config.AccessToken) == 0 {
		error.ThrowError(errorTypes.NoAccessTokenError{})
	}
	return config
}
