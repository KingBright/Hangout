package turing

import (
	"context"
	"log"
	"strings"

	"hacklife.fun/wechat/service/util"

	"github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/mp/message/callback/request"
	"github.com/chanxuehong/wechat/mp/message/callback/response"
	"github.com/otiai10/openaigo"
)

const (
	ERROR_MSG = "我坏掉了ヾ(;ﾟ;Д;ﾟ;)ﾉﾞ"
)

type Chat struct {
	client *openaigo.Client
}

func New(apiKey string) *Chat {
	var t Chat
	t.client = openaigo.NewClient(apiKey)
	return &t
}

func (chat Chat) chat(ctx context.Context, req openaigo.ChatCompletionRequestBody) (resp openaigo.ChatCompletionResponse, err error) {
	rsp, err := chat.client.Chat(ctx, req)
	return rsp, err
}

func (chat Chat) Reply(ctx *core.Context) {
	msg := request.GetText(ctx.MixedMsg)
	from := msg.FromUserName
	to := msg.ToUserName
	log.Printf("%s问: %s\n", msg.FromUserName, msg.Content)

	request := openaigo.ChatCompletionRequestBody{
		Model: "gpt-3.5-turbo",
		Messages: []openaigo.ChatMessage{
			{Role: "user", Content: msg.Content},
		},
	}
	rsp, err := chat.chat(context.Background(), request)

	log.Println(rsp, err)
	if err != nil {
		ctx.RawResponse(response.NewText(from, to, util.Now(), ERROR_MSG)) // 明文回复
		return
	}

	choices := rsp.Choices
	if len(choices) == 0 {
		ctx.RawResponse(response.NewText(from, to, util.Now(), ERROR_MSG)) // 明文回复
		return
	}

	if len(choices) == 1 {
		log.Printf("reply: %s", choices[0].Message.Content)
		str := strings.Trim(choices[0].Message.Content, " ")
		str = strings.ReplaceAll(str, "\n", "")
		ctx.RawResponse(response.NewText(from, to, util.Now(), str)) // 明文回复
		return
	}

	reply := "以下回答供您参考：\n"
	for i := 0; i < len(choices); i++ {
		str := strings.Trim(choices[i].Message.Content, " ")
		str = strings.ReplaceAll(str, "\n", "")
		reply += str + "\n"
	}
	log.Printf("reply: %s", reply)
	ctx.RawResponse(response.NewText(from, to, util.Now(), reply)) // 明文回复

}
