package jssdk

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/basicService/jssdk"
	kernel2 "github.com/ArtisanCloud/PowerWeChat/src/kernel"
	"net/url"
	"time"
)

type Client struct {
	*jssdk.Client
}

func NewClient(app *kernel2.ApplicationInterface) *Client {
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
		"nonceStr":  object.QuickRandom(32),
		"package":   fmt.Sprintf("prepay_id=%s", prepayId),
		"signType":  "RSA",
	}

	var err error
	(*options)["paySign"], err = comp.Signer.GenerateSign(fmt.Sprintf("%s\n%s\n%s\n%s\n",
		(*options)["appId"],
		(*options)["timeStamp"],
		(*options)["nonceStr"],
		(*options)["package"],
	))
	if err != nil {
		return nil, err
	}
	if !isJson {
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
		"noncestr":  object.UniqueID(""),
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
		"nonceStr":  object.UniqueID(""),
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
