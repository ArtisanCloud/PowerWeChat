package base

import (
	"github.com/ArtisanCloud/go-libs/http"
)

type Client struct {
	*http.BaseClient
}

func (comp *Client) GetCallbackIp() interface{} {
	return comp.HttpGet("cgi-bin/getcallbackip", nil)
}
