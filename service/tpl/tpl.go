package tpl

import (
	"net/http"
	"text/template"
)

type JumpInfo struct {
	TargetName string
	TargetUrl  string
}

type ErrorMsg struct {
	Title   string
	Message string
}

func init() {

}

func Home(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("service/tpl/index.html"))
	t.Execute(w, nil)
}

func Jump(w http.ResponseWriter, r *http.Request, info JumpInfo) {
	t := template.Must(template.ParseFiles("service/tpl/jump.html"))
	t.Execute(w, info)
}

func Error(w http.ResponseWriter, r *http.Request, err ErrorMsg) {
	t := template.Must(template.ParseFiles("service/tpl/error.html"))
	t.Execute(w, err)
}

func Help(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("service/tpl/help.html"))
	t.Execute(w, nil)
}

func Register(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("service/tpl/register.html"))
	t.Execute(w, nil)
}

func NoPwd(w http.ResponseWriter, r *http.Request) {
	template.Must(template.ParseFiles("service/tpl/nopwd.html")).Execute(w, nil)
}

func Publish(w http.ResponseWriter, r *http.Request) {
	template.Must(template.ParseFiles("service/tpl/publish.html")).Execute(w, nil)
}

func List(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("service/tpl/list.html"))
	t.Execute(w, nil)
}

func Test(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("service/tpl/test.html"))
	t.Execute(w, nil)
}
