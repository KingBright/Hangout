package turing

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"hacklife.fun/wechat/service/util"

	"github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/mp/message/callback/request"
	"github.com/chanxuehong/wechat/mp/message/callback/response"
)

const (
	TEXT_CODE     = 100000
	URL_CODE      = 200000
	NEWS_CODE     = 302000
	COOKBOOK_CODE = 308000

	ERROR0_CODE = 40000
	ERROR1_CODE = 40001
	ERROR2_CODE = 40002
	ERROR4_CODE = 40004
	ERROR7_CODE = 40007

	ERROR_MSG = "我坏掉了ヾ(;ﾟ;Д;ﾟ;)ﾉﾞ"
)

type Turing struct {
	Key string
	Url string
}

type Ask struct {
	Info   string `json:"info"`
	UserId string `json:"userid"`
	Loc    string `json:"loc"`
	Key    string `json:"key"`
}

// 40001	参数key错误
// 40002	请求内容info为空
// 40004	当天请求次数已使用完
// 40007	数据格式异常
// 100000	文本类
// 200000	链接类
// 302000	新闻类
// 308000	菜谱类
type CodeMessage struct {
	Code int `json:"code"`
}

type TextMessage struct {
	Text string `json:"text"`
}

type UrlMessage struct {
	Text string `json:"text"`
	Url  string `json:"url"`
}

type NewsMessage struct {
	Text string `json:"text"`
	Url  string `json:"url"`
}

type CookbookMessage struct {
	Text string `json:"text"`
	Url  string `json:"url"`
}

func New(url, key string) *Turing {
	var t Turing
	t.Key = key
	t.Url = url
	return &t
}

func (this Turing) Reply(ctx *core.Context) {
	msg := request.GetText(ctx.MixedMsg)
	from := msg.FromUserName
	to := msg.ToUserName

	log.Printf("收到:%s\n", msg.Content)
	ask := Ask{msg.Content, msg.FromUserName, "", this.Key}

	json := util.ToJsonBytes(ask)
	req, err := http.NewRequest("POST", this.Url, bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ctx.RawResponse(response.NewText(from, to, util.Now(), ERROR_MSG)) // 明文回复
		return
	}
	rep := genReply(from, to, string(body))
	log.Printf("回复:%s\n", util.ToJson(rep))
	ctx.RawResponse(rep) // 明文回复
}

// func genUrlReply(from, to, answer string) interface{} {
// 	return response.New
// }

func genReply(from, to, answer string) interface{} {
	log.Printf("Raw回复:\n%s\n", answer)
	decoder := json.NewDecoder(strings.NewReader(answer))
	message := new(CodeMessage)
	decoder.Decode(message)

	code := message.Code
	if code == TEXT_CODE {
		decoder = json.NewDecoder(strings.NewReader(answer))
		log.Printf("text")
		tm := new(TextMessage)
		decoder.Decode(tm)
		return response.NewText(from, to, util.Now(), tm.Text)
	} else if code == URL_CODE {
		log.Printf("url")
		um := new(UrlMessage)
		decoder.Decode(um)
		return response.NewText(from, to, util.Now(), um.Text)
	} else if code == NEWS_CODE {
		log.Printf("news")
	} else if code == COOKBOOK_CODE {
		log.Printf("cookbook")
	}
	return response.NewText(from, to, util.Now(), ERROR_MSG)
}
