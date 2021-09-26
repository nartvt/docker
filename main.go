package main

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetReportCaller(true)
	logger := logrus.New()
	buffer := &bytes.Buffer{}
	logger.Out = buffer
	logger.SetLevel(logrus.DebugLevel)
	logger.SetReportCaller(true)
	logger.WithTime(time.Now())

	formatter := &logrus.TextFormatter{
		TimestampFormat:  "2006-01-02 15:04:05",
		CallerPrettyfier: runTime,
	}
	enry := &logrus.Entry{}
	enry.Logger = logger
	formatter.Format(enry)
	logrus.SetFormatter(formatter)
}

func main() {

	logrus.Info("Hello Walrus after")
	logrus.Warn("Hello Walrus after")
}

func fileLine(toFile string) string {
	arr := strings.Split(toFile, "/")
	return arr[len(arr)-1]
}

func runTime(f *runtime.Frame) (string, string) {
	return "", fmt.Sprintf("%s - line %d", fileLine(f.File), f.Line)
}

// func init() {
// 	logrus.SetReportCaller(true)

// 	logger := logrus.New()
// 	buffer := &bytes.Buffer{}
// 	logger.Out = buffer
// 	logger.SetLevel(logrus.DebugLevel)
// 	logger.SetReportCaller(true)
// 	logger.WithTime(time.Now())

// 	formatter := &prefixed.TextFormatter{
// 		TimestampFormat: "2006-01-02 15:04:05",
// 		FullTimestamp:   true,
// 	}
// 	// formatLogrus := &logrus.TextFormatter{
// 	// 	TimestampFormat: "2006-01-02 15:04:05",
// 	// 	FullTimestamp:   true,
// 	// 	FieldMap: logrus.FieldMap{
// 	// 		logrus.FieldKeyLevel: "@level",
// 	// 		logrus.FieldKeyTime:  "@timestamp",
// 	// 		logrus.FieldKeyMsg:   "@message"},
// 	// 	DisableLevelTruncation: true,
// 	// 	CallerPrettyfier: func(f *runtime.Frame) (string, string) {
// 	// 		return "", fmt.Sprintf("\t%s - line %d", formatFilePath(f.File), f.Line)
// 	// 	},
// 	// }
// 	enry := &logrus.Entry{}
// 	enry.Logger = logger

// 	logrus.SetFormatter(formatter)
// }
// func main() {

// 	logrus.Info("Hello Walrus after")
// 	logrus.Warn("Hello Walrus after")
// }
