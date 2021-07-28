package notify

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	response2 "github.com/ArtisanCloud/go-libs/http/response"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-libs/str"
	"github.com/ArtisanCloud/power-wechat/src/kernel/support"
	"github.com/ArtisanCloud/power-wechat/src/payment/kernel"
	"io/ioutil"
	"net/http"
)

type Handler struct {
	App        *kernel.ApplicationPaymentInterface
	Message    *object.HashMap
	Fail       string
	Attributes *object.StringMap
	Check      bool
	Sign       bool

	Handle func(closure func(payload ...interface{}) interface{}) *http.Response
}

const SUCCESS = "SUCCESS"
const FAIL = "FAIL"

func NewHandler(app *kernel.ApplicationPaymentInterface) *Handler {
	return &Handler{
		App:   app,
		Check: true,
		Sign:  false,
	}
}

func (handler *Handler) fail(message string) {
	handler.Fail = message
}

func (handler *Handler) RespondWith(attributes *object.StringMap, sign bool) *Handler {

	handler.Attributes = attributes
	handler.Sign = sign

	return handler
}
func (handler *Handler) ToResponse() (response *http.Response, err error) {

	returnCode := SUCCESS
	if handler.Fail != "" {
		returnCode = FAIL
	}
	base := &object.StringMap{
		"return_code": returnCode,
		"return_msg":  handler.Fail,
	}

	attributes := object.MergeStringMap(base, handler.Attributes)
	baseClient := (*handler.App).GetComponent("Base").(*kernel.BaseClient)
	if handler.Sign {
		(*attributes)["sign"], err = baseClient.Signer.GenerateSign(attributes)
		if err != nil {
			return nil, err
		}
	}

	bodyBuffer, _ := json.Marshal(attributes)
	rs := response2.NewHttpResponse()
	rs.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBuffer))

	return rs.Response, nil
}

func (handler *Handler) GetMessage() (message *object.HashMap, err error) {

	if handler.Message != nil {
		return handler.Message, nil
	}

	externalRequest := (*handler.App).GetComponent("ExternalRequest").(*http.Request)
	requestBody, _ := ioutil.ReadAll(externalRequest.Body)
	err = object.JsonDecode(requestBody, handler.Message)
	if err != nil {
		return nil, err
	}

	return handler.Message, nil
}

func (handler *Handler) DecryptMessage(key string) (string, error) {
	message, err := handler.GetMessage()
	if err != nil {
		return "", err
	}
	if (*message)[key] == nil {
		return "", errors.New("message doesn't have the key value")
	}
	content := (*message)[key].(string)

	config := (*handler.App).GetConfig()
	wxKey := config.GetString("key", "")
	nonce := str.QuickRandom(12)
	return support.DecryptAES256GCM(
		wxKey,
		"",
		nonce,
		content,
	)

}

func (handler *Handler) Strict(result bool) {

	if result != true && handler.Fail == "" {
		handler.fail(fmt.Sprintf("%t", result))
	}
}
