package client

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"log"
	"time"
)

const authPayload = `{
	"userName" : "%s",
	"password" : "%s"
}`

const expiredDateLayout = "2006-01-02 15:04:05"

const expireOffsetInSeconds = 30

type Auth struct {
	Token  string
	Expiry time.Time
	offset int64
}

type AuthResponse struct {
	Data      interface{} `json:"data,omitempty"`
	ErrorCode string      `json:"errcode,omitempty"`
	ErrMsg    string      `json:"errmsg,omitempty"`
}

func (t *Auth) estimateExpireTime() int64 {
	return time.Now().Unix() + t.offset
}

func (t *Auth) CaclulateOffset() {
	t.offset = expireOffsetInSeconds
}

func (au *Auth) IsValid() bool {
	if au.Token != "" && au.Expiry.Unix() > au.estimateExpireTime() {
		return true
	}
	return false
}

func (client *Client) InjectAuthenticationHeader(req *resty.Request) error {
	log.Printf("[DEBUG] Begin Injection")
	client.l.Lock()
	defer client.l.Unlock()
	if client.password != "" {
		if client.AuthToken == nil || !client.AuthToken.IsValid() {
			err := client.Authenticate()
			if err != nil {
				return err
			}
		}
		req.Header.Add("X-ACCESS-TOKEN", client.AuthToken.Token)
		return nil
	}
	return fmt.Errorf("password is missing")
}
