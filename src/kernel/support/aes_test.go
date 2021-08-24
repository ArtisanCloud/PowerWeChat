package support

import (
	"crypto/aes"
	"encoding/base64"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"log"
	"os"
	"testing"
)

const (
	encodingAesKey = "CyjBvGm4tJNmcThkeL5Wajzom4jKlBygKtpTkYeqsnE"
)

func TestAES_Decrypt(t *testing.T) {
	aesKey, err := base64.StdEncoding.DecodeString(encodingAesKey + "=")
	if err != nil {
		log.Fatal(err)
	}

	// original ciphertext content
	// 原始密文内容
	ciphertext := "ZBY4h9SD1Y/v+ighMGU4yHSVWyMuf2UdQQfFuB8S/SY5il9MynR1rtKIJDG871vOqaCs0FVje6bX0VyhvL02PG7FXjRR+rAyiwVv7T7nBGgamkPeOv4n/jXgFQnY6NEt3Zym09t4aI+21pDcE6uN6JzcC+hxw5sBfVyza3NHyuLiyhvalsFbasA4HNcnK+4/3eZLMo+e6yvHUsY4WhmZWW5u5mHcDZm1jtI2fUtzod6Sx7PrCOhm+zUO1I98OD0EmICxHHMjMkGmSFuGhG8vXRKaRavzvV/fkscZK1HjW+XJOV0c9wYDZT+/LTjDS/yx5KkM5lIksiw57DsCro9gvmE1WdvB3bw3m+5ie2x1ZRSuh6ik1rbMTiXgJXt0MnbV8VsyPe1zhkbug1cwdhT9zcOTieb0oyiIl3wSAEkNNBNBaWB+78V1hM/ngMoZgKaaasJUCGtPH/rN+YnxGicDOg=="

	newAes := NewAES()
	iv := aesKey[:aes.BlockSize]
	msg, encryptError := newAes.Decrypt(ciphertext, aesKey, iv)
	if encryptError != nil {
		log.Fatal(encryptError)
	}

	assert.Contains(t, string(msg), "<xml><ToUserName><![CDATA[wwedab3d5fc43636d8]]></ToUserName><FromUserName><![CDATA[WangChaoYi]]></FromUserName><CreateTime>1623301925</CreateTime><MsgType><![CDATA[text]]></MsgType><Content><![CDATA[密文测试]]></Content><MsgId>6972028683187026694</MsgId><AgentID>1000005</AgentID></xml>")
}

func TestAES_Encrypt(t *testing.T) {
	aesKey, err := base64.StdEncoding.DecodeString(encodingAesKey + "=")
	if err != nil {
		log.Fatal(err)
	}

	plainText := "<xml><ToUserName><![CDATA[wwedab3d5fc43636d8]]></ToUserName><FromUserName><![CDATA[WangChaoYi]]></FromUserName><CreateTime>1623301925</CreateTime><MsgType><![CDATA[text]]></MsgType><Content><![CDATA[密文测试]]></Content><MsgId>6972028683187026694</MsgId><AgentID>1000005</AgentID></xml>"

	newAES := NewAES()
	iv := aesKey[:aes.BlockSize]
	ciphertext, encryptError := newAES.Encrypt([]byte(plainText), aesKey, iv)
	if encryptError != nil {
		log.Fatal(encryptError)
	}

	plainText2, encryptError := newAES.Decrypt(string(ciphertext), aesKey, iv)
	if encryptError != nil {
		log.Fatal(encryptError)
	}

	assert.Contains(t, string(plainText2), plainText)
}


func TestEncryptor_PKCS7Padding(t *testing.T) {
	newAES := NewAES()

	assert.Equal(t, newAES.PKCS7Padding([]byte("aaaaa"), 6), []byte{97, 97, 97, 97, 97, 1})
	assert.Equal(t, newAES.PKCS7Padding([]byte("aaaaa"), 7), []byte{97, 97, 97, 97, 97, 2, 2})
	assert.Equal(t, newAES.PKCS7Padding([]byte("aaaaa"), 20), []byte{97, 97, 97, 97, 97, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15, 15})
}

