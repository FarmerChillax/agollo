package agollo

import (
	"strconv"
	"time"
)

type signature struct {
	AppID           string `json:"appId"`
	AccesskeySecret string `json:"accesskey_secret"`
}

func newSignature(appId, accesskeySecret string) *signature {
	return &signature{AppID: appId, AccesskeySecret: accesskeySecret}
}

func (s *signature) getAuthorization(url, timestamp string) string {
	return s.AccesskeySecret
}

func (s *signature) getTimestamp() string {
	t := time.Now().UnixNano() / 1e6
	return strconv.Itoa(int(t))
}
