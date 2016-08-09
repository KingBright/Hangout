package util

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func SetCookie(w http.ResponseWriter, name string, value string, expiration time.Time) {
	cookie := http.Cookie{Name: name, Value: value, Expires: expiration}
	http.SetCookie(w, &cookie)
}

func GetCookie(r *http.Request, name string) string {
	log.Println("try get session : ", name)
	cookie, _ := r.Cookie(name)
	if cookie != nil {
		log.Println("got session : ", cookie.Value)
		return cookie.Value
	}
	return ""
}

// Get post form param from Url
func GetPostFormParam(r *http.Request, paramName string) string {
	log.Println("try get post form param : ", paramName)
	return r.PostFormValue(paramName)
}

// Get param from Url
func GetUrlParam(r *http.Request, paramName string) string {
	log.Println("try get url param : ", paramName)
	return r.URL.Query().Get(paramName)
}

func Now() int64 {
	return time.Now().Unix()
}

func ToJsonBytes(msg interface{}) []byte {
	bytes, err := json.Marshal(msg)
	if err != nil {
		return nil
	}
	return bytes
}

func ToJson(msg interface{}) string {
	bytes, err := json.Marshal(msg)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func Month() time.Time {
	return time.Now().AddDate(0, 1, 0)
}

func Year() time.Time {
	return time.Now().AddDate(1, 0, 0)
}

func Week() time.Time {
	return time.Now().AddDate(0, 0, 7)
}
