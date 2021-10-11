package contract



type EventInterface interface {
	GetToUserName() string
	GetFromUserName() string
	GetCreateTime() string
	GetMsgType() string
	GetEvent() string
	GetChangeType() string
}


type EventHandlerInterface interface {
	Handle(header EventInterface, content interface{}) interface{}
}