package logger

import (
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.InfoLevel)

	formatter := &log.TextFormatter{
		DisableLevelTruncation: true,
	}
	log.SetFormatter(formatter)
}
