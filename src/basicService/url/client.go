package url

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/basicService/url/response"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	"math"
)

const DAY int = 86400

type Client struct {
	*kernel.BaseClient

	BaseUri string
}

func NewClient(app *kernel.ApplicationInterface) *Client {

	client := &Client{
		BaseClient: kernel.NewBaseClient(app, nil),
	}
	client.BaseClient.HttpRequest.BaseURI = "https://api.weixin.qq.com/"

	return client
}

// 短key托管
// https://developers.weixin.qq.com/doc/offiaccount/Account_Management/KEY_Shortener.html
func (comp *Client) ShortGenKey(longData string, expireSecond int) (*response.ResponseShortGenKey, error) {

	result := &response.ResponseShortGenKey{}

	params := &object.HashMap{
		"long_data":      longData,
		"expire_seconds": int(math.Min(float64(expireSecond), float64(30*DAY))),
	}
	_, err := comp.HttpPostJson("cgi-bin/shorten/gen", params, nil, nil, result)

	return result, err

}

// 获取端链接
// https://developers.weixin.qq.com/doc/offiaccount/Account_Management/KEY_Shortener.html
func (comp *Client) FetchShorten(shortKey string) (*response.ResponseFetchShorten, error) {

	result := &response.ResponseFetchShorten{}

	params := &object.HashMap{
		"short_key": shortKey,
	}
	_, err := comp.HttpPostJson("cgi-bin/shorten/fetch", params, nil, nil, result)

	return result, err

}
