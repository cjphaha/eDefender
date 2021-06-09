package goplugin

import (
	"fmt"
	"github.com/cjphaha/eDefender/pkg/util"
	"github.com/cjphaha/eDefender/plugin"
	"strings"
)

type zookeeperUnauth struct {
	info   plugin.Plugin
	result []plugin.Plugin
}

func init() {
	plugin.Regist("zookeeper", &zookeeperUnauth{})
}
func (d *zookeeperUnauth) Init() plugin.Plugin {
	d.info = plugin.Plugin{
		Name:    "zookeeper 未授权访问",
		Remarks: "导致敏感信息泄露。",
		Level:   2,
		Type:    "UNAUTH",
		Author:  "cjp",
		References: plugin.References{
			KPID: "KP-0029",
		},
	}
	return d.info
}
func (d *zookeeperUnauth) GetResult() []plugin.Plugin {
	var result = d.result
	d.result = []plugin.Plugin{}
	return result
}
func (d *zookeeperUnauth) Check(netloc string, meta plugin.TaskMeta) bool {
	buf, err := util.TCPSend(netloc, []byte("envi"), 15)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(buf))
	if err == nil && strings.Contains(string(buf), "Environment") {
		result := d.info
		result.Request = fmt.Sprintf("zookeeper://%s", netloc)
		result.Response = string(buf)
		result.Remarks = fmt.Sprintf("未授权访问，%s", result.Remarks)
		d.result = append(d.result, result)
		fmt.Println(d.result)
		return true
	}

	return false
}

