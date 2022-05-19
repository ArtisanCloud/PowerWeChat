package url

import (
	"github.com/ArtisanCloud/PowerLibs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/basicService/url/response"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	"math"
)

const DAY int = 86400

type Client struct {
	*kernel.BaseClient

	BaseUri string
}

func NewClient(app *kernel.ApplicationInterface) *Client {

	return &Client{
		BaseClient: kernel.NewBaseClient(app, nil),

		BaseUri: "https://api.weixin.qq.com/",
	}

}

// 短key托管
// https://developers.weixin.qq.com/doc/offiaccount/Account_Management/KEY_Shortener.html
func (comp *Client) ShortGenKey(longData string, expireSecond int) (*response.ResponseShortGenKey, error) {

	result := &response.ResponseShortGenKey{}

	params := &object.HashMap{
		"long_data":longData,
		"expire_seconds": int(math.Min(float64(expireSecond), float64(30*DAY))),
	}
	_, err := comp.HttpPostJson("cgi-bin/shorten/gen", params, nil, nil, result)

	return result, err

}
