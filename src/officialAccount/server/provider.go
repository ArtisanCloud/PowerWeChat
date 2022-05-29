package server

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/messages"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/work/server/handlers"
	"log"
)

func RegisterProvider(app kernel.ApplicationInterface) (*kernel.Encryptor, *Guard) {
	config := app.GetConfig()

	encryptor, err := kernel.NewEncryptor(
		(*config).GetString("corp_id", ""),
		(*config).GetString("token", ""),
		(*config).GetString("aes_key", ""),
	)

	// 如果密钥之类初始化失败，提示一个警告，这个可能会导致后面的消息解密出错。
	if err != nil {
		log.Printf("encryptor init warn: %v", err)
		//panic(err)
	}

	guard := NewGuard(&app)
	echoStrHandler := handlers.NewEchoStrHandler(&app)
	guard.Push(echoStrHandler, messages.VOID)

	return encryptor, guard
}
