package kernel

import (
	"bytes"
	"crypto/aes"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/support"
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
)

type CDATA struct {
	Value string `xml:",cdata"`
}

type WeComReplyMsg struct {
	XMLName   xml.Name `xml:"xml" json:"-"`
	Encrypt   CDATA    `xml:"Encrypt" json:"Encrypt"`
	Signature CDATA    `xml:"MsgSignature" json:"Signature"`
	Timestamp string   `xml:"TimeStamp" json:"Timestamp"`
	Nonce     CDATA    `xml:"Nonce" json:"Nonce"`
}

type WeComRecvMsg struct {
	ToUserName string `xml:"ToUserName" json:"ToUserName"`
	Encrypt    string `xml:"Encrypt" json:"Encrypt"`
	AgentID    string `xml:"AgentID" json:"AgentID"`
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
	//if aesKey == "" {
	//	fmt2.Dump("AES Key is empty, this may occur errors when decode callbacks message")
	//}
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

// Encrypt preserves the XML envelope used by application callbacks.
func (encryptor *Encryptor) Encrypt(msg, nonce, timestamp string) ([]byte, *support.CryptError) {
	envelope, err := encryptor.EncryptEnvelope(msg, nonce, timestamp)
	if err != nil {
		return nil, err
	}
	data, marshalErr := xml.Marshal(envelope)
	if marshalErr != nil {
		return nil, support.NewCryptError(ErrorXmlBuild, marshalErr.Error())
	}
	return data, nil
}

// EncryptEnvelope returns the signed fields for XML or JSON callback replies.
func (encryptor *Encryptor) EncryptEnvelope(msg, nonce, timestamp string) (*WeComReplyMsg, *support.CryptError) {
	if len(encryptor.aesKey) < aes.BlockSize || uint64(len(msg)) > uint64(^uint32(0)) {
		return nil, support.NewCryptError(ErrorEncryptAes, "invalid key or message length")
	}
	// Keep the 16-byte random prefix printable for existing callback consumers.
	var entropy [12]byte
	if _, err := rand.Read(entropy[:]); err != nil {
		return nil, support.NewCryptError(ErrorEncryptAes, err.Error())
	}
	var buffer bytes.Buffer
	buffer.WriteString(base64.RawStdEncoding.EncodeToString(entropy[:]))
	msgLenBuf := make([]byte, 4)
	binary.BigEndian.PutUint32(msgLenBuf, uint32(len(msg)))
	buffer.Write(msgLenBuf)
	buffer.WriteString(msg)

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

	return &WeComReplyMsg{
		Encrypt:   CDATA{Value: ciphertext},
		Signature: CDATA{Value: signature},
		Timestamp: timestamp,
		Nonce:     CDATA{Value: nonce},
	}, nil
}

// Decrypt decrypt xml msg and return xml
func (encryptor *Encryptor) Decrypt(content []byte, msgSignature, nonce, timestamp string) ([]byte, *support.CryptError) {
	var msg4Recv WeComRecvMsg
	var err error
	// 根据第一个字符判断是json还是xml
	if content[0] == '<' {
		err = xml.Unmarshal(content, &msg4Recv)
	} else if content[0] == '{' {
		err = json.Unmarshal(content, &msg4Recv)
	}

	if err != nil {
		return nil, support.NewCryptError(ErrorDecryptAes, err.Error())
	}

	signature := encryptor.Signature(encryptor.token, timestamp, nonce, msg4Recv.Encrypt)

	if strings.Compare(signature, msgSignature) != 0 {
		return nil, support.NewCryptError(ErrorCalcSignature, "signature not equal")
	}

	return encryptor.DecryptContent(msg4Recv.Encrypt)
}

// OpenCallback authenticates the ciphertext and checks the receiver in the decrypted frame.
func (encryptor *Encryptor) OpenCallback(ciphertext, signature, nonce, timestamp string) ([]byte, *support.CryptError) {
	want := encryptor.Signature(encryptor.token, timestamp, nonce, ciphertext)
	if !hmac.Equal([]byte(want), []byte(signature)) {
		return nil, support.NewCryptError(ErrorInvalidSignature, "invalid callback signature")
	}
	return encryptor.decryptContent(ciphertext, true)
}

func (encryptor *Encryptor) DecryptContent(content string) ([]byte, *support.CryptError) {
	return encryptor.decryptContent(content, false)
}

func (encryptor *Encryptor) decryptContent(content string, checkReceiver bool) ([]byte, *support.CryptError) {
	if len(encryptor.aesKey) < aes.BlockSize {
		return nil, support.NewCryptError(ErrorInvalidAesKey, "invalid AES key length")
	}
	plain, err := encryptor.aes.Decrypt(content, encryptor.aesKey, encryptor.aesKey[:aes.BlockSize])
	if err != nil {
		return nil, err
	}
	if len(plain) < 20 {
		return nil, support.NewCryptError(IllegalBuffer, "invalid callback frame")
	}
	size := binary.BigEndian.Uint32(plain[16:20])
	if uint64(size) > uint64(len(plain)-20) {
		return nil, support.NewCryptError(IllegalBuffer, "invalid callback message length")
	}
	end := 20 + int(size)
	if checkReceiver && string(plain[end:]) != encryptor.appID {
		return nil, support.NewCryptError(ErrorInvalidAppID, "invalid callback receiver")
	}
	return plain[20:end], nil
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

func CalcPaySig(uri, postBody, appkey string) string {
	needSignMsg := uri + "&" + postBody
	h := hmac.New(sha256.New, []byte(appkey))
	h.Write([]byte(needSignMsg))
	paySig := hex.EncodeToString(h.Sum(nil))
	return paySig
}

func CalcSignature(postBody, sessionKey string) string {
	needSignMsg := postBody
	h := hmac.New(sha256.New, []byte(sessionKey))
	h.Write([]byte(needSignMsg))
	signature := hex.EncodeToString(h.Sum(nil))
	return signature
}
