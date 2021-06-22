package handlers

import (
	"bytes"
	"github.com/ArtisanCloud/go-wechat/src/kernel"
	"github.com/ArtisanCloud/go-wechat/src/kernel/contract"
	"net/http"
)

type EchoStrHandler struct {
	contract.EventHandlerInterface

	App *kernel.ApplicationInterface
}

func NewEchoStrHandler(app *kernel.ApplicationInterface) *EchoStrHandler {
	handler := &EchoStrHandler{
		App: app,
	}

	return handler
}

func (handler *EchoStrHandler) handle(payload interface{}) interface{} {

	request := (*handler.App).GetComponent("ExternalRequest").(*http.Request)
	encryptor := (*handler.App).GetComponent("Encryptor").(*kernel.Encryptor)

	query := request.URL.Query()

	strDecrypted := query.Get("echostr")
	if strDecrypted != "" {
		decrypted:= bytes.NewBufferString(strDecrypted).Bytes()
		str, _ := encryptor.Decrypt(
			decrypted,
			query.Get("msg_signature"),
			query.Get("nonce"),
			query.Get("timestamp"),
		)
		return str
	}

	return nil
}
