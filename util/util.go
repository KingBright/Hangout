package util

import (
	"encoding/json"
	"time"
)

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
