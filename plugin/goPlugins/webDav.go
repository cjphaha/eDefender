package goplugin

import (
	"fmt"
	"github.com/cjphaha/eDefender/pkg/util"
	"github.com/cjphaha/eDefender/plugin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

type webDav struct {
	info   plugin.Plugin
	result []plugin.Plugin
}

func init() {
	plugin.Regist("web", &webDav{})
}
func (d *webDav) Init() plugin.Plugin {
	d.info = plugin.Plugin{
		Name:    "WebDav Put开启",
		Remarks: "开启了WebDav且配置不当导致攻击者可上传文件到web目录",
		Level:   1,
		Type:    "CONF",
		Author:  "cjp",
		References: plugin.References{
			KPID: "KP-0022",
		},
	}
	return d.info
}
func (d *webDav) GetResult() []plugin.Plugin {
	var result = d.result
	d.result = []plugin.Plugin{}
	return result
}
func (d *webDav) Check(URL string, meta plugin.TaskMeta) bool {
	fmt.Println(URL)
	putURL := URL + "/" + util.GetRandomString(6) + ".txt"
	request, err := http.NewRequest("PUT", putURL, strings.NewReader("vultest"))
	if err != nil {
		fmt.Println(err.Error())
		log.Error(err)
		return false
	}
	_, err = util.RequestDo(request, false, 15)
	if err != nil {
		fmt.Println(err.Error())
		log.Error(err)
		return false
	}
	vRequest, err := http.NewRequest("GET", putURL, nil)
	if err != nil {
		fmt.Println(err.Error())
		log.Error(err)
		return false
	}
	resp, err := util.RequestDo(vRequest, true, 15)
	if err != nil {
		fmt.Println(err.Error())
		log.Error(err)
		return false
	}
	if strings.Contains(resp.ResponseRaw, "vultest") {
		result := d.info
		result.Response = resp.ResponseRaw
		result.Request = resp.RequestRaw
		d.result = append(d.result, result)
		return true
	}
	return false
}

