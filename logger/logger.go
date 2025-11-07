package logger

import (
	"fmt"
	"time"
)

type logger struct {
}

func NewLogger() *logger {
	return &logger{}
}

var Logger *logger

func init() {
	Logger = NewLogger()
}

func (l *logger) Info(message string) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "INFO:", message)
}

func (l *logger) Debug(message string) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "DEBUG:", message)
}

func (l *logger) Warn(message string) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "WARN:", message)
}

func (l *logger) Error(message string) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "ERROR:", message)
}
