package kernel

import (
	"encoding/xml"
	"github.com/go-playground/assert/v2"
	"log"
	"testing"
)

const (
	token          = "pcoLldqMu2WQQtYBeDX5HiXPsBFlTnDM"
	encodingAesKey = "CyjBvGm4tJNmcThkeL5Wajzom4jKlBygKtpTkYeqsnE"
	appID          = "1000005"
)

func TestEncryptor_Decrypt(t *testing.T) {
	msgSignature := "d728d2bea7e28585ebd14eacbd330d84073fe7f1"
	timestamp := "1623301925"
	nonce := "1624217036"

	// original xml ciphertext content
	// 微信发过来的xml原始密文内容
	ciphertext := "<xml><ToUserName><![CDATA[wwedab3d5fc43636d8]]></ToUserName><Encrypt><![CDATA[ZBY4h9SD1Y/v+ighMGU4yHSVWyMuf2UdQQfFuB8S/SY5il9MynR1rtKIJDG871vOqaCs0FVje6bX0VyhvL02PG7FXjRR+rAyiwVv7T7nBGgamkPeOv4n/jXgFQnY6NEt3Zym09t4aI+21pDcE6uN6JzcC+hxw5sBfVyza3NHyuLiyhvalsFbasA4HNcnK+4/3eZLMo+e6yvHUsY4WhmZWW5u5mHcDZm1jtI2fUtzod6Sx7PrCOhm+zUO1I98OD0EmICxHHMjMkGmSFuGhG8vXRKaRavzvV/fkscZK1HjW+XJOV0c9wYDZT+/LTjDS/yx5KkM5lIksiw57DsCro9gvmE1WdvB3bw3m+5ie2x1ZRSuh6ik1rbMTiXgJXt0MnbV8VsyPe1zhkbug1cwdhT9zcOTieb0oyiIl3wSAEkNNBNBaWB+78V1hM/ngMoZgKaaasJUCGtPH/rN+YnxGicDOg==]]></Encrypt><AgentID><![CDATA[1000005]]></AgentID></xml>"

	// Expected content after decryption
	// 解密之后的内容
	decryptionText := "<xml><ToUserName><![CDATA[wwedab3d5fc43636d8]]></ToUserName><FromUserName><![CDATA[WangChaoYi]]></FromUserName><CreateTime>1623301925</CreateTime><MsgType><![CDATA[text]]></MsgType><Content><![CDATA[密文测试]]></Content><MsgId>6972028683187026694</MsgId><AgentID>1000005</AgentID></xml>"

	encryptor, err := NewEncryptor(appID, token, encodingAesKey)
	if err != nil {
		log.Fatal("NewEncryptor init fail: ", err)
	}

	// Start parsing the ciphertext
	// 开始解析密文内容
	msg, encryptError := encryptor.Decrypt([]byte(ciphertext), msgSignature, nonce, timestamp)

	if encryptError != nil {
		log.Fatal("encryptor.Decrypt error: ", encryptError)
	}

	// The inferred ciphertext is parsed and matches the expected result
	// 推断密文解析出来之后和预期结果一致
	assert.Equal(t, string(msg), decryptionText)
}

func TestEncryptor_Encrypt(t *testing.T) {
	timestamp := "1623301925"
	nonce := "1624217036"
	//timestamp := ""
	//nonce := ""

	// Original plain text
	// 原始明文消息
	plainText := "<xml><ToUserName><![CDATA[wwedab3d5fc43636d8]]></ToUserName><FromUserName><![CDATA[WangChaoYi]]></FromUserName><CreateTime>1623301925</CreateTime><MsgType><![CDATA[text]]></MsgType><Content><![CDATA[密文测试]]></Content><MsgId>6972028683187026694</MsgId><AgentID>1000005</AgentID></xml>"

	encryptor, err := NewEncryptor("1000005", token, encodingAesKey)
	if err != nil {
		log.Fatal(err)
	}

	// Execute encryption
	// 执行加密
	encryptMsg, encryptError := encryptor.Encrypt(plainText, nonce, timestamp)
	if encryptError != nil {
		log.Fatal(encryptError)
	}

	// Because the parsed cipher text is in xml format, it is parsed back again.
	// 因为解析出来的密文是xml格式，所以重新再解析回来。
	var msgBiz WeComReplyMsg
	err = xml.Unmarshal(encryptMsg, &msgBiz)
	if err != nil {
		log.Fatal(err)
	}
	//log.Println("msgBiz: ", msgBiz.Signature, msgBiz.Nonce, msgBiz.Timestamp, msgBiz.Encrypt)

	// Simulate a raw cipher uniformMessage sent from WeChat
	// 模拟一个微信发送过来的原始密文消息
	msgRecv := &WeComRecvMsg{
		ToUserName: "232323",
		Encrypt:    msgBiz.Encrypt.Value,
		AgentID:    appID,
	}
	msgRecvXmlStr, err := xml.Marshal(msgRecv)
	if err != nil {
		log.Fatal(err)
	}

	// Perform verification signature and decryption again
	// 再次执行验证签名和解密
	msg, encryptError := encryptor.Decrypt(msgRecvXmlStr, msgBiz.Signature.Value, msgBiz.Nonce.Value, msgBiz.Timestamp)

	// The timestamp and nonce passed in should not change
	// 传入的时间戳和随机数应该没有变化
	assert.Equal(t, msgBiz.Timestamp, timestamp)
	assert.Equal(t, msgBiz.Nonce.Value, nonce)

	// Encryption => Decryption The returned uniformMessage should be the same
	// 加密 => 解密 回来的消息应该一致
	assert.Equal(t, string(msg), plainText)
}

