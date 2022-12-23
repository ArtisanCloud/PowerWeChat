package server

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
)

type Guard struct {
	*kernel.ServerGuard
}

func NewGuard(app *kernel.ApplicationInterface) *Guard {
	//config := (*app).GetContainer().GetConfig()

	guard := &Guard{
		kernel.NewServerGuard(app),
	}

	guard.OverrideResolve()
	guard.OverrideNotify()

	return guard

}

//func (guard *Guard) DecryptEvent(content string) (bufDecrypted []byte, err error) {
//
//	encryptor := (*guard.App).GetComponent("Encryptor").(*miniProgram.Encryptor)
//
//	bufDecrypted, cryptErr := encryptor.DecryptData(content)
//	if cryptErr != nil {
//		return nil, errors.New(cryptErr.ErrMsg)
//	}
//
//	return bufDecrypted, err
//
//}