func TestEncryptor_PKCS7UnPadding(t *testing.T) {
	newAES := NewAES()

	text, encryptError := newAES.PKCS7UnPadding([]byte{97, 97, 97, 97, 97, 1})
	if encryptError != nil {
		log.Fatal(encryptError)
	}
	assert.Equal(t, text, []byte{97, 97, 97, 97, 97})

}


const (
	testAESUtilAPIV3Key       = "testAPIv3Key0000"
	testAESUtilCiphertext     = "FsdXzxryWfKwvLJKf8LG/ToRPTRh8RN9wROC"
	testAESUtilPlaintext      = "Hello World"
	testAESUtilNonce          = "wCq51qZv4Yfg"
	testAESUtilAssociatedData = "Dl56FrqWQJF1t9LtC3vEUsniZKvbqdR8"
)

// 解析微信支付回调数据
// 2021/07/27 10:29:42 {"id":"c61507ac-d975-56c9-8ab3-d9c0ca38743f","create_time":"2021-07-27T10:29:42+08:00","resource_type":"encrypt-resource","event_type":"TRANSACTION.SUCCESS","summary":"支付成功","resource":{"original_type":"transaction","algorithm":"AEAD_AES_256_GCM","ciphertext":"TdSY7TauzzJoVzvL3hven8NN/ENrjkhzBZXpCV/hEmRltFltHokAi/G52wGWHDvq63qjutx1BN9VKWPKICMpJQVuAQZ7ktx9YPIyOLL1UzfDkbcZBdMnBoEbR0m0/nbUrWjPLBYlDYzkAnLo1e3+4DXG86ewMCJZwbHmOp+ihlBWFgyh/1dvFatIM0BPI7dMjF4xg2Rtn2ibgeaZ/GE8Yjo5NHDws000uymScvOqHNcJPcNsv5IrJz2HmztxV2Hp+wjaWFAFvv1EHHJ+r4+YsFw45yR2jaehzERf1HB6MxUkouJMY5ZSnxTKNFMOnSAioBtGV1Wg2ZQomNo4i60vdR5ePoGQUNrgmTDelhBE58ihZu3uBtUMHkufyXo2NpIhd8pSw3lpLj5P5zkQn1UBYvB/92WBBAhl4/72BU+JdfVHpk+NOwHC3UFi6I3l1fxrT66dVHSVIFdbS32YFt5VvXrdwt0N186dtLg2/sw2SwAn+vsXSjyzIHgbxPpdSwISyUFLE1NrE3kkJUwH1extkclsWPhYNmz8fyBPD9vNegktszcH+Mgohl13MSGy3tcU0OjkcRDEX8Rhw+EaD2gc8lFdkrjqptrhDPXB/SzWzUVr9ycEVg==","associated_data":"transaction","nonce":"LUgxtPpJynB6"}}
// [GIN] 2021/07/27 - 10:29:42 | 200 |      81.333µs |   121.51.58.172 | POST     "/wx/notify"
func TestDecryptAES256GCM(t *testing.T) {

	ciphertext := "TdSY7TauzzJoVzvL3hven8NN/ENrjkhzBZXpCV/hEmRltFltHokAi/G52wGWHDvq63qjutx1BN9VKWPKICMpJQVuAQZ7ktx9YPIyOLL1UzfDkbcZBdMnBoEbR0m0/nbUrWjPLBYlDYzkAnLo1e3+4DXG86ewMCJZwbHmOp+ihlBWFgyh/1dvFatIM0BPI7dMjF4xg2Rtn2ibgeaZ/GE8Yjo5NHDws000uymScvOqHNcJPcNsv5IrJz2HmztxV2Hp+wjaWFAFvv1EHHJ+r4+YsFw45yR2jaehzERf1HB6MxUkouJMY5ZSnxTKNFMOnSAioBtGV1Wg2ZQomNo4i60vdR5ePoGQUNrgmTDelhBE58ihZu3uBtUMHkufyXo2NpIhd8pSw3lpLj5P5zkQn1UBYvB/92WBBAhl4/72BU+JdfVHpk+NOwHC3UFi6I3l1fxrT66dVHSVIFdbS32YFt5VvXrdwt0N186dtLg2/sw2SwAn+vsXSjyzIHgbxPpdSwISyUFLE1NrE3kkJUwH1extkclsWPhYNmz8fyBPD9vNegktszcH+Mgohl13MSGy3tcU0OjkcRDEX8Rhw+EaD2gc8lFdkrjqptrhDPXB/SzWzUVr9ycEVg=="
	nonce := "LUgxtPpJynB6"
	associatedData := "transaction"
	apiV3Key := os.Getenv("API_V3_KEY")
	//log.Println("apiV3Key: ", apiV3Key)

	if apiV3Key == "" {
		panic(errors.New("API_V3_KEY not found from env"))
	}

	plaintext, err := DecryptAES256GCM(apiV3Key, associatedData, nonce, ciphertext)
	if err != nil {
		panic(err)
		return
	}

	assert.Equal(t, plaintext, `{"mchid":"1611854986","appid":"ww16143ea0101327c7","out_trade_no":"1826186758589177228934882556","transaction_id":"4200001136202107273023033988","trade_type":"JSAPI","trade_state":"SUCCESS","trade_state_desc":"支付成功","bank_type":"BOSH_CREDIT","attach":"自定义数据说明","success_time":"2021-07-27T10:29:42+08:00","payer":{"openid":"oAuaP0TRUMwP169nQfg7XCEAw3HQ"},"amount":{"total":1,"payer_total":1,"currency":"CNY","payer_currency":"CNY"}}`)
	//log.Printf("Plaintext: %s\n", string(plaintext))
}


