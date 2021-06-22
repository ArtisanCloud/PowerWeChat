package support

import "github.com/ArtisanCloud/go-wechat/src/kernel/contract"

type Observable struct {
	handlers []*contract.EventHandlerInterface
}

func NewObservable() *Observable {
	return &Observable{}
}

func (observable *Observable) Push(closure contract.EventHandlerInterface) interface{} {
	observable.handlers = append(observable.handlers, &closure)

	return observable
}

func (observable *Observable) SetHandle(handlers []*contract.EventHandlerInterface) *Observable {
	observable.handlers = handlers

	return observable
}

func (observable *Observable) Dispatch(event string, payload interface{}) interface{} {
	return observable.notify(event , payload)
}

func (observable *Observable) notify(event string, payload interface{}) interface{} {

	var finalResult interface{}
	for _,handlers := range observable.handlers{
		finalResult = (*handlers).Handle(payload)
	}

	return finalResult
}




