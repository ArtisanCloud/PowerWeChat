package redpack

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	payment "github.com/ArtisanCloud/power-wechat/src/payment/kernel"
	"net/http"
)

type Client struct {
	*payment.BaseClient
}

func NewClient(app *payment.ApplicationPaymentInterface) *Client {
	return &Client{
		payment.NewBaseClient(app),
	}
}

// Send Miniprogram Normal redpack.
// https://pay.weixin.qq.com/wiki/doc/api/tools/cash_coupon.php?chapter=18_2&index=2
func (comp *Client) SendMiniProgramNormal(params *power.HashMap) (interface{}, error) {
	config := (*comp.App).GetConfig()

	externalRequest := (*comp.App).GetComponent("ExternalRequest").(*http.Request)
	clientIP := externalRequest.Host
	if (*params)["client_ip"] != nil && (*params)["client_ip"].(string) != "" {
		clientIP = (*params)["client_ip"].(string)
	}
	base := &object.HashMap{
		"total_num": 1,
		"client_ip": clientIP,
		"wxappid":   config.GetString("app_id", ""),
	}

	options := object.MergeHashMap(base, params.ToHashMap())

	endpoint := comp.Wrap("/mmpaymkttransfers/sendminiprogramhb")
	result, err := comp.SafeRequest(endpoint, nil, "POST", options, false, nil)

	return result, err
}

