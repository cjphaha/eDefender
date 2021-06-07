package log

import (
	"errors"
	"log"
)

type Config struct {
	LogFilePath string
	LogFileName string
	Level       string
	DeployStage string
	MaxAge      int
	MaxSize     int
}

func (c *Config) ConfigValidate() error {
	if c == nil {
		return errors.New("log config is nil")
	}

	if c.LogFilePath == "" {
		c.LogFilePath = "log"
	}

	if c.LogFileName == "" {
		c.LogFileName = "toy_vote.log"
	}

	if c.DeployStage == "" {
		c.DeployStage = "develop"
	}

	if c.MaxAge == 0 {
		c.MaxAge = 3
	}

	if c.MaxSize == 0 {
		c.MaxSize = 5
	}

	return nil
}

func InitLog(c *Config) (err error) {
	// check config
	err = c.ConfigValidate()
	if err != nil {
		log.Fatal(err)
		return
	}
	initLogger(c)
	//registerControl(c)
	return
}
