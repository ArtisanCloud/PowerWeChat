package contract

type EventHandlerInterface interface {

	handle(payload interface{})
}