package handlers

import (
	"github.com/ArtisanCloud/go-wechat/src/kernel"
	"github.com/ArtisanCloud/go-wechat/src/kernel/contract"
	"github.com/ArtisanCloud/go-wechat/src/kernel/decorators"
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

func (handler *EchoStrHandler) Handle(payload ...interface{}) interface{} {

	request := (*handler.App).GetComponent("ExternalRequest").(*http.Request)
	if request == nil {
		println("request is invalid")
		return nil
	}
	encryptor := (*handler.App).GetComponent("Encryptor").(*kernel.Encryptor)

	query := request.URL.Query()
	strDecrypted := query.Get("echostr")

	//strDecrypted := "wAHjjtg5VmVwpLzy7YoJXGT1SAVvFACAZbThfdeq2bWrhDCfDv6o9JWMImq2N7SPKiG8Pci+2W7pP9GlpELVEA=="
	//strDecrypted = url.QueryEscape(strDecrypted)
	//fmt.Dump(request.RequestURI, strDecrypted)
	if strDecrypted != "" {
		str, err := encryptor.VerifyUrl(
			strDecrypted,
			query.Get("msg_signature"),
			query.Get("nonce"),
			query.Get("timestamp"),
		)

		if err != nil {
			println(err.ErrCode, err.ErrMsg)
			return nil
		}

		return decorators.NewFinallyResult(string(str))
	}

	return nil
}
