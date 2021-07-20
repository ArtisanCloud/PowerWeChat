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
	"io/ioutil"
	"strings"
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
