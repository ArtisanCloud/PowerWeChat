package support

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

const (
	ErrorBase64Decode  = -40010 // base64 解码失败
	ErrorEncryptAes    = -40006 // AES 加密失败
	ErrorDecryptAes    = -40007 // AES 解密失败
	ErrorInvalidAesKey = -40004 // AESKey 非法
	ErrorInvalidIv     = -10001 // AESKey 非法
)

type CryptError struct {
	ErrCode int
	ErrMsg  string
}

func NewCryptError(errCode int, errMsg string) *CryptError {
	return &CryptError{ErrCode: errCode, ErrMsg: errMsg}
}

type AES struct{}

func NewAES() *AES {
	return &AES{}
}

func (a AES) Encrypt(text []byte, key, iv []byte) ([]byte, *CryptError) {
	const blockSize = 32
	padMsg := a.PKCS7Padding(text, blockSize)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, NewCryptError(ErrorEncryptAes, err.Error())
	}

	ciphertext := make([]byte, len(padMsg))

	mode := cipher.NewCBCEncrypter(block, iv)

	mode.CryptBlocks(ciphertext, padMsg)
	base64Msg := make([]byte, base64.StdEncoding.EncodedLen(len(ciphertext)))
	base64.StdEncoding.Encode(base64Msg, ciphertext)

	return base64Msg, nil
}

// Decrypt AES解密
// cipherText 是密文的base64字符串
// key和iv是经过反base64之后的byte内容
// 例如：原始值是"tiihtNczf5v6AKRyjwEUhQ=="，那么应该用base64.StdEncoding.DecodeString("tiihtNczf5v6AKRyjwEUhQ==")解码之后传过来
func (a AES) Decrypt(cipherText string, key, iv []byte) ([]byte, *CryptError) {

	encryptMsg, err := base64.StdEncoding.DecodeString(cipherText)
	//log.Println("cipherText: ", cipherText, err)
	if nil != err {
		return nil, NewCryptError(ErrorBase64Decode, err.Error())
	}

	//log.Println("ready to gen block")
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, NewCryptError(ErrorDecryptAes, err.Error())
	}

	if len(encryptMsg) < aes.BlockSize {
		return nil, NewCryptError(ErrorDecryptAes, "encrypt_msg size is not valid")
	}

	//iv := aesKey[:aes.BlockSize]

	if len(encryptMsg)%aes.BlockSize != 0 {
		return nil, NewCryptError(ErrorDecryptAes, "encrypt_msg not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	mode.CryptBlocks(encryptMsg, encryptMsg)

	//log.Println("encryptMsg: ", string(encryptMsg))
	originalMsg, err2 := a.PKCS7UnPadding(encryptMsg)
	if err2 != nil {
		return nil, err2
	}

	return originalMsg, nil
}

// PKCS7Padding PKCS#7 padding.
func (a *AES) PKCS7Padding(text []byte, blockSize int) []byte {
	padding := blockSize - (len(text) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	var buffer bytes.Buffer
	buffer.Write(text)
	buffer.Write(padText)
	return buffer.Bytes()
}

// PKCS7UnPadding KCS#7 unPadding.
func (a *AES) PKCS7UnPadding(text []byte) ([]byte, *CryptError) {
	plaintextLen := len(text)
	if text == nil || plaintextLen == 0 {
		return nil, NewCryptError(ErrorDecryptAes, "pKCS7UnPadding error nil or zero")
	}

	// TODO： 不清楚微信官方为什么需要这里校验一下
	//if plaintextLen%encryptor.blockSize != 0 {
	//	return nil, NewCryptError(ErrorDecryptAes, "pKCS7UnPadding text not a multiple of the block size")
	//}
	paddingLen := int(text[plaintextLen-1])
	return text[:plaintextLen-paddingLen], nil
}

// DecryptAES256GCM 使用 AEAD_AES_256_GCM 算法进行解密
//
// 你可以使用此算法完成微信支付平台证书和回调报文解密，详见：
// https://wechatpay-api.gitbook.io/wechatpay-api-v3/qian-ming-zhi-nan-1/zheng-shu-he-hui-tiao-bao-wen-jie-mi
func DecryptAES256GCM(aesKey, associatedData, nonce, ciphertext string) (plaintext string, err error) {
	decodedCiphertext, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	c, err := aes.NewCipher([]byte(aesKey))
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", err
	}
	dataBytes, err := gcm.Open(nil, []byte(nonce), decodedCiphertext, []byte(associatedData))
	if err != nil {
		return "", err
	}
	return string(dataBytes), nil
}
