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
	"github.com/ArtisanCloud/power-wechat/src/payment/notify/request"
	"io/ioutil"
	"net/http"
)

type Handler struct {
	App        *kernel.ApplicationPaymentInterface
	Message    *request.RequestNotify
	fail       string
	Attributes *object.StringMap
	Check      bool
	Sign       bool

	ExternalRequest *http.Request

	//Handle func(closure func(message *request.RequestNotify, transaction *models.Transaction, fail func(groupWelcomeTemplate string)) interface{}) *http.Response
}

const SUCCESS = "SUCCESS"
const FAIL = "FAIL"

func NewHandler(app *kernel.ApplicationPaymentInterface, r *http.Request) *Handler {

	//-------------- external request --------------
	request := &http.Request{}
	if r != nil {
		request = r
	}

	return &Handler{
		App:             app,
		Check:           true,
		Sign:            false,
		ExternalRequest: request,
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
func (handler *Handler) ToResponse() (response *response2.HttpResponse, err error) {

	returnCode := SUCCESS
	returnMsg := "成功"
	if handler.fail != "" {
		returnCode = FAIL
		returnMsg = handler.fail
		err = errors.New(handler.fail)
	}
	base := &object.StringMap{
		"code":           returnCode,
		"uniformMessage": returnMsg,
	}

	attributes := object.MergeStringMap(base, handler.Attributes)
	baseClient := (*handler.App).GetComponent("Base").(*base2.Client)
	if handler.Sign {
		(*attributes)["sign"], err = baseClient.Signer.GenerateSign("")
		if err != nil {
			return nil, err
		}
	}

	bodyBuffer, _ := json.Marshal(attributes)
	rs := response2.NewHttpResponse(http.StatusOK)
	rs.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBuffer))

	return rs, err
}

func (handler *Handler) GetMessage() (notify *request.RequestNotify, err error) {

	if handler.Message != nil {
		return handler.Message, nil
	}

	externalRequest := handler.ExternalRequest

	requestBody, err := ioutil.ReadAll(externalRequest.Body)
	if err!=nil{
		return nil, err
	}
	handler.Message = &request.RequestNotify{}
	err = object.JsonDecode(requestBody, handler.Message)
	if err != nil {
		return nil, err
	}
	handler.Message.RawRequest = externalRequest

	return handler.Message, nil
}

func (handler *Handler) DecryptMessage() (string, error) {
	message, err := handler.GetMessage()
	if err != nil {
		return "", err
	}
	if message.Resource == nil {
		return "", errors.New("uniformMessage doesn't have the key value")
	}

	config := (*handler.App).GetConfig()
	wxKey := config.GetString("mch_api_v3_key", "")
	nonce := message.Resource.Nonce
	associatedData := message.Resource.AssociatedData
	cipherText := message.Resource.Ciphertext
	return support.DecryptAES256GCM(
		wxKey,
		associatedData,
		nonce,
		cipherText,
	)

}

func (handler *Handler) Strict(result interface{}) {

	bResult := true
	strResult := ""
	switch result.(type) {
	case bool:
		bResult = result.(bool)
		strResult = fmt.Sprintf("%t", result)
	case string:
		strResult = result.(string)
		if strResult != "" {
			bResult = false
		}
	default:
		return
	}

	if bResult != true && handler.fail == "" {
		handler.Fail(strResult)
	}
}


func (handler *Handler) reqInfo() (content string, err error) {

	content, err = handler.DecryptMessage()
	if err != nil {
		return "", err
	}

	// save the decoded content to message resource
	handler.Message.Resource.Plaintext = content

	return content, nil
}
