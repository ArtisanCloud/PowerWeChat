package support

import (
	"crypto/aes"
	"encoding/base64"
	"github.com/stretchr/testify/assert"
	"log"
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