package miniProgram

import (
	"encoding/base64"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/support"
)

type Encryptor struct {
	aes *support.AES
}

func (encryptor Encryptor) DecryptData(encrypted string, sessionKey string, iv string) ([]byte, *support.CryptError) {
	encryptData := encrypted
	_key, _ := base64.StdEncoding.DecodeString(sessionKey)
	_iv, _ := base64.StdEncoding.DecodeString(iv)

	aes := support.AES{}
	data, err := aes.Decrypt(encryptData, _key, _iv)
	if err != nil {
		return nil, err
	}
	return data, err
}
