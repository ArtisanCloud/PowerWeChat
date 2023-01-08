package support

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/go-playground/assert/v2"
	"testing"
)

// Authorization: WECHATPAY2-SHA256-RSA2048 mchid="1611854986",nonce_str="sacfXg6R9YwKqo3sfPoOVJQd4jPb0KOe",timestamp="1626623583",serial_no="2655A2CD634B06C2A86B28780228A997D047B01C",signature="dvB2r+z5v8JOsKw0goAmXNNTdLtpwmCkZHzuVdAi63kBZIvFugQuYx1nCUiZQckIV7ebq0JkK5/3N2gkmX70RJSnEa/Rjq7n2K//0OahXPGI+2qgFr5qUZ586en66QZjuQVeqoW6aYsaAwHPnszva56uJmopvHnuPPdUzKTTWf8sDNdIph/y+BpDXGTIvgifYR3RnJc2qh5n9eOo1Tqr4Ei6y6HhdPhMWMrr9RXY4bOCjtDkZhQ+mUXEP6aHLPau+5Th2cGlb5dyUY3o/MzgfjvvXjv4JDXhHFuo9BZAwp4XQcs/6jh/XAakf9lHx7ESvoQyT406Sfn30An3Y+p4wg=="
//func TestGetEncryptMethod(t *testing.T) {
//	signer := &SHA256WithRSASigner{
//		MchID:               "161186666",
//		CertificateSerialNo: "2355A2CD634B06C2A86B28780228A997D017B011",
//		PrivateKey:          getPrivateKey(),
//	}
//
//	signBody := "{\"amount\":{\"total\":1},\"appid\":\"ww16143ea0101327c7\",\"attach\":\"自定义数据说明\",\"description\":\"Image形象店-深圳腾大-QQ公仔\",\"mchid\":\"1611854986\",\"notify_url\":\"https://pay.wangchaoyi.com/power/notify\",\"out_trade_no\":\"5519778939773395659222199361\",\"payer\":{\"openid\":\"oAuaP0TRUMwP169nQfg7XCEAw3HQ\"}}"
//
//	authorization, err := GenerateSign(signer, GenerateSigner{
//		Method:       http.MethodPost,
//		CanonicalURL: "/v3/pay/transactions/jsapi",
//		SignBody:     signBody,
//		timestamp:    1626747079,
//		nonce:        "W2XLk2c8KYM1aRNBzwmeGBnVqZ3QbvHS",
//	})
//
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	expectAuthorization := "WECHATPAY2-SHA256-RSA2048 mchid=\"161186666\",nonce_str=\"W2XLk2c8KYM1aRNBzwmeGBnVqZ3QbvHS\",timestamp=\"1626747079\",serial_no=\"2355A2CD634B06C2A86B28780228A997D017B011\",signature=\"eUF/u5p9UjRtflGwOIk5jXUzD2Aduaj/NLhlwWwkliFPpF2I9jtTtM7gARrMHYuX2tQNS6OfY5Jf350D6OsJ4YKaKK4C8HOQ62maQ90DASJUrcqRI/EA4uyCrkqbUWnl+Xm2dE5wuVpfTSbRaMzOdXQwFB376uZgfUQtnD8C5PUfyqJ07qjxvh6NGi+R1vNCyG2rHhWityYtd66CZX4lBTOG5bJocn4GpOZOnyJO5/paQVQ8rKmQn+Wm7XCSCFL+4QLIa7ATyra/JMy0SLswq8ORfjCV0wFsHNR3h0u0vJo9JFRcqhhr/L6uQRc5x0vAC/wciOiDejAWYWBY90LEnA==\""
//
//	assert.Equal(t, authorization, expectAuthorization)
//}

func TestPaymentV2ParamsJoin(t *testing.T) {
	params := &power.StringMap{
		"appid":  "f323",
		"c":      "12",
		"d":      "34",
		"mch_id": "2323532",
	}
	key := "HelloPowerWeChat"
	expectText := "appid=f323&c=12&d=34&mch_id=2323532&key=HelloPowerWeChat"
	text := PaymentV2ParamsJoin(params, key)

	assert.Equal(t, expectText, text)
}

func TestGenerateSignMD5(t *testing.T) {
	params := &power.StringMap{
		"appid":  "f323",
		"c":      "12",
		"d":      "34",
		"mch_id": "2323532",
	}
	key := "HelloPowerWeChat"
	expectSignMD5 := "D18B5A9F4D01EB18CDEFFA39C78EB0F5"
	signMD5 := GenerateSignMD5(params, key)

	assert.Equal(t, expectSignMD5, signMD5)
}

func TestGenerateSignHmacSHA256(t *testing.T) {
	params := &power.StringMap{
		"appid":  "f323",
		"c":      "12",
		"d":      "34",
		"mch_id": "2323532",
	}
	key := "HelloPowerWeChat"
	expectSignHmacSHA256 := "CF3C3C7B038A12682967DC5ABADDAD56CE612FEE9B0E0A885B3B41E9E72B9A10"
	signHmacSHA256 := GenerateSignHmacSHA256(params, key)

	assert.Equal(t, expectSignHmacSHA256, signHmacSHA256)
}
