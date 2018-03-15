package main

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/yunhor/alisms"
)

func main() {
	var cfg alisms.UserParams
	if _, err := toml.DecodeFile("sms.toml", &cfg); err != nil {
		return
	}
	fmt.Println(cfg.AccessKeyId)
	fmt.Println(cfg.AppSecret)
	cfg.TemplateParam = fmt.Sprintf("{\"name\":\"%s\",\"money\":\"%s\",\"time\":\"%s\"}", "中文名字", "8000", "1月20至2月22日")
	rt, str, err := alisms.SendMessage(&cfg)
	fmt.Println(rt)
	fmt.Println(str)
	if err != nil {
		fmt.Println(err.Error())
	}

}
