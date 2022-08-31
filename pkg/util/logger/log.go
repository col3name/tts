package logger

import log "github.com/sirupsen/logrus"

func LogFatalIfNeed(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
