package virtualPayment

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram/virtualPayment/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/device/response"
)

// 小程序虚拟支付
// uri固定写死为 requestVirtualPayment
// 官方示例是写死为这个postbody 没有任何相关文档表示这个postbody是怎么来的 以及各个参数的意义，只能直译
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/virtual-payment.html#_2-4-%E7%AD%BE%E5%90%8D%E8%AF%A6%E8%A7%A3

func (client *Client) TransactionVirtual(ctx context.Context, in *request.VirtualPaymentOrderRequest) (result *object.StringMap, err error) {

	if client.offerId == "" {

		return nil, errors.New("offerId is empty")
	}

	if client.appKey == "" {

		return nil, errors.New("appKey is empty")
	}

	if in.SessionKey == "" {

		return nil, errors.New("sessionKey is empty")
	}

	uri := "requestVirtualPayment"

	postBody := fmt.Sprintf(`{"offerId":"%s","buyQuantity":1,"env":0,"currencyType":"CNY","platform":"android","productId":"%d","goodsPrice":%d,"outTradeNo":"%s","attach":"%s"}`, client.offerId, in.ProductId, in.Price, in.OutTradeNo, in.Attach)

	paySign := kernel.CalcPaySig(uri, postBody, client.appKey)
	signature := kernel.CalcSignature(postBody, in.SessionKey)

	// todo 是否要变为强类型
	return &object.StringMap{
		"post_body": postBody,
		"pay_sign":  paySign,
		"signature": signature,
	}, nil

}

// 商品上传
// https://developers.weixin.qq.com/miniprogram/dev/platform-capabilities/industry/virtual-payment.html#_2-3-%E6%9C%8D%E5%8A%A1%E5%99%A8API

func (client *Client) StartUploadGoods(ctx context.Context, params *request.UploadProductsRequest) (result *response.BaseResp, err error) {

	fmt.Printf("appid: %s, app_key: %s, offer_id: %s \n", (*client.App).GetConfig().GetString("app_id", ""), client.appKey, client.offerId)

	if err != nil {

		return nil, err
	}

	postBody, err := json.Marshal(params)

	if err != nil {

		return nil, err
	}

	signPost := string(postBody)

	// paySign
	endpoint := "/xpay/start_upload_goods"
	paySign := kernel.CalcPaySig(endpoint, signPost, client.appKey)

	uri := fmt.Sprintf("%s?pay_sig=%s", endpoint, paySign)

	_, err = client.HttpPostJson(ctx, uri, signPost, nil, nil, &result)

	return result, err
}
