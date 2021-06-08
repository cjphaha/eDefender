package server

import "errors"

type Config struct {
	Port   string
	IsCORS bool
}

func (c *Config) ConfigValidate() error {
	if c == nil {
		return errors.New("config is nil")
	}

	if c.Port == "" {
		c.Port = "8999"
	}

	return nil
}


