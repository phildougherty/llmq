// logger/logger.go
package logger

import (
    "github.com/sirupsen/logrus"
    "os"
)

func SetupLogger(level string) *logrus.Logger {
    log := logrus.New()
    log.Out = os.Stdout

    switch level {
    case "debug":
        log.SetLevel(logrus.DebugLevel)
    case "info":
        log.SetLevel(logrus.InfoLevel)
    case "warn":
        log.SetLevel(logrus.WarnLevel)
    case "error":
        log.SetLevel(logrus.ErrorLevel)
    default:
        log.SetLevel(logrus.InfoLevel)
    }

    return log
}
