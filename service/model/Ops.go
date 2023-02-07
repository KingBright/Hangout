package model

import (
	"encoding/base64"
	"errors"
	"log"
	"net/http"
	"time"

	"hacklife.fun/wechat/service/constant"
	"hacklife.fun/wechat/service/token"
	"hacklife.fun/wechat/util"
)

// Create a token
func CreateToken(id string) string {
	encodedId := base64.StdEncoding.EncodeToString([]byte(id))

	tk := new(Token)
	if has, err := engine.Table("token").Where(constant.USER_ID+" = ?", encodedId).Get(tk); err == nil && has {
		log.Println("refresh token")
		tk.Token = token.Decode(tk.Token).Refresh().Encode()
		engine.Update(tk)
	} else {
		log.Println("create new token : ", err)
		tk.UserId = encodedId
		tk.Token = token.New(encodedId).Encode()
		engine.Insert(tk)
	}
	return tk.Token
}

func CheckToken(w http.ResponseWriter, r *http.Request) (string, bool) {
	if tokenStr := util.GetUrlParam(r, constant.TOKEN); tokenStr != "" {
		log.Println("try validate token from url")
		return validToken(tokenStr)
	} else if tokenStr = util.GetCookie(r, constant.TOKEN); tokenStr != "" {
		log.Println("try validate token from cookie")
		return validToken(tokenStr)
	} else {
		log.Println("no token found from request")
		return "", false
	}
}

func validToken(tokenStr string) (string, bool) {
	token1 := token.Decode(tokenStr)
	if token1 == nil {
		return "", false
	}
	tk := new(Token)
	if has, err := engine.Table("token").Where(constant.USER_ID+" = ?", token1.UserId).Get(tk); err == nil && has {
		token2 := token.Decode(tk.Token)
		if token1.Expire > time.Now().Unix() {
			if token1.Sequence == token2.Sequence {
				log.Println("token valid")
				return tokenStr, true
			} else {
				log.Println("token sequence not right : ", token1.Sequence, token2.Sequence)
				return tokenStr, false
			}
		} else {
			log.Println("token expired : ", token1.Expire, token2.Expire)
			return tokenStr, false
		}
	} else {
		log.Println("no token found for id : " + token1.UserId)
		return "", false
	}
}

func CheckUserExist(tokenStr string) bool {
	tk := token.Decode(tokenStr)
	return tk != nil && checkUserExist(tk.UserId)
}

// Check if a user exists
func checkUserExist(id string) bool {
	user := new(User)
	if has, err := engine.Table("user").Where(constant.USER_ID+" = ?", id).Get(user); err == nil && has {
		return true
	}
	return false
}

func CreateUserIfNotExist(name string, email string, tokenStr string) error {
	tk := token.Decode(tokenStr)
	if tk != nil && checkUserExist(tk.UserId) {
		return nil
	} else {
		return insertUser(name, email, tk.UserId)
	}
}

// Create a user
func insertUser(name string, email string, id string) error {
	user := new(User)
	user.Name = name
	user.Email = email
	user.UserId = id

	if _, err := engine.InsertOne(user); err == nil {
		log.Println("successfully created user")
		return nil
	} else {
		log.Println("failed to create user", err)
		return errors.New("Can't create your profile. Please contact the administrator.")
	}
}

func CheckAuth(w http.ResponseWriter, r *http.Request) (tokenValid bool, logined bool) {
	tokenStr := ""
	if tokenStr = util.GetUrlParam(r, constant.TOKEN); tokenStr == "" {
		log.Println("token not found in url")
		if tokenStr = util.GetCookie(r, constant.TOKEN); tokenStr == "" {
			log.Println("token not found in cookies")
			return false, false
		}
	}

	if tk := token.Decode(tokenStr); tk == nil {
		log.Println("Can't decode to token")
		return false, false
	} else {
		if has := checkUserExist(tk.UserId); has {
			log.Println("User exists")
			return true, true
		} else {
			log.Println("Token is ok, but user info not exist")
			return true, false
		}
	}

}
