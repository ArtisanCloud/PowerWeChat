package server

import (
	"github.com/ArtisanCloud/go-wechat/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) (*kernel.Encryptor, *Guard) {
	config := app.GetConfig()

	client, _ := kernel.NewEncryptor(
		(*config).Get("corp_id", "").(string),
		(*config).Get("token", "").(string),
		(*config).Get("aes_key", "").(string),
	)

	server := NewGuard(&app)
	//server.Push(NewEchoStrHandler(&app))

	return client, server
}
