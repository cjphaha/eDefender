package log

import (
	"errors"
	"os"
	"os/signal"
	"path"
	"strings"
	"syscall"

	"github.com/natefinch/lumberjack"
	log "github.com/sirupsen/logrus"
)

var logFilePath string

func initLogger(c *Config) {
	log.Info("init logger now!!!")
	logFilePath = path.Join(c.LogFilePath, "logs", c.LogFileName+".log")
	if c.DeployStage == "production" {
		log.SetFormatter(&log.JSONFormatter{})
	} else {
		log.SetFormatter(&log.TextFormatter{
			DisableColors: true,
		})
	}
	log.SetOutput(&lumberjack.Logger{
		Filename: logFilePath,
		MaxSize:  c.MaxSize,
		MaxAge:   c.MaxAge, //days
	})
	setLevel(c.Level)

	reloadLogFile()
	go registerRotate()
}

func registerRotate() {
	c := make(chan os.Signal)
	for {
		signal.Notify(c, syscall.SIGHUP)
		s := <-c
		if s == syscall.SIGHUP {
			reloadLogFile()
		}
	}
}

func reloadLogFile() {
	folder := path.Dir(logFilePath)
	err := os.MkdirAll(folder, os.ModePerm)
	if err != nil {
		log.WithError(err).Fatal("Fail to mk folder")
	}
	f, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		log.WithError(err).Fatal("Fail to open log file")
	}
	log.SetOutput(f)
}

func setLevel(level string) error {
	level = strings.ToLower(level)
	switch level {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	default:
		return errors.New("Unknown level")
	}
	return nil
}
