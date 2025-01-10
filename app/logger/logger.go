package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

type Logger struct {
	*logrus.Logger
}

func NewLogger() *Logger {
	log := logrus.New()
	log.SetReportCaller(true)

	// Custom formatter with colors
	log.Formatter = &logrus.TextFormatter{
		ForceColors:   true,
		DisableColors: false,
		// CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
		// 	filename := path.Base(frame.File)
		// 	return fmt.Sprintf("%s()", frame.Function), fmt.Sprintf("%s:%d", filename, frame.Line)
		// },
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	}
	// log.ReportCaller = false

	// Log to file
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	mw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(mw)

	// Set log level
	log.SetLevel(logrus.DebugLevel)

	return &Logger{log}
}
