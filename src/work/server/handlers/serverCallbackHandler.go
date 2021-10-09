package handlers

import (
	"github.com/ArtisanCloud/go-libs/fmt"
	"github.com/ArtisanCloud/power-wechat/src/work/server"
)

type ServerCallbackHandler struct {
	Callback func(payload interface{}) interface{}
}

func NewServerCallbackHandler() *ServerCallbackHandler {
	return &ServerCallbackHandler{}
}

func (handler *ServerCallbackHandler) Handle(payload server.CallbackInterface, content interface) interface{} {

	if handler.Callback != nil {
		event := payload.GetEvent()
		switch event {
		case "change_contact":
			eventUserCreate:=content.(*server.EventUserCreate)
			for item:= range eventUserCreate.ExtAttr.Item{
				fmt.Dump(item)
			}
			
		}
		return handler.Callback(payload)
	}

	return nil
}

