package service

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"time"
)

type Service interface {
	GetInfo() (info Info)
}

type service struct {
	c *Config
}

func New(c *Config) (Service, error) {
	srv := &service{
		c: c,
	}
	return srv, nil
}

func (s *service) GetInfo() (info Info){
	info.Host = getHostInfo()
	//获取内存使用率 同时定时
	percent, _ := cpu.Percent(time.Second * 14, false)//设置间隔时间
	info.Cpu = getCpuInfo(fmt.Sprintf("%.2f",percent[0]))
	info.Mem = getMemInfo()
	return
}

