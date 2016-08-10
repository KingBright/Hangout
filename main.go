package main

import (
	"log"
	"net/http"

	"./service"
	"./service/config"
	"./service/constant"
	"./turing"

	"github.com/chanxuehong/wechat.v2/mp/core"
	"github.com/chanxuehong/wechat.v2/mp/menu"
	"github.com/chanxuehong/wechat.v2/mp/message/callback/request"
	"github.com/chanxuehong/wechat.v2/mp/message/callback/response"
	"github.com/julienschmidt/httprouter"
)

var (
	msgHandler core.Handler
	msgServer  *core.Server
	robot      *turing.Turing

	tokenServer *core.DefaultAccessTokenServer

	hangoutService *service.HangoutService
)

func init() {
	config.Load("config.yml")

	mux := core.NewServeMux()
	mux.DefaultMsgHandleFunc(defaultMsgHandler)
	mux.DefaultEventHandleFunc(defaultEventHandler)
	mux.MsgHandleFunc(request.MsgTypeText, textMsgHandler)
	mux.EventHandleFunc(menu.EventTypeClick, menuClickEventHandler)
	msgHandler = mux

	msgServer = core.NewServer(config.WxOriId(), config.WxAppId(), config.WxToken(), config.WxEncodedAESKey(), msgHandler, nil)

	tokenServer = core.NewDefaultAccessTokenServer(config.WxAppId(), config.WxAppSecret(), nil)
	robot = turing.New(config.TuringApi(), config.TuringAppKey())

	hangoutService = service.New()
	hangoutService.Init()

	hangoutService.POST(constant.WEIXIN_CALLBACK, wxCallbackHandler)
}

func textMsgHandler(ctx *core.Context) {
	if !hangoutService.HandleMsg(ctx) {
		robot.Reply(ctx)
	}
	//ctx.AESResponse(resp, 0, "", nil) // aes密文回复
}

func defaultMsgHandler(ctx *core.Context) {
	log.Printf("收到消息:\n%s\n", ctx.MsgPlaintext)
	ctx.NoneResponse()
}

func menuClickEventHandler(ctx *core.Context) {
	log.Printf("收到菜单 click 事件:\n%s\n", ctx.MsgPlaintext)

	event := menu.GetClickEvent(ctx.MixedMsg)
	resp := response.NewText(event.FromUserName, event.ToUserName, event.CreateTime, "收到 click 类型的事件")
	ctx.RawResponse(resp) // 明文回复
	//ctx.AESResponse(resp, 0, "", nil) // aes密文回复
}

func defaultEventHandler(ctx *core.Context) {
	log.Printf("收到事件:\n%s\n", ctx.MsgPlaintext)
	ctx.NoneResponse()
}

func createMenu() {
	// wxClient := core.NewClient(tokenServer, nil)
	// hangoutBtn := menu.Button{menu.ButtonTypeText, "Hangout", "hangout", "", "hangout", nil}
	// menuToCreate := menu.Menu{[]menu.Button{hangoutBtn}, nil, 0}
	// err := menu.Create(wxClient, &menuToCreate)
	// if err != nil {
	// 	log.Printf("创建菜单失败%s", err.Error())
	// } else {
	// 	log.Printf("创建菜单%s\n", "View Hangout")
	// }
}

// wxCallbackHandler 是处理回调请求的 http handler.
//  1. 不同的 web 框架有不同的实现
//  2. 一般一个 handler 处理一个公众号的回调请求(当然也可以处理多个, 这里我只处理一个)
func wxCallbackHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	msgServer.ServeHTTP(w, r, nil)
}

func main() {
	createMenu()
	log.Println(http.ListenAndServe(":"+config.Port(), hangoutService.Router))
}
