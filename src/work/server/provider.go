package server

import (
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/messages"
	"github.com/ArtisanCloud/PowerWeChat/src/work/server/handlers"
)

func RegisterProvider(app kernel.ApplicationInterface) (*kernel.Encryptor, *Guard) {
	config := app.GetConfig()

	encryptor, _ := kernel.NewEncryptor(
		(*config).GetString("corp_id", ""),
		(*config).GetString("token", ""),
		(*config).GetString("aes_key", ""),
	)

	guard := NewGuard(&app)
	echoStrHandler := handlers.NewEchoStrHandler(&app)
	guard.Push(echoStrHandler, messages.VOID)

	return encryptor, guard
}
