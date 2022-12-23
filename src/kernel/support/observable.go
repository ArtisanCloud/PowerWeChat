package support

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/contract"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/decorators"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/messages"
	"net/http"
	"reflect"
)

type Observable struct {
	handlers [][]*contract.EventHandlerInterface
}

func NewObservable() *Observable {
	return &Observable{
		make([][]*contract.EventHandlerInterface, 1),
	}
}

func (observable *Observable) PushMessage(closure contract.EventHandlerInterface, condition int) *Observable {

	if observable.handlers[condition] == nil {
		observable.handlers[condition] = []*contract.EventHandlerInterface{}
	}

	observable.handlers[condition] = append(observable.handlers[condition], &closure)

	// tbd
	// clause

	return observable
}

func (observable *Observable) Push(closure contract.EventHandlerInterface, condition int) *Observable {

	if observable.handlers[condition] == nil {
		observable.handlers[condition] = []*contract.EventHandlerInterface{}
	}

	observable.handlers[condition] = append(observable.handlers[condition], &closure)

	// tbd
	// clause

	return observable
}

func (observable *Observable) SetHandlers(handlers [][]*contract.EventHandlerInterface) *Observable {
	observable.handlers = handlers

	return observable
}

func (observable *Observable) Observe(condition int, handler contract.EventHandlerInterface) *Observable {
	return observable.Push(handler, condition)
}

func (observable *Observable) On(condition int, handler contract.EventHandlerInterface) *Observable {
	return observable.Push(handler, condition)
}

func (observable *Observable) Dispatch(request *http.Request, event int, header contract.EventInterface, content interface{}) interface{} {
	return observable.notify(request, event, header, content)
}

func (observable *Observable) notify(request *http.Request, event int, header contract.EventInterface, content interface{}) interface{} {

	var (
		finalResult interface{}
		result      interface{}
		response    interface{}
		//err         error
	)

Loop1:
	for condition, handlers := range observable.handlers {

		if messages.VOID == condition || ((condition & event) == event) {
		Loop2:
			for _, handler := range handlers {
				// tbd
				// intercepted

				response = observable.callHandler(request, handler, header, content)

				switch response.(type) {

				case decorators.TerminateResult:
					return response.(decorators.TerminateResult).Content
				case bool:
					typeValue := response.(bool)
					if typeValue {
						continue Loop2
					} else {
						break Loop1
					}
				case nil:
					break
				default:
					// make sure response is not nil
					if response != nil {
						// result is not init with a single response
						if result == nil {
							result = response
						} else {
							// if current result is not final result
							objType := reflect.TypeOf(result).String()
							if objType != "decorators.FinallyResult" &&
								objType != "*decorators.FinallyResult" {
								result = response
							}
						}
					}

				}
			}
		}
	}

	switch result.(type) {
	case *decorators.FinallyResult:
		finalResult = result.(*decorators.FinallyResult).Content
		break
	case decorators.FinallyResult:
		finalResult = result.(decorators.FinallyResult).Content
		break
	default:
		finalResult = result
	}

	return finalResult
}

func (observable *Observable) callHandler(request *http.Request, callable *contract.EventHandlerInterface, header contract.EventInterface, content interface{}) interface{} {
	return (*callable).Handle(request, header, content)
}
