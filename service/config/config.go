package config

import (
	"fmt"
	"path/filepath"

	config "github.com/go-ozzo/ozzo-config"
)

var c *config.Config

func Port() int {
	return c.GetInt("Port", 8787)
}

func WxAppId() string {
	return c.GetString("Weixin.AppId", "")
}

func WxAppSecret() string {
	return c.GetString("Weixin.AppSecret", "")
}
func WxOriId() string {
	return c.GetString("Weixin.OriId", "")
}
func WxToken() string {
	return c.GetString("Weixin.Token", "")
}
func WxEncodedAESKey() string {
	return c.GetString("Weixin.EncodedAESKey", "")
}

func ChatApiKey() string {
	return c.GetString("Chat.ApiKey", "")
}

func init() {
	c = config.New()
}

func Load(configFile string) {
	path, err := filepath.Abs(configFile)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(path)
	}
	c.Load(configFile)
}