func TestDecryptAes256Gcm(t *testing.T) {
	type args struct {
		apiv3Key       string
		associatedData string
		nonce          string
		ciphertext     string
	}
	tests := []struct {
		name      string
		args      args
		plaintext string
		wantErr   bool
	}{
		{
			name: "decrypt success",
			args: args{
				apiv3Key:       testAESUtilAPIV3Key,
				associatedData: testAESUtilAssociatedData,
				nonce:          testAESUtilNonce,
				ciphertext:     testAESUtilCiphertext,
			},
			wantErr:   false,
			plaintext: testAESUtilPlaintext,
		},
		{
			name: "invalid base64 ciphertext",
			args: args{
				apiv3Key:       testAESUtilAPIV3Key,
				associatedData: testAESUtilAssociatedData,
				nonce:          testAESUtilNonce,
				ciphertext:     "invalid cipher",
			},
			wantErr: true,
		},
		{
			name: "invalid ciphertext",
			args: args{
				apiv3Key:       testAESUtilAPIV3Key,
				associatedData: testAESUtilAssociatedData,
				nonce:          testAESUtilNonce,
				ciphertext:     "SGVsbG8gV29ybGQK",
			},
			wantErr: true,
		},
		{
			name: "invalid aes key",
			args: args{
				apiv3Key:       "not a aes key",
				associatedData: testAESUtilAssociatedData,
				nonce:          testAESUtilNonce,
				ciphertext:     testAESUtilCiphertext,
			},
			wantErr: true,
		},
		{
			name: "wrong aes key",
			args: args{
				apiv3Key:       "testAPIv3Key1111",
				associatedData: testAESUtilAssociatedData,
				nonce:          testAESUtilNonce,
				ciphertext:     testAESUtilCiphertext,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				plaintext, err := DecryptAES256GCM(
					tt.args.apiv3Key, tt.args.associatedData, tt.args.nonce, tt.args.ciphertext,
				)
				require.Equal(t, tt.wantErr, err != nil)
				assert.Equal(t, tt.plaintext, plaintext)
			},
		)
	}
}