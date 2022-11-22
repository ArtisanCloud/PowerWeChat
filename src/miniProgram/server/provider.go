package server

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/messages"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/officialAccount/server"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/work/server/handlers"
)

func RegisterProvider(app kernel.ApplicationInterface) (*server.Guard, error) {

	guard := server.NewGuard(&app)
	echoStrHandler := handlers.NewEchoStrHandler(&app)
	guard.Push(echoStrHandler, messages.VOID)

	return guard, nil
}
