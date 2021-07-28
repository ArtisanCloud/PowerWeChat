package notify

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	response2 "github.com/ArtisanCloud/go-libs/http/response"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel/support"
	base2 "github.com/ArtisanCloud/power-wechat/src/payment/base"
	"github.com/ArtisanCloud/power-wechat/src/payment/kernel"
	"io/ioutil"
	"net/http"
)

type Handler struct {
	App        *kernel.ApplicationPaymentInterface
	Message    *object.HashMap
	fail       string
	Attributes *object.StringMap
	Check      bool
	Sign       bool

	Handle func(closure func(message *object.HashMap, content *object.HashMap, fail string) interface{}) *http.Response
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

func (handler *Handler) Fail(message string) {
	handler.fail = message
}

func (handler *Handler) RespondWith(attributes *object.StringMap, sign bool) *Handler {

	handler.Attributes = attributes
	handler.Sign = sign

	return handler
}
func (handler *Handler) ToResponse() (response *http.Response, err error) {

	returnCode := SUCCESS
	returnMsg:="成功"
	if handler.fail != "" {
		returnCode = FAIL
	}
	base := &object.StringMap{
		"code": returnCode,
		"message":  returnMsg,
	}

	attributes := object.MergeStringMap(base, handler.Attributes)
	baseClient := (*handler.App).GetComponent("Base").(*base2.Client)
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
	handler.Message = &object.HashMap{}
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
	content := (*message)[key].(map[string]interface{})
	config := (*handler.App).GetConfig()
	wxKey := config.GetString("key", "")
	nonce := content["nonce"].(string)
	associatedData := content["associated_data"].(string)
	cipherText := content["ciphertext"].(string)
	return support.DecryptAES256GCM(
		wxKey,
		associatedData,
		nonce,
		cipherText,
	)

}

func (handler *Handler) Strict(result interface{}) {

	bResult := false
	strResult := ""
	switch result.(type) {
	case bool:
		bResult = result.(bool)
		strResult = fmt.Sprintf("%t", result)
	case string:
		strResult = result.(string)
		if strResult != "" {
			bResult = true
		}
	default:
		return
	}

	if bResult != true && handler.fail == "" {
		handler.Fail(strResult)
	}
}
