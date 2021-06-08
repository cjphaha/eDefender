package util

import (
	"net"
	"time"

	log "github.com/sirupsen/logrus"
)

// TCPSend 指定目标发送tcp报文，返回结果（仅适用于一次交互即可判断漏洞的场景）
func TCPSend(netloc string, data []byte, timeout int) ([]byte, error) {
	conn, err := net.DialTimeout("tcp", netloc, time.Second*time.Duration(timeout))
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	log.Info("tcp send", len(data))
	_, err = conn.Write(data)
	if err != nil {
		return nil, err
	}
	buf := make([]byte, 20480)
	n, err := conn.Read(buf)
	if err != nil {
		return nil, err
	}
	log.Info("tcp recv", n)
	return buf[:n], nil
}
