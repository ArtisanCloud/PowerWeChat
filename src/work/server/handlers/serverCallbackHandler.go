package handlers

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/contract"
)

type ServerCallbackHandler struct {
	Callback func(header contract.EventInterface, content interface{}) interface{}
}

func NewServerCallbackHandler() *ServerCallbackHandler {
	return &ServerCallbackHandler{}
}

func (handler *ServerCallbackHandler) Handle(header contract.EventInterface, content interface{}) interface{} {

	if handler.Callback != nil {
		return handler.Callback(header, content)
	}

	return nil
}
