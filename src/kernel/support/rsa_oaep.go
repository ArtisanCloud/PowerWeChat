package support

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"hash"
	"io/ioutil"
)

type RSAOaep struct {
	PublicKeyPath  string // RSA公钥路径，会自动读取出*rsa.PrivateKey
	PublicKey      *rsa.PublicKey
	PrivateKeyPath string
	PrivateKey     *rsa.PrivateKey
}

func (r *RSAOaep) EncryptOAEP(text []byte) ([]byte, error) {
	if r.PublicKey == nil && r.PublicKeyPath == "" {
		return nil, fmt.Errorf("you must set publickey to use EncryptOAEP")
	}

	// Parsing the certificate inside the PrivateKeyPath
	// 读取PublicKeyPath里面的证书
	if r.PublicKey == nil {
		publicKey, err := r.getPublicKey()
		if err != nil {
			return nil, err
		}
		r.PublicKey = publicKey
	}

	rng := rand.Reader
	cipherData, err := rsa.EncryptOAEP(sha1.New(), rng, r.PublicKey, text, nil)
	if err != nil {
		return cipherData, err
	}

	return cipherData, nil
}

func (r *RSAOaep) DecryptOAEP(hash hash.Hash,cipherMsg []byte) ([]byte, error) {
	if r.PrivateKey == nil && r.PrivateKeyPath == "" {
		return nil, fmt.Errorf("you must set privatekey to use EncryptOAEP")
	}

	// Parsing the certificate inside the PrivateKeyPath
	// 读取PrivateKeyPath里面的证书
	if r.PrivateKey == nil {
		privateKey, err := r.getPrivateKey()
		if err != nil {
			return nil, err
		}
		r.PrivateKey = privateKey
	}

	decryptedData, decryptErr := rsa.DecryptOAEP(hash, rand.Reader, r.PrivateKey, cipherMsg, nil)
	if decryptErr != nil {
		fmt.Println("Decrypt data error")
		return nil, decryptErr
	}

	return decryptedData, nil
}

func (r *RSAOaep) getPublicKey() (*rsa.PublicKey, error) {
	keyStr, err := ioutil.ReadFile(r.PublicKeyPath)

	if err != nil {
		return nil, err
	}

	publicPem, _ := pem.Decode(keyStr)
	publicPemBytes := publicPem.Bytes

	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKIXPublicKey(publicPemBytes); err != nil {
		// note this returns type `interface{}`
		return nil, err
	}

	publicKey, ok := parsedKey.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("%s is not rsa public key", r.PublicKeyPath)
	}
	return publicKey, err
}

func (r *RSAOaep) getPrivateKey() (*rsa.PrivateKey, error) {
	keyStr, err := ioutil.ReadFile(r.PrivateKeyPath)

	if err != nil {
		return nil, err
	}

	privPem, _ := pem.Decode(keyStr)
	privPemBytes := privPem.Bytes

	var parsedKey interface{}
	if parsedKey, err = x509.ParsePKCS1PrivateKey(privPemBytes); err != nil { // note this returns type `interface{}`
		return nil, err
	}

	privateKey, ok := parsedKey.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("%s is not rsa private key", r.PrivateKeyPath)
	}
	return privateKey, err
}
