package token

import (
	"encoding/base64"
	"encoding/json"
	"strings"
	"time"

	"hacklife.fun/wechat/util"
)

type Token struct {
	UserId   string `json:"user_id"`
	Expire   int64  `json:"expire"`
	Sequence int64  `json:"sequence"`
}

func New(id string) *Token {
	token := new(Token)
	token.UserId = id
	token.Expire = time.Now().AddDate(0, 1, 0).Unix()
	token.Sequence = 1
	return token
}

func (this *Token) Refresh() *Token {
	this.Sequence = this.Sequence + 1
	this.Expire = time.Now().AddDate(0, 1, 0).Unix()
	return this
}

func (this *Token) Encode() string {
	return base64.StdEncoding.EncodeToString([]byte(base64.StdEncoding.EncodeToString(util.ToJsonBytes(this))))
}

func Decode(str string) *Token {
	if result, err := base64.StdEncoding.DecodeString(str); err == nil {
		if result, err = base64.StdEncoding.DecodeString(string(result)); err == nil {
			decoder := json.NewDecoder(strings.NewReader(string(result)))
			token := new(Token)
			decoder.Decode(token)
			return token
		}
	}
	return nil
}
