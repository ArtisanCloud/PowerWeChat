package server

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/kernel/messages"
	"github.com/ArtisanCloud/power-wechat/src/work/server/handlers"
)

func RegisterProvider(app kernel.ApplicationInterface) (*kernel.Encryptor, *Guard) {
	config := app.GetConfig()

	encryptor, _ := kernel.NewEncryptor(
		(*config).Get("corp_id", "").(string),
		(*config).Get("token", "").(string),
		(*config).Get("aes_key", "").(string),
	)

	guard := NewGuard(&app)
	echoStrHandler := handlers.NewEchoStrHandler(&app)
	guard.Push(echoStrHandler, messages.VOID)

	return encryptor, guard
}
