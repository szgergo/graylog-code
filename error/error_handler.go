package error

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func Check(e error) {
	if e != nil {
		ThrowError(e.Error())
	}
}

func ThrowError(args... interface{}) {
	log.Error(args)
	os.Exit(1)
}
