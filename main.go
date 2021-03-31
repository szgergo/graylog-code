package main

import (
	configHandler "com/github/graylog-code/config/handler"
	"com/github/graylog-code/config/template"
	"com/github/graylog-code/graylog/resource/orchestrator"
	"flag"
	log "github.com/sirupsen/logrus"
	"os"
)

var config template.GraylogConfigTemplate

func init() {
	configFilePath := flag.String("file", "", "The config util for graylog")
	accessToken := flag.String("access-token", "", "Access token for Graylog")
	flag.Parse()
	// Log as JSON instead of the default ASCII formatter.
	//log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
	config = configHandler.LoadConfiguration(configFilePath, accessToken)
}

func main() {
	orchestrator.StartManagers(config)
}




