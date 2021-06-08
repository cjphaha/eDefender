package config

import (
	"github.com/cjphaha/eDefender/internal/server"
	"github.com/cjphaha/eDefender/internal/service"
	"github.com/cjphaha/eDefender/pkg/log"
)

var (
	Conf     = new(Root)
	confPath string
)

// 配置文件根结点
type Root struct {
	Log     *log.Config
	Server *server.Config
	Service *service.Config
	Base *Base
}

type Base struct {
	Welecome string
}