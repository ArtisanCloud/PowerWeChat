package contract

import "github.com/ArtisanCloud/power-wechat/src/work/server/handlers"

type EventHandlerInterface interface {
	//Handle(payload interface{}) interface{}
	Handle(header handlers.CallbackInterface, content interface{}) interface{}
}
