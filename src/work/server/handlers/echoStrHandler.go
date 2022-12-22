package handlers

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/contract"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/decorators"
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

func (handler *EchoStrHandler) Handle(request *http.Request, header contract.EventInterface, content interface{}) interface{} {

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
			panic(err.ErrMsg)
			return nil
		}

		return decorators.NewFinallyResult(string(str))
	}

	return nil
}
