package logging

import (
	"log"
	"os"
)

var logger *log.Logger = nil
var warner *log.Logger = nil

func setLogger() {
	if logger == nil {
		logger = log.New(os.Stdout, "[LOG] ", log.LstdFlags)
	}
	if warner == nil {
		warner = log.New(os.Stdout, "[WARN] ", log.Llongfile)
	}
}

func Log(v ...interface{}) {
	setLogger()
	logger.Println(v...)
}

func Warn(v ...interface{}) {
	setLogger()
	warner.Println(v...)
}
