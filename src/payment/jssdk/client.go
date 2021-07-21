package jssdk

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-libs/str"
	"github.com/ArtisanCloud/power-wechat/src/basicService/jssdk"
	"github.com/ArtisanCloud/power-wechat/src/kernel/support"
	"github.com/ArtisanCloud/power-wechat/src/payment/kernel"
	"net/url"
	"time"
)

type Client struct {
	*jssdk.Client
}

func NewClient(app *kernel.ApplicationPaymentInterface) *Client {
	return &Client{
		jssdk.NewClient(app),
	}
}

// JSAPI调起支付API
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_4.shtml
func (comp *Client) BridgeConfig(prepayId string, isJson bool) (interface{}, error) {

	config := (*comp.App).GetConfig()
	appID := config.GetString("app_id", "")

	options := &object.StringMap{
		"appId":     appID,
		"timeStamp": fmt.Sprintf("%d", time.Now().Unix()),
		"nonceStr":  str.UniqueID(""),
		"package":   fmt.Sprintf("prepay_id=%s", prepayId),
		"signType":  "RSA",
	}

	signBody, err := object.JsonEncode(options)
	if err != nil {
		return nil, err
	}
	(*options)["paySign"], err = support.GenerateSign(comp.Signer, support.GenerateSigner{
		Method:       "POST",
		CanonicalURL: "/v3/pay/transactions/jsapi",
		SignBody:     signBody,
	})
	if isJson {
		return json.Marshal(options)
	} else {
		return options, nil
	}

}

func (comp *Client) SDKConfig(prepayId string) (*object.StringMap, error) {

	result, err := comp.BridgeConfig(prepayId, false)

	config := result.(*object.StringMap)
	(*config)["timestamp"] = (*config)["timeStamp"]
	delete(*config, "timeStamp")

	return config, err

}

func (comp *Client) AppConfig(prepayId string) (*object.StringMap, error) {

	config := (*comp.App).GetConfig()
	appID := config.GetString("app_id", "")
	mchID := config.GetString("mch_id", "")

	params := &object.StringMap{
		"appid":     appID,
		"partnerid": mchID,
		"prepayid":  prepayId,
		"noncestr":  str.UniqueID(""),
		"timestamp": fmt.Sprintf("%d", time.Now().Unix()),
		"package":   "Sign=WXPay",
	}

	//(*params)["sign"] = support.GenerateSign(params, config.GetString("key", ""), "RSA")

	return params, nil
}

func (comp *Client) ShareAddressConfig(accessToken string, isJson bool) (interface{}, error) {

	config := (*comp.App).GetConfig()
	appID := config.GetString("app_id", "")

	params := &object.StringMap{
		"appId":     appID,
		"scope":     "jsapi_address",
		"timeStamp": fmt.Sprintf("%d", time.Now().Unix()),
		"nonceStr":  str.UniqueID(""),
		"signType":  "SHA1",
	}

	signParams := &object.StringMap{
		"appid":       (*params)["appId"],
		"url":         comp.GetUrl(),
		"timestamp":   (*params)["timeStamp"],
		"noncestr":    (*params)["nonceStr"],
		"accesstoken": accessToken,
	}

	strJoinedSignParams := object.GetJoinedWithKSort(signParams)
	u, err := url.Parse(strJoinedSignParams)
	if err != nil {
		return nil, err
	}
	strJoinedSignParams = u.EscapedPath()

	hash := sha1.New()
	hash.Write([]byte(strJoinedSignParams))
	(*params)["paySign"] = string(hash.Sum(nil))
	if isJson {
		return json.Marshal(params)
	} else {
		return params, nil
	}

}

func (comp *Client) ContractConfig(params *object.StringMap) (*object.StringMap, error) {

	config := (*comp.App).GetConfig()
	(*params)["appid"] = config.GetString("app_id", "")
	(*params)["timestamp"] = fmt.Sprintf("%d", time.Now().Unix())
	//(*params)["sign"] = support.GenerateSign(params, config.GetString("key", ""), "md5")

	return params, nil

}
