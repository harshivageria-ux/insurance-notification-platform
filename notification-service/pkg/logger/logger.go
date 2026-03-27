package logger

import (
	"log"
	"os"
)

var logger *log.Logger

func init() {
	logger = log.New(os.Stdout, "[NOTIFICATION_SERVICE] ", log.LstdFlags|log.Lshortfile)
}

func Info(msg string) {
	logger.Println("INFO:", msg)
}

func Error(msg string) {
	logger.Println("ERROR:", msg)
}

func Debug(msg string) {
	logger.Println("DEBUG:", msg)
}

func Warn(msg string) {
	logger.Println("WARN:", msg)
}
