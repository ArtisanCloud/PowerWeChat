package contract

import "net/http"

type EventInterface interface {
	GetToUserName() string
	GetFromUserName() string
	GetCreateTime() string
	GetMsgType() string
	GetEvent() string
	GetChangeType() string
	ReadMessage(msg interface{}) error
}

type EventHandlerInterface interface {
	Handle(request *http.Request, header EventInterface, content interface{}) interface{}
}
