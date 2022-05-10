package support

import (
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"github.com/go-playground/assert/v2"
	"log"
	"testing"
)

var (
	testPlainMsg           = []byte("Hello PowerWechat")
	testEncryptedMSGBase64 = "YYQEzEq5lYJPbYF7r2-UfoOPp7LzeEjbEbRYK_H3EYkDZNe_LvBz3f1Q7n3FNQtoXQciFV59CWXA9vo9RvsXKojARs7wZHGjp9WvOGNkM90s6pHI1lSP5lvsLi3VoxNbLRwsLwhTELGD75SENgmSBmaSNyv0NkVJiBbmYbW7kKBTD3A9KQYS12e8P0YCWoKWzPk1jJoVlpuv-gsI8BHZcoxIvouZG3sx9G0l6OvJDrATJV4kPaTwnq7LAFIOBJhV0xfbdGC979h6piYZ8yXFd2Ruy1zgOU4tXERSWQxhDb5PmvCAgrsLbOOHcfSqB-UOM0E64HQGFgmoxS6-UwyFyg=="
	testPublicKeyData      = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAygGoUiTD+LjwZIgwFZyj
iibWNQ2LM9xZ2pjKQGP8iUBtAuAW629/Ofw8qxToMyixPrG4A7j8+KOPwYrWPGV6
Og//4zm3cG+1hQvnNUWtMjHHBY8OByUPQ6/T8XHER1DxFBfnWfFLZ1yFX6oNNuvt
LgOreI6ehehJd5IB/4mOjMvFEBgOEejado2n55VNdcFpdQ3RcvGV+f/rl/lsIM08
QvL3lc5gqawj53sW9YZi1DL/uN48R+ghvAYhtx2jpHDBvlH1NCF1rU6CynYsgV9Q
Iksv0ihwl4T+k5F9ir0uv0WIS6kKKS1SRpAprRKunos4PlE8l2+jC6LaJUPhDZlj
/wIDAQAB
-----END PUBLIC KEY-----
`
	testPrivateKeyData = `
-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEAygGoUiTD+LjwZIgwFZyjiibWNQ2LM9xZ2pjKQGP8iUBtAuAW
629/Ofw8qxToMyixPrG4A7j8+KOPwYrWPGV6Og//4zm3cG+1hQvnNUWtMjHHBY8O
ByUPQ6/T8XHER1DxFBfnWfFLZ1yFX6oNNuvtLgOreI6ehehJd5IB/4mOjMvFEBgO
Eejado2n55VNdcFpdQ3RcvGV+f/rl/lsIM08QvL3lc5gqawj53sW9YZi1DL/uN48
R+ghvAYhtx2jpHDBvlH1NCF1rU6CynYsgV9QIksv0ihwl4T+k5F9ir0uv0WIS6kK
KS1SRpAprRKunos4PlE8l2+jC6LaJUPhDZlj/wIDAQABAoIBAHIcX5YPeLie2AUi
PW9n7aYT7DtJ7FGebw+h8dZP5Q8vWqUeKzRR5p+90hOemtCTcxSEVfucWyKlWoat
Q/oYJOR5t0YHi40zPWnr4G7ibkUFg3Sra/QzRh0pTON+La9PlO+R1TmkqcC4rgrt
R8u3mGK+5fUTM49XOXEXBJPyg5kaXQpiA4BoIRdRnCSitNxWA8kxMkQYJYlwAYab
cKo4Ik/J6+YGG7m2FtrUAWpWVUMBzEYOmGJ7JhSJ1u0UC/Oh1HOS1xlGopkmexbd
EygY3hTNWzHmYaYcYQs0f+8aVcVL64Gm0dtqvAHNnBvudMThhQgdYPc39mNLbrwI
ks4uS8ECgYEA9XfvcGKsNrHA0nqoPUPMT0Nfvv/4XCaKOYk25brH4LbqJPm6CiU6
uNlKFQsxzHPmx7OEK7EYVVZCbSO9s4t/xCzDVNbOZ9kDL6bkTX9DArLE4d6IRF/1
WW/AlNPuwVgxl0kcJILFtLqA1WoC5UWMhbYe2YB/Q3rCozmn0AiwyqECgYEA0qxd
KClKAMIsrB0WJ9gZEsJOpFi4q4g6T1BwT40Xj6Ul6o6DHi6hFhPgZAstqmnY0ANz
ezQ2yxtIm7zSy7S+nwDUycjY9riJcomc/YQZNA2QVM16hEv84VLwH1MVV2wkTb41
DWjbcg/ZNofZHl9AQIw7es+R3mmtDN+8BZOZSp8CgYBHtwmaUQm1VQtbswAyHfuz
8KApgklCSvQ5SRBj38UDrw0LTnZ+/k+Ar+MH8ORUskvrblQgG7ZbQD9Z+YYzzX6/
hsBuqe9Vwb4/jsfGqHagdDA3OTegmlRpE9A06xInJKggZfi15gry+UYok7dS2pXq
fsHWk8capOP2oiKYEeHs4QKBgF2KcLaDVrtte/5Tz+GTHtbodZidWCm5jAJpeeSo
hfye3G4AJxHArH+sBacGG5md88mwrpbWwTl/fMbBmWsfbsAU02ZhCozJtSWpGo6q
F7K4DwzIS4zwXHEDrWCLOF+fwaLPQKkalM1ZYh3HRc0ph9LhMQu/nEn/6/laYhar
yZWLAoGASvCrpFKn0qllMKNUetBmYFpgtjmnNuW7l0xT2UftkW6AuFjU19gKgXhe
I+uZciHQ8kIUHfNLYBbhETsF3iqsklKfeoIr23zYHLE5GpoC151IpKf4guoPbCHX
a1oCDuZm//f5HMePb9juJN0WR//d5jWuizAycZf41XoEd8Bqydg=
-----END RSA PRIVATE KEY-----
`
)

// demo from github gist https://gist.github.com/jemygraw/4da51d58b349bfddd0c5
func TestRSA_EncryptOAEP(t *testing.T) {
	pubKeyBlock, _ := pem.Decode([]byte(testPublicKeyData))
	pubInterface, parseErr := x509.ParsePKIXPublicKey(pubKeyBlock.Bytes)
	if parseErr != nil {
		fmt.Println("Load public key error")
		panic(parseErr)
	}
	PublicKey := pubInterface.(*rsa.PublicKey)
	privateBlock, _ := pem.Decode([]byte(testPrivateKeyData))
	privateKey, parseErr := x509.ParsePKCS1PrivateKey(privateBlock.Bytes)
	if parseErr != nil {
		fmt.Println("Load private key error")
		panic(parseErr)
	}

	rsaOaep := &RSAOaep{
		PublicKey:  PublicKey,
		PrivateKey: privateKey,
	}

	encryptedData, err := rsaOaep.EncryptOAEP(testPlainMsg)
	if err != nil {
		panic(err)
	}
	encryptedDataBase64 := base64.URLEncoding.EncodeToString(encryptedData)
	log.Println("encryptedData: ", encryptedDataBase64)

	// 因为每次加密的随机数和hash都是不一样的，所以每次同一个明文加密出来的密文都会不一样
	// 验证方法就是从加密出来的密文再解密，看下明文是否一致
	plainMsg, _ := rsaOaep.DecryptOAEP(sha1.New(),encryptedData)

	assert.Equal(t, testPlainMsg, plainMsg)
}

func TestRSAOaep_DecryptOAEP(t *testing.T) {
	privateBlock, _ := pem.Decode([]byte(testPrivateKeyData))
	//msg := []byte("helloworldjemyme")
	privateKey, parseErr := x509.ParsePKCS1PrivateKey(privateBlock.Bytes)
	if parseErr != nil {
		fmt.Println("Load private key error")
		panic(parseErr)
	}

	rsaOaep := &RSAOaep{
		//PublicKeyPath: "",
		PrivateKey: privateKey,
	}

	testEncryptedMSG, _ := base64.URLEncoding.DecodeString(testEncryptedMSGBase64)
	decryptedData, err := rsaOaep.DecryptOAEP(sha1.New(),[]byte(testEncryptedMSG))
	if err != nil {
		panic(err)
	}

	log.Println("decryptedData: ", string(decryptedData))
	assert.Equal(t, decryptedData, testPlainMsg)
}
