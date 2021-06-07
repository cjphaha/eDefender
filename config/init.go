package config

import (
	"flag"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

//AppConfig 全局可访问配置文件对象
var (
	AppConfig *Root
)

const (
	name     = "setting"
	fileType = "yaml"
)

//InitConfig 初始化Config读取
//InitConfig 初始化Config读取
func init() {
	configPath := flag.String("config_path", ".", "config path")
	flag.Parse()
	v := viper.New()
	v.AddConfigPath(*configPath)
	v.SetConfigName(name)
	v.SetConfigType(fileType)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.ReadInConfig(); err != nil {
			panic(err)
		}
	})

	if err := v.Unmarshal(&AppConfig); err != nil {
		fmt.Println(err)
	}
}
