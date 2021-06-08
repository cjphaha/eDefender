package service

import "errors"

type Config struct {
}

func (c *Config) ConfigValidate() error {
	if c == nil {
		return errors.New("log config is nil")
	}

	return nil
}
