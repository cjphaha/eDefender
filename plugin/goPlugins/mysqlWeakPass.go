package goplugin

import (
	"database/sql"
	"fmt"
	"github.com/cjphaha/eDefender/plugin"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

type mysqlWeakPass struct {
	info   plugin.Plugin
	result []plugin.Plugin
}

func init() {
	plugin.Regist("mysql", &mysqlWeakPass{})
}
func (d *mysqlWeakPass) Init() plugin.Plugin {
	d.info = plugin.Plugin{
		Name:    "MySQL 弱口令",
		Remarks: "导致数据库敏感信息泄露，严重可导致服务器直接被入侵控制。",
		Level:   0,
		Type:    "WEAKPWD",
		Author: "cjp",
		References: plugin.References{
			URL:  "https://www.cnblogs.com/yunsicai/p/4080864.html",
			KPID: "KP-0005",
		},
	}
	return d.info
}
func (d *mysqlWeakPass) GetResult() []plugin.Plugin {
	var result = d.result
	d.result = []plugin.Plugin{}
	return result
}
func (d *mysqlWeakPass) Check(netloc string, meta plugin.TaskMeta) (b bool) {
	if strings.IndexAny(netloc, "http") == 0 {
		return
	}
	userList := []string{
		"root", "www", "bbs", "web", "admin",
	}
	for _, user := range userList {
		for _, pass := range meta.PassList {
			pass = strings.Replace(pass, "{user}", user, -1)
			connStr := fmt.Sprintf("%s:%s@tcp(%s)/?timeout=%ds", user, pass, netloc, 15)
			db, err := sql.Open("mysql", connStr)
			if err != nil {
				fmt.Println(err.Error())
				break
			}
			err = db.Ping()
			if err == nil {
				db.Close()
				result := d.info
				result.Request = connStr
				result.Remarks = fmt.Sprintf("弱口令：%s,%s,%s", user, pass, result.Remarks)
				d.result = append(d.result, result)
				b = true
				break
			} else if strings.Contains(err.Error(), "Access denied") {
				continue
			} else {
				return
			}
		}
	}
	return b
}

