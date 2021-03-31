package creator

import (
	"com/github/graylog-code/config/template"
	"com/github/graylog-code/graylog/api"
	"com/github/graylog-code/graylog/client"
	log "github.com/sirupsen/logrus"
	"strings"
)

func CreateShareApiObjectWithUserNames(config template.GraylogConfigTemplate,usernames []string) api.ShareApi {
	var share = api.ShareApi{}
	share.SelectedGranteeCapabilities = make(map[string]string)
	for i, username := range usernames {
		log.Debugf("Preparing %d. user: %s", i+1, username)
		userId := getGraylogUserIdByUsername(config,username)
		preparedUserId := prepareUserIdForShare(userId)
		share.SelectedGranteeCapabilities[preparedUserId] = "view"
	}
	return share
}

func getGraylogUserIdByUsername(config template.GraylogConfigTemplate, username string) string{
	var users api.UsersApi
	client.GetFromGraylog(config,"/users",&users)
	for i, user := range *users.Users {
		log.Debugf("Checking %d. stream",i + 1)
		if strings.EqualFold(user.Username, username) {
			log.Debugf("Found id(%s) for user: %s ",user.Id, username)
			return user.Id
		}
	}
	return ""
}

func prepareUserIdForShare(userId string) string {
	return "grn::::user:" + userId
}

func PrepareViewIdForShare(viewId string) string {
	return "grn::::search:" + viewId
}