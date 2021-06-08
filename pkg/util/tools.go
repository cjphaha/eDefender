package util

import "encoding/json"

// Struct2Map 结构转map
func Struct2Map(obj interface{}) map[string]interface{} {
	jsonBytes, _ := json.Marshal(obj)
	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)
	return result
}
