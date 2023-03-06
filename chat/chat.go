package turing

import (
	"context"
	"log"
	"strings"
	"time"

	"hacklife.fun/wechat/service/util"

	"github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/mp/message/callback/request"
	"github.com/chanxuehong/wechat/mp/message/callback/response"
	cache "github.com/karlseguin/ccache"
	openai "github.com/sashabaranov/go-openai"
)

const (
	ERROR_MSG = "我坏掉了ヾ(;ﾟ;Д;ﾟ;)ﾉﾞ"
)

type Chat struct {
	cli   *openai.Client
	cache *cache.Cache
}

type ChatCache struct {
	Msg   *string
	Reply *string
}

func New(apiKey string) *Chat {
	var t Chat
	t.cli = openai.NewClient(apiKey)
	t.cache = cache.New(cache.Configure().MaxSize(100).ItemsToPrune(50))
	return &t
}

func (chat Chat) chat(ctx context.Context, req openai.ChatCompletionRequest) (resp openai.ChatCompletionResponse, err error) {
	rsp, err := chat.cli.CreateChatCompletion(ctx, req)
	return rsp, err
}

func (chat Chat) Reply(ctx *core.Context) {
	msg := request.GetText(ctx.MixedMsg)
	from := msg.FromUserName
	to := msg.ToUserName
	log.Printf("%s问: %s\n", msg.FromUserName, msg.Content)

	if strings.EqualFold(strings.ToLower(msg.Content), "retry") {
		// reply from cache
		result := chat.cache.Get(msg.FromUserName).Value()
		if result != nil {
			c := result.(ChatCache)
			if c.Reply != nil {
				log.Printf("reply: %s", *c.Reply)
				ctx.RawResponse(response.NewText(from, to, util.Now(), *c.Reply)) // 明文回复
				return
			}
		}
	}

	request := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{Role: "user", Content: msg.Content},
		},
		User: from,
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

		// put to cache
		chat.cache.Set(from, ChatCache{
			Msg:   &msg.Content,
			Reply: &str,
		}, time.Minute*2)
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

	// put to cache
	chat.cache.Set(from, ChatCache{
		Msg:   &msg.Content,
		Reply: &reply,
	}, time.Minute*2)
	ctx.RawResponse(response.NewText(from, to, util.Now(), reply)) // 明文回复

}
