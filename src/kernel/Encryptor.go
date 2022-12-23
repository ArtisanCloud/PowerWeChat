package kernel

import (
	"bytes"
	"crypto/aes"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"encoding/xml"
	"fmt"
	fmt2 "github.com/ArtisanCloud/PowerLibs/v3/fmt"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/support"
	"math/rand"
	"sort"
	"strings"
	"time"
)

// Wechat Docs: https://open.work.weixin.qq.com/api/doc/90000/90138/90307
const (
	ErrorInvalidSignature = -40001 // 签名验证错误
	ErrorParseXml         = -40002 // xml/json解析失败
	ErrorCalcSignature    = -40003 // sha加密生成签名失败
	ErrorInvalidAesKey    = -40004 // AESKey 非法
	ErrorInvalidAppID     = -40005 // ReceiveID 校验错误
	ErrorEncryptAes       = -40006 // AES 加密失败
	ErrorDecryptAes       = -40007 // AES 解密失败
	ErrorInvalidXml       = -40008 // 解密后得到的buffer非法
	ErrorBase64Encode     = -40009 // base64 编码失败
	ErrorBase64Decode     = -40010 // base64 解码失败
	ErrorXmlBuild         = -40011 // 生成xml失败
	IllegalBuffer         = -41003 // Illegal buffer
	letterBytes           = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

type CDATA struct {
	Value string `xml:",cdata"`
}

type WeComReplyMsg struct {
	XMLName   xml.Name `xml:"xml"`
	Encrypt   CDATA    `xml:"Encrypt"`
	Signature CDATA    `xml:"MsgSignature"`
	Timestamp string   `xml:"TimeStamp"`
	Nonce     CDATA    `xml:"Nonce"`
}

type WeComRecvMsg struct {
	ToUserName string `xml:"ToUserName"`
	Encrypt    string `xml:"Encrypt"`
	AgentID    string `xml:"AgentID"`
}

type Encryptor struct {
	appID     string // App token
	token     string
	aesKey    []byte
	blockSize int64
	aes       *support.AES
}

func NewEncryptor(appID, token, aesKey string) (*Encryptor, error) {

	var aesKeyByte []byte
	var err error
	if aesKey == "" {
		fmt2.Dump("AES Key is empty, this may occur errors when decode callbacks message")
	}
	aesKeyByte, err = base64.StdEncoding.DecodeString(aesKey)
	if err != nil {
		aesKeyByte, err = base64.RawStdEncoding.DecodeString(aesKey)
		if err != nil {
			return nil, err
		}
	}

	return &Encryptor{
		appID:     appID,
		token:     token,
		aesKey:    aesKeyByte,
		blockSize: 32,
		aes:       &support.AES{},
	}, nil
}

// GetToken Get the app token
func (encryptor *Encryptor) GetToken() string {
	return encryptor.token
}

func (encryptor *Encryptor) randString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

// Encrypt encrypt xml msg and return xml
func (encryptor *Encryptor) Encrypt(msg, nonce, timestamp string) ([]byte, *support.CryptError) {

	randStr := encryptor.randString(16)
	var buffer bytes.Buffer
	buffer.WriteString(randStr)

	msgLenBuf := make([]byte, 4)
	binary.BigEndian.PutUint32(msgLenBuf, uint32(len(msg)))
	buffer.Write(msgLenBuf)
	buffer.WriteString(msg)

	// TODO 暂时不知道用途
	//buffer.WriteString(wxBizMsgCrypt.receiverID)
	buffer.WriteString(encryptor.appID)

	tmpCiphertext, err := encryptor.aes.Encrypt(buffer.Bytes(), encryptor.aesKey, encryptor.aesKey[:aes.BlockSize])
	if err != nil {
		return nil, (*support.CryptError)(err)
	}
	ciphertext := string(tmpCiphertext)

	if timestamp == "" {
		timestamp = fmt.Sprintf("%d", time.Now().Unix())
	}
	if nonce == "" {
		nonce = object.QuickRandom(10)
	}
	signature := encryptor.Signature(encryptor.token, timestamp, nonce, ciphertext)

	msg4Send := &WeComReplyMsg{
		Encrypt:   CDATA{Value: ciphertext},
		Signature: CDATA{Value: signature},
		Timestamp: timestamp,
		Nonce:     CDATA{Value: nonce},
	}
	//msg4Send := NewWXBizMsg4Send(ciphertext, signature, timestamp, nonce)
	//msg.Marshal()
	xmlByte, err2 := xml.Marshal(msg4Send)
	if err2 != nil {
		return nil, support.NewCryptError(ErrorXmlBuild, err2.Error())
	}
	return xmlByte, nil
}

// Decrypt decrypt xml msg and return xml
func (encryptor *Encryptor) Decrypt(content []byte, msgSignature, nonce, timestamp string) ([]byte, *support.CryptError) {
	var msg4Recv WeComRecvMsg
	err := xml.Unmarshal(content, &msg4Recv)

	if err != nil {
		return nil, support.NewCryptError(ErrorDecryptAes, err.Error())
	}

	signature := encryptor.Signature(encryptor.token, timestamp, nonce, msg4Recv.Encrypt)

	if strings.Compare(signature, msgSignature) != 0 {
		return nil, support.NewCryptError(ErrorCalcSignature, "signature not equal")
	}

	iv := encryptor.aesKey[:aes.BlockSize]
	plaintext, cryptErr := encryptor.aes.Decrypt(msg4Recv.Encrypt, encryptor.aesKey, iv)
	//plaintext, cryptErr := wxBizMsgCrypt.cbcDecrypter(msg4Recv.Encrypt)
	if nil != cryptErr {
		return nil, cryptErr
	}

	textLen := uint32(len(plaintext))
	if textLen < 20 {
		return nil, support.NewCryptError(IllegalBuffer, "plain is to small 1")
	}
	//random := plaintext[:16]
	msgLen := binary.BigEndian.Uint32(plaintext[16:20])
	if textLen < (20 + msgLen) {
		return nil, support.NewCryptError(IllegalBuffer, "plain is to small 2")
	}

	msg := plaintext[20 : 20+msgLen]
	//receiverID := plaintext[20+msgLen:]

	return msg, nil
}

// Decrypt decrypt xml msg and return xml
func (encryptor *Encryptor) DecryptContent(content string) ([]byte, *support.CryptError) {

	iv := encryptor.aesKey[:aes.BlockSize]
	plaintext, cryptErr := encryptor.aes.Decrypt(content, encryptor.aesKey, iv)
	//plaintext, cryptErr := wxBizMsgCrypt.cbcDecrypter(msg4Recv.Encrypt)
	if nil != cryptErr {
		return nil, cryptErr
	}

	textLen := uint32(len(plaintext))
	if textLen < 20 {
		return nil, support.NewCryptError(IllegalBuffer, "plain is to small 1")
	}
	//random := plaintext[:16]
	msgLen := binary.BigEndian.Uint32(plaintext[16:20])
	if textLen < (20 + msgLen) {
		return nil, support.NewCryptError(IllegalBuffer, "plain is to small 2")
	}

	msg := plaintext[20 : 20+msgLen]
	//receiverID := plaintext[20+msgLen:]

	return msg, nil
}

// VerifyUrl Parsing the official WeChat callbacks validation.
// Wechat Docs: https://work.weixin.qq.com/api/doc/90000/90135/90235
// When adding URLs to the WeChat admin backend, WeChat will trigger a GET request to verify whether the server can process the encrypted information properly, and the uniformMessage needs to be decrypted and returned.
// 在微信管理后台添加URL的时候，微信会触发一条GET请求用于验证服务器能否正常处理加密信息，需要将消息解密出来返回。
// eg: "/app-callbacks?msg_signature=1495c4dfd4958d4e5faf618978ae66943a042f87&timestamp=1623292419&nonce=1623324060&echostr=o1XtmVltGmUAqoWee54yd4Q5ZBgrw4%2F9lFo5qdZoVPd1DybzarjuYCfFlR2AFbAcWHwFgmbrVBD%2Bf9910QIF6g%3D%3D"
func (encryptor *Encryptor) VerifyUrl(content string, msgSignature, nonce, timestamp string) ([]byte, *support.CryptError) {
	msg4Recv := &WeComRecvMsg{
		Encrypt: content,
	}
	msg4RecvByte, err := xml.Marshal(msg4Recv)
	if err != nil {
		return nil, support.NewCryptError(ErrorXmlBuild, err.Error())
	}
	return encryptor.Decrypt(msg4RecvByte, msgSignature, nonce, timestamp)
}

// Signature get sha1

func (encryptor *Encryptor) Signature(token, timestamp, nonce, data string) string {
	sortArr := []string{token, timestamp, nonce, data}
	sort.Strings(sortArr)
	var buffer bytes.Buffer
	for _, value := range sortArr {
		buffer.WriteString(value)
	}

	sha := sha1.New()
	sha.Write(buffer.Bytes())
	signature := fmt.Sprintf("%x", sha.Sum(nil))
	return signature
}
