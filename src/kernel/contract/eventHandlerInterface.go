package contract

type EventHandlerInterface interface {
	Handle(payload ...interface{}) interface{}
}
