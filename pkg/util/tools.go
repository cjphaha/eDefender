package util

import (
	"encoding/json"
	"math/rand"
	"time"
)

// Struct2Map 结构转map
func Struct2Map(obj interface{}) map[string]interface{} {
	jsonBytes, _ := json.Marshal(obj)
	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)
	return result
}

// GetRandomString 获取随机字符串
func GetRandomString(l int) string {
	return string(GetRandomBytes(l))
}

// GetRandomBytes 获取包含随机字母的 byte 数组
func GetRandomBytes(l int) []byte {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return result
}
