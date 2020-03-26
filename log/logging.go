package log

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	. "os"
)

func InitializeFileLogging(logFile string) {

	var file, err = OpenFile(logFile, O_RDWR|O_CREATE|O_APPEND, 0666)
	if err != nil {
		fmt.Println("Could Not Open Log File : " + err.Error())
	}

	Formatter := new(log.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)
	log.SetOutput(file)
	//log.SetFormatter(&log.TextFormatter{})
	log.SetFormatter(&log.JSONFormatter{})

}
func InitializeLogging() {
	Formatter := new(log.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)
}
