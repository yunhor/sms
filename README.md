## 阿里短消息服务

### 由来

   - 阿里大于合并至阿里云后，摇身一变成了阿里云通信，多了许多新功能，有了新版的api,虽然旧版的仍旧可用。但毕竟不是长久之计。  
   - 官方没有golang sdk
   - 已有好几个go库，但都几乎只实现发送的功能。

### feature
   - golang
   - [x] 阿里大于 2.0版 备份
   - [x] 现版本基于阿里云协议2017-05-25版
   - [x] 客户端调用支持命令行与toml参数配置
   - [x] 短信发送
   - [ ] 验证码接口   
   - [ ] 群发接口
   - [ ] 发送反馈查询接口
   - [ ] 其它业务查询接口
   - [ ] 语音功能接口
   - [ ] 流量功能接口
   - [ ] 物联网功能接口
   - [ ] 私密小号功能接口

### 参考
 
   - [阿里云官网](https://dayu.aliyun.com/?spm=a3142.10677814.0.0.23716ebcliNC2w)

     - [阿里大于api](http://open.taobao.com/doc2/apiDetail.htm?spm=0.0.0.0.bkKKhG&apiId=25450)
     - [云通信短信api](https://help.aliyun.com/document_detail/56189.html?spm=a2c4g.11186623.6.580.bDKh92)

   - [holdno/alidayu](https://github.com/holdno/alidayu)
   - [ltt1987/alidayu](https://github.com/ltt1987/alidayu)
   - [gwpp/alidayu-go](https://github.com/gwpp/alidayu-go)

### 安装

    go get github.com/yunhor/alisms

### 示例
#### 命令行/toml配置，单条发送

```
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/urfave/cli"
	"github.com/yunhor/alisms"
)

func main() {

	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: "sms.toml",
			Usage: "从配置文件中读取配置信息 `FILE`",
		},
	}
	app.Action = func(c *cli.Context) error {
		var cfg alisms.UserParams
		if _, err := toml.DecodeFile(c.String("config"), &cfg); err != nil {
			log.Fatal(err)
		}
		cfg.TemplateParam = fmt.Sprintf("{\"name\":\"%s\",\"money\":\"%s\",\"time\":\"%s\"}", "中文名字", "8000", "1月20至2月22日")
		//模板其它参数修改
		rt, str, err := alisms.SendMessage(&cfg)
		fmt.Println(rt)
		fmt.Println(str)
		if err != nil {
			log.Fatal(err)
		}
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
```
#### 更多例子详见_examples目录

### 杂记

- 2018-03-15 旧的大于的accesskey和secret并不相同，迁移至阿里云后需要重新生成一对。

### 阿里大于2.0

#### 安装

	go get github.com/yunhor/alisms/dayu
	
#### 下列参数写入相应的toml

	HTTPSURL = "https://eco.taobao.com/router/rest"
	HTTPURL = "http://gw.api.taobao.com/router/rest"
	sendSms = "alibaba.aliqin.fc.sms.num.send"
	callTTS = "alibaba.aliqin.fc.tts.num.singlecall"
	callVoice = "alibaba.aliqin.fc.voice.num.singlecall"
	callDouble = "alibaba.aliqin.fc.voice.num.doublecall"
	msgConsume = "taobao.tmc.messages.consume"
	msgConfirm = "taobao.tmc.messages.confirm"