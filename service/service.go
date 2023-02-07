package service

import (
	"errors"
	"log"
	"net/http"
	"net/url"

	"github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/mp/message/callback/request"
	"github.com/chanxuehong/wechat/mp/message/callback/response"
	"github.com/julienschmidt/httprouter"
	"hacklife.fun/wechat/service/constant"
	"hacklife.fun/wechat/service/model"
	"hacklife.fun/wechat/service/mwchain"
	"hacklife.fun/wechat/service/tpl"
	"hacklife.fun/wechat/service/util"
)

type HangoutService struct {
	Router *httprouter.Router
}

func New() *HangoutService {
	service := HangoutService{}
	service.Router = httprouter.New()
	return &service
}

func (this *HangoutService) Init() {
	this.GET(constant.PATH_TEST, mwchain.New(testHandler))

	this.GET(constant.PATH_HOME, mwchain.New(tokenHandler, authHandler, homeHandler))
	this.GET(constant.PATH_REGISTER, mwchain.New(tokenHandler, registerHandler))
	this.GET(constant.PATH_LIST, mwchain.New(tokenHandler, authHandler, listHandler))
	this.GET(constant.PATH_PUBLISH, mwchain.New(tokenHandler, authHandler, publishHandler))

	this.GET(constant.PATH_NOPWD, mwchain.New(noPwdHandler))
	this.GET(constant.PATH_HELP, mwchain.New(helpHandler))

	this.POST(constant.ACTION_DO_REGISTER, mwchain.New(tokenHandler, doRegisterHandler))
	this.POST(constant.ACTION_DO_PUBLISH, mwchain.New(tokenHandler, doPublishHandler))

	this.NotFound(http.FileServer(http.Dir(constant.PATH_STATIC)))
}

func (this *HangoutService) POST(path string, handler httprouter.Handle) {
	this.Router.POST(path, handler)
}

func (this *HangoutService) GET(path string, handler httprouter.Handle) {
	this.Router.GET(path, handler)
}

func (this *HangoutService) NotFound(handler http.Handler) {
	this.Router.NotFound = handler
}

func (this *HangoutService) HandleMsg(ctx *core.Context) bool {
	handle := false
	msg := request.GetText(ctx.MixedMsg)
	from := msg.FromUserName
	to := msg.ToUserName

	if msg.Content == constant.CMD_TEST {
		log.Println("handle cmd : ", constant.CMD_TEST)
		token := model.CreateToken(from)
		url, err := url.Parse(constant.BASE_URL + "?" + constant.TOKEN + "=" + token)
		if err == nil {
			ctx.RawResponse(response.NewText(from, to, util.Now(), url.String()))
			handle = true
		}
		return handle
	}
	return handle
}

func tokenHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	log.Println("token handler")
	if token, valid := model.CheckToken(w, r); valid {
		log.Println("put token into cookies")
		util.SetCookie(w, constant.TOKEN, token, util.Month())
		return nil
	} else {
		log.Println("token not valid")
		errorPage(w, r, tpl.ErrorMsg{Title: "出错啦", Message: "您的登录状态有问题。"})
		return errors.New("")
	}
}

func authHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	log.Println("auth handler")
	if tokenOk, login := model.CheckAuth(w, r); tokenOk && login {
		log.Println("everything ok")
		return nil
	} else {
		if tokenOk {
			log.Println("redirect to register")
			jump(w, r, tpl.JumpInfo{TargetName: "注册", TargetUrl: constant.PATH_REGISTER})
			return errors.New("")
		} else {
			log.Println("redirect to error page")
			errorPage(w, r, tpl.ErrorMsg{Title: "出错啦", Message: "登录超时了！请参考帮助页面刷新登录状态。"})
			return errors.New("")
		}
	}
}

func doRegisterHandler(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	log.Println("do login handler")
	name := util.GetPostFormParam(r, constant.NAME)
	email := util.GetPostFormParam(r, constant.EMAIL)
	token := util.GetCookie(r, constant.TOKEN)
	if err := model.CreateUserIfNotExist(name, email, token); err == nil {
		log.Println("user created or exists")
		jump(w, r, tpl.JumpInfo{TargetName: "您的首页", TargetUrl: constant.PATH_HOME})
	} else {
		log.Println("error occured")
		errorPage(w, r, tpl.ErrorMsg{Title: "出错啦", Message: "认证失败啦，无法注册！"})
	}
	return nil
}

func doPublishHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) error {
	errorPage(w, r, tpl.ErrorMsg{Title: "出错啦", Message: "还没有开发完成！"})
	return nil
}
func homeHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) error {
	log.Println("serve home : ", r.URL.String())
	tpl.Home(w, r)
	return nil
}

func jump(w http.ResponseWriter, r *http.Request, info tpl.JumpInfo) error {
	log.Println("serve jump : ", r.URL.String())
	tpl.Jump(w, r, info)
	return nil
}

func errorPage(w http.ResponseWriter, r *http.Request, err tpl.ErrorMsg) error {
	log.Println("serve error : ", r.URL.String())
	tpl.Error(w, r, err)
	return nil
}

func helpHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) error {
	log.Println("serve help : ", r.URL.String())
	tpl.Help(w, r)
	return nil
}

func noPwdHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) error {
	log.Println("serve nopwd: ", r.URL.String())
	tpl.NoPwd(w, r)
	return nil
}

func registerHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) error {
	log.Println("serve register: ", r.URL.String())
	token := util.GetCookie(r, constant.TOKEN)
	if model.CheckUserExist(token) {
		log.Println("user created or exists")
		jump(w, r, tpl.JumpInfo{TargetName: "您的首页", TargetUrl: constant.PATH_HOME})
	} else {
		tpl.Register(w, r)
	}
	return nil
}

func listHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) error {
	log.Println("serve register: ", r.URL.String())
	tpl.List(w, r)
	return nil
}

func publishHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) error {
	log.Println("serve publish: ", r.URL.String())
	tpl.Publish(w, r)
	return nil
}

func testHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) error {
	log.Println("serve publish: ", r.URL.String())
	tpl.Test(w, r)
	return nil
}
