package handlers

type ServerCallbackHandler struct {
	Callback func(payload interface{}) interface{}
}

func NewServerCallbackHandler()*ServerCallbackHandler{
	return &ServerCallbackHandler{}
}

func (handler *ServerCallbackHandler) Handle(payload interface{}) interface{} {

	if handler.Callback!=nil{
		return handler.Callback(payload)
	}

	return nil
}