package support

import (
	"context"
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/object"
	"time"
)

type GenerateSigner struct {
	Method       string // 接口提交方法。"GET", "POST"等
	CanonicalURL string // 微信支付接口路径。 例如： /v3/pay/transactions/jsapi
	SignBody     string // 提交的body字符串。 例如； {"amount":{"total":1},"appid":"ww16143ea0101327c7","attach":"自定义数据说明","description":"Image形象店-深圳腾大-QQ公仔","mchid":"1611854986","notify_url":"https://pay.wangchaoyi.com/wx/notify","out_trade_no":"5519778939773395659222199361","payer":{"openid":"oAuaP0TRUMwP169nQfg7XCEAw3HQ"}}
	timestamp    int64  // 单元测试传入的固定时间戳
	nonce        string // 单元测试传入的固定随机数
}

func GenerateSign(signer *SHA256WithRSASigner, gs GenerateSigner) (authorization string, err error) {

	timestamp := time.Now().Unix()
	nonce := object.QuickRandom(32)

	// Under ci mode, go fixed value
	// 在ci模式下面，走固定值
	if gs.timestamp != 0 && gs.nonce != "" {
		timestamp = gs.timestamp
		nonce = gs.nonce
	}

	// Splice the string to be signed
	// 拼接出需要签名的字符串
	message := fmt.Sprintf(SignatureMessageFormat, gs.Method, gs.CanonicalURL, fmt.Sprintf("%d", timestamp), nonce, gs.SignBody)

	signatureResult, err := signer.Sign(context.TODO(), message)
	if err != nil {
		return "", err
	}

	authorization = fmt.Sprintf(
		HeaderAuthorizationFormat,
		signer.GetAuthorizationType(),
		signatureResult.MchID,
		nonce,
		timestamp,
		signatureResult.CertificateSerialNo,
		signatureResult.Signature,
	)

	return authorization, err
}

func GetEncryptMethod(signType string, secretKey string) string {
	return ""
}
