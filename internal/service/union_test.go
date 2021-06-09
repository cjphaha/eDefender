package service

import (
	"fmt"
	"testing"
)

/*
info.Host = getHostInfo()
	//获取内存使用率 同时定时
	percent, _ := cpu.Percent(time.Second * 14, false)//设置间隔时间
	info.Cpu = getCpuInfo(fmt.Sprintf("%.2f",percent[0]))
	info.Mem = getMemInfo()
 */

func TestGetHost(t *testing.T) {
	Host := getHostInfo()
	fmt.Println(Host)
}