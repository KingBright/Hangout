package config

import (
	"fmt"

	"github.com/jinzhu/configor"
)

var Config = struct {
	Port string `default:"8787"`

	WX struct {
		AppId         string
		AppSecret     string
		OriId         string
		Token         string
		EncodedAESKey string
	}

	Turing struct {
		Api    string
		AppKey string
	}
}{}

func Load(configFilePath string) {
	configor.Load(&Config, configFilePath)
	fmt.Printf("config: %#v", Config)
}
