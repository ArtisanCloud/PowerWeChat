package handlers

import (
	"github.com/ArtisanCloud/go-libs/fmt"
	"github.com/ArtisanCloud/power-wechat/src/kernel/contract"
	"github.com/ArtisanCloud/power-wechat/src/work/server/handlers/models"
)

type ServerCallbackHandler struct {
	Callback func(header contract.EventInterface, content interface{}) interface{}
}

func NewServerCallbackHandler() *ServerCallbackHandler {
	return &ServerCallbackHandler{}
}

func (handler *ServerCallbackHandler) Handle(header contract.EventInterface, content interface{}) interface{} {

	if handler.Callback != nil {
		event := header.GetEvent()
		switch event {
		case "change_contact":
			eventUserCreate := content.(*models.EventUserCreate)
			for item := range eventUserCreate.ExtAttr.Item {
				fmt.Dump(item)
			}
		}
		return handler.Callback(header, content)
	}

	return nil
}
