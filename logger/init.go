package logger

import (
	"os"
	"time"
)

func NowDate() string {
	return time.Now().Format("2006-01-02")
}

func Init() (*os.File, error) {
	fileName := "logs/" + NowDate() + ".log"
	logFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	return logFile, err
}
