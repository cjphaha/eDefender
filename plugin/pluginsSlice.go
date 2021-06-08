package plugin

import (
	"strconv"
	"strings"
)

type pluginsSlice []map[string]interface{}

func (s pluginsSlice) Len() int      { return len(s) }
func (s pluginsSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s pluginsSlice) Less(i, j int) bool {
	ir := s[i]["references"].(map[string]interface{})
	jr := s[j]["references"].(map[string]interface{})
	return getKPINTID(ir["kpid"].(string)) < getKPINTID(jr["kpid"].(string))
}

func getKPINTID(kpid string) int {
	tmp1 := strings.Split(kpid, "-")
	if len(tmp1) != 2 {
		return 0
	}
	kpiniid, err := strconv.Atoi(tmp1[1])
	if err != nil {
		return 0
	}
	return kpiniid
}