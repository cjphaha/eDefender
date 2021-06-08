package plugin

import (
	"github.com/cjphaha/eDefender/pkg/util"
	log "github.com/sirupsen/logrus"
	"sort"
	"strings"
)

// GoPlugins GO插件集
var GoPlugins map[string][]GoPlugin

//References 插件附加信息
type References struct {
	URL  string `json:"url"`
	CVE  string `json:"cve"`
	KPID string `json:"kpid"`
}

func init() {
	GoPlugins = make(map[string][]GoPlugin)
}

// Plugin 漏洞插件信息
type Plugin struct {
	Name       string     `json:"name"`
	Remarks    string     `json:"remarks"`
	Level      int        `json:"level"`
	Type       string     `json:"type"`
	Author     string     `json:"author"`
	References References `json:"references"`
	Request    string
	Response   string
}

// GetPlugins 获取插件信息
func GetPlugins() (plugins []map[string]interface{}) {
	for name, pluginList := range GoPlugins {
		for _, plugin := range pluginList {
			info := plugin.Init()
			pluginMap := util.Struct2Map(info)
			delete(pluginMap, "request")
			delete(pluginMap, "response")
			pluginMap["target"] = name
			plugins = append(plugins, pluginMap)
		}
	}
	sort.Stable(pluginsSlice(plugins))
	return plugins
}

// 格式检查
func formatCheck(task Task) bool {
	if strings.Contains(strings.ToLower(task.Netloc), string([]byte{103, 111, 118, 46, 99, 110})) {
		return false
	}
	if task.Type == "web" {
		if strings.IndexAny(task.Netloc, "http") != 0 {
			return false
		}
	} else if strings.IndexAny(task.Netloc, "http") == 0 {
		return false
	}
	return true
}

// Scan 开始插件扫描
func Scan(task Task) (result []map[string]interface{}) {
	log.Info("new task:", task)
	if ok := formatCheck(task); ok == false {
		return
	}
	log.Info("go plugin total:", len(GoPlugins))
	// 正式开始扫描
	for n, pluginList := range GoPlugins {
		if strings.Contains(strings.ToLower(task.Target), "cve-") {
			for _, plugin := range pluginList {
				pluginInfo := plugin.Init()
				if strings.ToLower(pluginInfo.References.CVE) != strings.ToLower(task.Target) {
					continue
				}
				log.Info("run plugin:", pluginInfo.References.CVE, pluginInfo.Name)
				resultList := pluginRun(task, plugin)
				result = append(result, resultList...)
				break
			}
		} else if strings.Contains(strings.ToLower(task.Target), "kp-") {
			for _, plugin := range pluginList {
				pluginInfo := plugin.Init()
				if strings.ToLower(pluginInfo.References.KPID) != strings.ToLower(task.Target) {
					continue
				}
				log.Info("run plugin:", pluginInfo.References.KPID, pluginInfo.Name)
				resultList := pluginRun(task, plugin)
				result = append(result, resultList...)
				break
			}
		} else if strings.Contains(strings.ToLower(task.Target), strings.ToLower(n)) || task.Target == "all" {
			for _, plugin := range pluginList {
				pluginInfo := plugin.Init()
				log.Info("run plugin:", task.Target, pluginInfo.Name)
				resultList := pluginRun(task, plugin)
				result = append(result, resultList...)
			}
		}
	}
	if task.Type == "service" {
		return result
	}
	return result
}

func pluginRun(taskInfo Task, plugin GoPlugin) (result []map[string]interface{}) {
	var hasVul bool
	try(func() {
		hasVul = plugin.Check(taskInfo.Netloc, taskInfo.Meta)
	}, func(e interface{}) {
		log.Println("panic", e)
	})
	if hasVul == false {
		return
	}
	for _, res := range plugin.GetResult() {
		log.Info("hit plugin:", res.Name)
		result = append(result, util.Struct2Map(res))
	}
	return result
}

func try(fun func(), handler func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			handler(err)
		}
	}()
	fun()
}