func TestEncryptor_GetToken(t *testing.T) {
	encryptor, err := NewEncryptor("1000005", token, encodingAesKey)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, encryptor.GetToken(), token)
}

func TestEncryptor_randString(t *testing.T) {
	encryptor, err := NewEncryptor("1000005", token, encodingAesKey)
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, len(encryptor.randString(10)), 10)
	assert.Equal(t, len(encryptor.randString(16)), 16)
	assert.Equal(t, len(encryptor.randString(32)), 32)
}

func TestEncryptor_Signature(t *testing.T) {
	msgSignature := "d728d2bea7e28585ebd14eacbd330d84073fe7f1"
	timestamp := "1623301925"
	nonce := "1624217036"

	// original ciphertext content
	// 原始密文内容
	ciphertext := "ZBY4h9SD1Y/v+ighMGU4yHSVWyMuf2UdQQfFuB8S/SY5il9MynR1rtKIJDG871vOqaCs0FVje6bX0VyhvL02PG7FXjRR+rAyiwVv7T7nBGgamkPeOv4n/jXgFQnY6NEt3Zym09t4aI+21pDcE6uN6JzcC+hxw5sBfVyza3NHyuLiyhvalsFbasA4HNcnK+4/3eZLMo+e6yvHUsY4WhmZWW5u5mHcDZm1jtI2fUtzod6Sx7PrCOhm+zUO1I98OD0EmICxHHMjMkGmSFuGhG8vXRKaRavzvV/fkscZK1HjW+XJOV0c9wYDZT+/LTjDS/yx5KkM5lIksiw57DsCro9gvmE1WdvB3bw3m+5ie2x1ZRSuh6ik1rbMTiXgJXt0MnbV8VsyPe1zhkbug1cwdhT9zcOTieb0oyiIl3wSAEkNNBNBaWB+78V1hM/ngMoZgKaaasJUCGtPH/rN+YnxGicDOg=="

	ciphertext2 := ciphertext[1:]

	encryptor, err := NewEncryptor(appID, token, encodingAesKey)
	if err != nil {
		log.Fatal("NewEncryptor init fail: ", err)
	}

	// The signature should match the expected
	// 签名应该和预期的一致
	assert.Equal(t, encryptor.Signature(token, timestamp, nonce, ciphertext), msgSignature)

	// Since the ciphertext data has been changed, the signatures should not be equal.
	// 因为密文数据被改过了，所以签名不应该相等。
	assert.NotEqual(t, encryptor.Signature(token, timestamp, nonce, ciphertext2), msgSignature)
}

func TestEncryptor_VerifyUrl(t *testing.T) {
	msgSignature := "1495c4dfd4958d4e5faf618978ae66943a042f87"
	timestamp := "1623292419"
	nonce := "1623324060"

	// original ciphertext content
	// 原始密文内容
	ciphertext := "o1XtmVltGmUAqoWee54yd4Q5ZBgrw4/9lFo5qdZoVPd1DybzarjuYCfFlR2AFbAcWHwFgmbrVBD+f9910QIF6g=="

	encryptor, err := NewEncryptor(appID, token, encodingAesKey)
	if err != nil {
		log.Fatal("NewEncryptor init fail: ", err)
	}

	msg, encryptError := encryptor.VerifyUrl(ciphertext, msgSignature, nonce, timestamp)
	if encryptError != nil {
		log.Fatal(encryptError)
	}

	assert.Equal(t, string(msg), "6058207018001857403")
}
