package util

import (
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httputil"
	"time"

	log "github.com/sirupsen/logrus"
)

var client *http.Client

// Resp 封装的http返回包
type Resp struct {
	Body        []byte
	Other       *http.Response
	RequestRaw  string
	ResponseRaw string
}

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

// RequestDo 发送指定的request，返回结果结构，hasRaw参数决定是否返回原始请求包和返回包内容
func RequestDo(request *http.Request, hasRaw bool, timeout int) (Resp, error) {
	var result Resp
	var err error
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.113 Safari/537.36")
	if hasRaw {
		requestOut, err := httputil.DumpRequestOut(request, true)
		if err == nil {
			result.RequestRaw = string(requestOut)
		}
	}
	log.Info("request do", request.URL.String())
	client.Timeout = time.Second * time.Duration(timeout)
	result.Other, err = client.Do(request)
	if err != nil {
		log.Error(err.Error())
		return result, err
	}
	log.Info("response code:", result.Other.StatusCode, "len:", result.Other.ContentLength)
	defer result.Other.Body.Close()
	if hasRaw {
		ResponseOut, err := httputil.DumpResponse(result.Other, true)
		if err == nil {
			result.ResponseRaw = string(ResponseOut)
		}
	}
	result.Body, _ = ioutil.ReadAll(result.Other.Body)
	return result, err
}
