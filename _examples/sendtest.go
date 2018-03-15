package main

import (
	"github.com/BurntSushi/toml"
)

func main() {
	var Cfg string
	if _, err := toml.DecodeFile("sms.toml", &Cfg); err != nil {
		return
	}
}
