package config

import (
	"github.com/easy-project-templete/pkg/log"
)

var (
	Conf     = new(Root)
	confPath string
)

// 配置文件根结点
type Root struct {
	Log     *log.Config
}