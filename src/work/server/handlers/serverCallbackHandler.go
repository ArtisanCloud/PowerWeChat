package handlers

import (
	"github.com/ArtisanCloud/go-libs/fmt"
)

type ServerCallbackHandler struct {
	Callback func(header CallbackInterface, content interface{}) interface{}
}

func NewServerCallbackHandler() *ServerCallbackHandler {
	return &ServerCallbackHandler{}
}

func (handler *ServerCallbackHandler) Handle(header CallbackInterface, content interface{}) interface{} {

	if handler.Callback != nil {
		event := header.GetEvent()
		switch event {
		case "change_contact":
			eventUserCreate := content.(*EventUserCreate)
			for item := range eventUserCreate.ExtAttr.Item {
				fmt.Dump(item)
			}
		}
		return handler.Callback(header, content)
	}

	return nil
}
