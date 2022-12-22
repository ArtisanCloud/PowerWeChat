package server

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/messages"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/server/handlers"
)

func RegisterProvider(app kernel.ApplicationInterface) (*kernel.Encryptor, *Guard, error) {
	config := app.GetConfig()

	encryptor, err := kernel.NewEncryptor(
		(*config).GetString("app_id", ""),
		(*config).GetString("token", ""),
		(*config).GetString("aes_key", ""),
	)

	// 如果密钥之类初始化失败，提示一个警告，这个可能会导致后面的消息解密出错。
	if err != nil {
		return nil, nil, err
	}

	guard := NewGuard(&app)
	echoStrHandler := handlers.NewEchoStrHandler(&app)
	guard.Push(echoStrHandler, messages.VOID)

	return encryptor, guard, nil
}
