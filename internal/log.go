package tfa

import (
	"os"
	"runtime"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

// NewDefaultLogger creates a new logger based on the current configuration
func NewDefaultLogger() *logrus.Logger {
	// Setup logger
	log = logrus.StandardLogger()

	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)

	// Set logger format
	switch config.LogFormat {
	case "pretty":
		break
	case "json":
		log.SetFormatter(&logrus.JSONFormatter{})
	// "text" is the default
	default:
		log.SetFormatter(&logrus.TextFormatter{
			ForceColors:      true,
			FullTimestamp:    true,
			TimestampFormat:  "2006-01-02 15:04:05",
			CallerPrettyfier: func(f *runtime.Frame) (string, string) { return "", "" },
		})
	}

	// Set logger level
	switch config.LogLevel {
	case "trace":
		log.SetLevel(logrus.TraceLevel)
	case "debug":
		log.SetLevel(logrus.DebugLevel)
	case "info":
		log.SetLevel(logrus.InfoLevel)
	case "error":
		log.SetLevel(logrus.ErrorLevel)
	case "fatal":
		log.SetLevel(logrus.FatalLevel)
	case "panic":
		log.SetLevel(logrus.PanicLevel)
	// warn is the default
	default:
		log.SetLevel(logrus.WarnLevel)
	}

	return log
}
