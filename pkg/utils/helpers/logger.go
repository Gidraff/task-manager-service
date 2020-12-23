package helpers

import (
	log "github.com/sirupsen/logrus"
	"os"
)
type Logger struct {
	Class string
}

func NewLogger(cls string) *Logger {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
	return &Logger{Class: cls}
}

func (lg *Logger) Info(msg string){
	log.WithFields(log.Fields{"class":lg.Class}).Info(msg)
}
func (lg *Logger) Error(msg string){
	log.WithFields(log.Fields{"class":lg.Class}).Error(msg)
}
func (lg *Logger) Warn(msg string){
	log.WithFields(log.Fields{
		"class":lg.Class,
	}).Warn(msg)
}



