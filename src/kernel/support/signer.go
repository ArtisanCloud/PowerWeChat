package support

import (
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"io/ioutil"
	"strings"
	"time"
)

// 请求报文签名相关常量
const (
	SignatureMessageFormat = "%s\n%s\n%s\n%s\n%s\n" // 数字签名原文格式

	// HeaderAuthorizationFormat 请求头中的 Authorization 拼接格式
	HeaderAuthorizationFormat = "%s mchid=\"%s\",nonce_str=\"%s\",timestamp=\"%d\",serial_no=\"%s\",signature=\"%s\""
)

// SHA256WithRSASigner Sha256WithRSA 数字签名生成器
type SHA256WithRSASigner struct {
	MchID               string          // 商户号
	CertificateSerialNo string          // 商户证书序列号
	PrivateKeyPath      string          // 商户私钥路径，会自动读取出*rsa.PrivateKey
	PrivateKey          *rsa.PrivateKey // 商户私钥
}

// SignatureResult 数字签名结果
type SignatureResult struct {
	MchID               string // 商户号
	CertificateSerialNo string // 签名对应的证书序列号
	Signature           string // 签名内容
}

type RequestSignChain struct {
	Method       string // 接口提交方法。http.MethodPost, http.MethodPost等
	CanonicalURL string // 微信支付接口路径。 例如： /v3/pay/transactions/jsapi
	SignBody     string // 提交的body字符串。 例如； {"amount":{"total":1},"appid":"ww16143ea0101327c7","attach":"自定义数据说明","description":"Image形象店-深圳腾大-QQ公仔","mchid":"1611854986","notify_url":"https://pay.wangchaoyi.com/wx/notify","out_trade_no":"5519778939773395659222199361","payer":{"openid":"oAuaP0TRUMwP169nQfg7XCEAw3HQ"}}
	timestamp    int64  // 单元测试传入的固定时间戳
	nonce        string // 单元测试传入的固定随机数
}

func (s *SHA256WithRSASigner) GenerateRequestSign(signChain *RequestSignChain) (authorization string, err error) {
	timestamp := time.Now().Unix()
	nonce := object.QuickRandom(32)

	// Under ci mode, go fixed value
	// 在ci模式下面，走固定值
	if signChain.timestamp != 0 && signChain.nonce != "" {
		timestamp = signChain.timestamp
		nonce = signChain.nonce
	}

	// Splice the string to be signed
	// 拼接出需要签名的字符串
	arrayParams := []string{signChain.Method, signChain.CanonicalURL, fmt.Sprintf("%d", timestamp), nonce, signChain.SignBody}
	message := strings.Join(arrayParams, "\n") + "\n"

	// sign the uniformMessage
	signatureResult, err := s.Sign(context.TODO(), message)
	if err != nil {
		return "", err
	}

	// get the header authorization
	authorization = fmt.Sprintf(
		HeaderAuthorizationFormat,
		s.GetAuthorizationType(),
		signatureResult.MchID,
		nonce,
		timestamp,
		signatureResult.CertificateSerialNo,
		signatureResult.Signature,
	)

	return authorization, err
}

func (s *SHA256WithRSASigner) GenerateSign(message string) (sign string, err error) {

	// sign the uniformMessage
	signatureResult, err := s.Sign(context.TODO(), message)
	if err != nil {
		return "", err
	}

	return signatureResult.Signature, err
}

// Sign 对信息使用 SHA256WithRSA 算法进行签名
func (s *SHA256WithRSASigner) Sign(_ context.Context, message string) (*SignatureResult, error) {
	if s.PrivateKey == nil && s.PrivateKeyPath == "" {
		return nil, fmt.Errorf("you must set privatekey to use SHA256WithRSASigner")
	}

	// Parsing the certificate inside the PrivateKeyPath
	// 读取PrivateKeyPath里面的证书
	if s.PrivateKey == nil {
		privateKey, err := s.getPrivateKey()
		if err != nil {
			return nil, err
		}
		s.PrivateKey = privateKey
	}

	if strings.TrimSpace(s.CertificateSerialNo) == "" {
		return nil, fmt.Errorf("you must set mch certificate serial no to use SHA256WithRSASigner")
	}

	signature, err := SignSHA256WithRSA(message, s.PrivateKey)
	if err != nil {
		return nil, err
	}
	return &SignatureResult{MchID: s.MchID, CertificateSerialNo: s.CertificateSerialNo, Signature: signature}, nil
}

func (s *SHA256WithRSASigner) getPrivateKey() (*rsa.PrivateKey, error) {
	keyStr, err := ioutil.ReadFile(s.PrivateKeyPath)

	if err != nil {
		return nil, err
	}

	privPem, _ := pem.Decode(keyStr)
	privPemBytes := privPem.Bytes

	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKCS8PrivateKey(privPemBytes); err != nil { // note this returns type `interface{}`
		return nil, err
	}

	privateKey, ok := parsedKey.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("%s is not rsa private key", s.PrivateKeyPath)
	}
	return privateKey, err
}

// Algorithm 返回使用的签名算法：SHA256-RSA2048
func (s *SHA256WithRSASigner) Algorithm() string {
	return "SHA256-RSA2048"
}

func (s *SHA256WithRSASigner) GetAuthorizationType() string {
	return "WECHATPAY2-" + s.Algorithm()
}

// SignSHA256WithRSA 通过私钥对字符串以 SHA256WithRSA 算法生成签名信息
func SignSHA256WithRSA(source string, privateKey *rsa.PrivateKey) (signature string, err error) {
	if privateKey == nil {
		return "", fmt.Errorf("private key should not be nil")
	}
	h := crypto.Hash.New(crypto.SHA256)
	_, err = h.Write([]byte(source))
	if err != nil {
		return "", nil
	}
	hashed := h.Sum(nil)
	signatureByte, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signatureByte), nil
}
