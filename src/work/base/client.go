package base

import (
	"github.com/ArtisanCloud/go-wechat/src/kernel/http"
)

type Client struct {
	BaseClient *http.BaseClient
}

func (client *Client) GetCallbackIp() interface{} {
	return client.BaseClient.HttpGet("cgi-bin/getcallbackip", nil)
}
