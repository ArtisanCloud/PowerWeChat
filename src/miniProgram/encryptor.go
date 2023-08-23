package miniProgram

import (
	"encoding/base64"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/support"
)

type Encryptor struct {
	BaseEncryptor *kernel.Encryptor
	aes           *support.AES
}

func NewEncryptor(appID, token, aesKey string) (*Encryptor, error) {
	baseEncryptor, err := kernel.NewEncryptor(appID, token, aesKey)
	if err != nil {
		return nil, err
	}
	return &Encryptor{
		BaseEncryptor: baseEncryptor,
		aes:           &support.AES{},
	}, nil

}

func (encryptor Encryptor) DecryptData(encrypted string, sessionKey string, iv string) ([]byte, *support.CryptError) {
	encryptData := encrypted
	_key, errDecode := base64.StdEncoding.DecodeString(sessionKey)
	if errDecode != nil {
		return nil, support.NewCryptError(-1, errDecode.Error())
	}

	_iv, errDecode := base64.StdEncoding.DecodeString(iv)
	if errDecode != nil {
		return nil, support.NewCryptError(-1, errDecode.Error())
	}

	aes := support.AES{}
	data, err := aes.Decrypt(encryptData, _key, _iv)
	if err != nil {
		return nil, err
	}
	return data, err
}
