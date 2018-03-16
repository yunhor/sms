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
		fmt.Println(cfg.AccessKeyId)
		fmt.Println(cfg.AppSecret)
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
