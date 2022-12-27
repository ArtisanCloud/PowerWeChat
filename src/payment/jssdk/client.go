package jssdk

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/basicService/jssdk"
	kernel2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	*jssdk.Client
}

func NewClient(app *kernel2.ApplicationInterface) (*Client, error) {
	jssdkClient, err := jssdk.NewClient(app)
	if err != nil {
		return nil, err
	}
	return &Client{
		jssdkClient,
	}, nil
}

// JSAPI调起支付API
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_4.shtml
func (comp *Client) BridgeConfig(prepayID string, isJson bool) (interface{}, error) {

	config := (*comp.BaseClient.App).GetConfig()
	appID := config.GetString("app_id", "")

	options := &object.StringMap{
		"appId":     appID,
		"timeStamp": fmt.Sprintf("%d", time.Now().Unix()),
		"nonceStr":  object.QuickRandom(32),
		"package":   fmt.Sprintf("prepay_id=%s", prepayID),
		"signType":  "RSA",
	}

	var err error
	(*options)["paySign"], err = comp.BaseClient.Signer.GenerateSign(fmt.Sprintf("%s\n%s\n%s\n%s\n",
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

func (comp *Client) SDKConfig(prepayID string) (*object.StringMap, error) {

	result, err := comp.BridgeConfig(prepayID, false)

	config := result.(*object.StringMap)
	(*config)["timestamp"] = (*config)["timeStamp"]
	delete(*config, "timeStamp")

	return config, err

}

func (comp *Client) AppConfig(prepayID string) (*object.StringMap, error) {

	config := (*comp.BaseClient.App).GetConfig()
	appID := config.GetString("app_id", "")
	mchID := config.GetString("mch_id", "")

	params := &object.StringMap{
		"appid":     appID,
		"partnerid": mchID,
		"prepayid":  prepayID,
		"noncestr":  object.UniqueID(""),
		"timestamp": fmt.Sprintf("%d", time.Now().Unix()),
		"package":   "Sign=WXPay",
	}

	//(*params)["sign"] = support.GenerateSign(params, config.GetString("key", ""), "RSA")

	return params, nil
}

func (comp *Client) ShareAddressConfig(request *http.Request, accessToken string, isJson bool) (interface{}, error) {

	config := (*comp.BaseClient.App).GetConfig()
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
		"url":         comp.GetUrl(request),
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
	(*params)["paySign"] = fmt.Sprintf("%x", hash.Sum(nil))
	if isJson {
		return json.Marshal(params)
	} else {
		return params, nil
	}

}

func (comp *Client) ContractConfig(params *object.StringMap) (*object.StringMap, error) {

	config := (*comp.BaseClient.App).GetConfig()
	(*params)["appid"] = config.GetString("app_id", "")
	(*params)["timestamp"] = fmt.Sprintf("%d", time.Now().Unix())
	//(*params)["sign"] = support.GenerateSign(params, config.GetString("key", ""), "md5")

	return params, nil

}
