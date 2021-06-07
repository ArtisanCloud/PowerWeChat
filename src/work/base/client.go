package base

import (
	"github.com/ArtisanCloud/go-wechat/src/kernel/http"
)

type Client struct {
	BaseClient *http.BaseClient
}

func (comp *Client) GetCallbackIp() interface{} {
	return comp.BaseClient.HttpGet("cgi-bin/getcallbackip", nil)
}
