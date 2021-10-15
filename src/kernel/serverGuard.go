package kernel

import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"encoding/xml"
	"errors"
	"github.com/ArtisanCloud/go-libs/http/response"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel/contract"
	"github.com/ArtisanCloud/power-wechat/src/kernel/messages"
	"github.com/ArtisanCloud/power-wechat/src/kernel/models"
	"github.com/ArtisanCloud/power-wechat/src/kernel/support"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

const SUCCESS_EMPTY_RESPONSE = "success"

var MESSAGE_TYPE_MAPPING = map[string]int{
	"*":               messages.VOID,
	"text":            messages.TEXT,
	"image":           messages.IMAGE,
	"voice":           messages.VOICE,
	"video":           messages.VIDEO,
	"shortvideo":      messages.SHORT_VIDEO,
	"location":        messages.LOCATION,
	"link":            messages.LINK,
	"device_event":    messages.DEVICE_EVENT,
	"device_text":     messages.DEVICE_TEXT,
	"event":           messages.EVENT,
	"file":            messages.FILE,
	"miniprogrampage": messages.MINIPROGRAM_PAGE,
}

type ServerGuard struct {
	*support.Observable
	*support.ResponseCastable

	alwaysValidate bool
	App            *ApplicationInterface

	IsSafeMode              func() bool
	Validate                func() (*ServerGuard, error)
	ShouldReturnRawResponse func() bool

	ToCallbackType func(callbackHeader contract.EventInterface, buf []byte) (decryptMessage interface{}, err error)
}

func NewServerGuard(app *ApplicationInterface) *ServerGuard {

	serverGuard := &ServerGuard{
		Observable: support.NewObservable(),
		App:        app,
	}

	serverGuard.IsSafeMode = func() bool {
		return serverGuard.isSafeMode()
	}
	serverGuard.Validate = func() (*ServerGuard, error) {
		return serverGuard.validate()
	}
	serverGuard.ShouldReturnRawResponse = func() bool {
		return serverGuard.shouldReturnRawResponse()
	}

	return serverGuard

}

func (serverGuard *ServerGuard) Serve(r *http.Request) (response *response.HttpResponse, err error) {

	//-------------- external request --------------
	request := &http.Request{}
	if r != nil {
		request = r
	}

	(*serverGuard.App).SetExternalRequest(request)

	validatedGuard, err := serverGuard.Validate()
	if err != nil {
		return nil, err
	}

	response, err = validatedGuard.resolve()

	// tbd
	// log here

	return response, err
}

func (serverGuard *ServerGuard) validate() (*ServerGuard, error) {

	if !serverGuard.alwaysValidate && serverGuard.IsSafeMode() {
		return serverGuard, nil
	}

	request := (*serverGuard.App).GetExternalRequest()
	if request == nil {
		return nil, errors.New("request is invalid")
	}
	query := request.URL.Query()

	sign := serverGuard.signature([]string{
		serverGuard.getToken(),
		query.Get("timestamp"),
		query.Get("nonce"),
	})

	if query.Get("signature") != sign {
		return serverGuard, errors.New("invalid request signature")
	}

	return serverGuard, nil
}

func (serverGuard *ServerGuard) getMessage() (callback *models.Callback, callbackHeader *models.CallbackMessageHeader, Decrypted interface{}, err error) {

	request := (*serverGuard.App).GetExternalRequest()
	if request == nil {
		return nil, nil, nil, errors.New("request is invalid")
	}
	b, err := io.ReadAll(request.Body)
	if err != nil || b == nil {
		return nil, nil, nil, err
	}

	callback, err = serverGuard.parseMessage(string(b))
	if err != nil {
		return nil, nil, nil, err
	}

	if serverGuard.IsSafeMode() && callback.Encrypt != "" {
		callbackHeader, Decrypted, err = serverGuard.decryptMessage(string(b))
	}

	return callback, callbackHeader, Decrypted, err

}

func (serverGuard *ServerGuard) resolve() (httpRS *response.HttpResponse, err error) {
	result, err := serverGuard.handleRequest()
	if err != nil {
		return nil, err
	}

	var rs *http.Response
	if serverGuard.ShouldReturnRawResponse() {
		rs = &http.Response{
			Body: ioutil.NopCloser(bytes.NewBufferString((*result)["response"].(string))),
		}

	} else {
		strBuiltResponse := serverGuard.buildResponse((*result)["to"].(string), (*result)["from"].(string), (*result)["response"])
		header := http.Header{}
		header.Set("Content-Type", "application/xml")
		rs = &http.Response{
			Body:       ioutil.NopCloser(bytes.NewBufferString(strBuiltResponse)),
			StatusCode: http.StatusOK,
			Header:     header,
		}
	}
	httpRS = response.NewHttpResponse(http.StatusOK)
	httpRS.Response = rs

	return httpRS, nil
}

func (serverGuard *ServerGuard) getToken() string {
	config := (*serverGuard.App).GetConfig()
	token := config.Get("token", "").(string)

	return token
}

func (serverGuard *ServerGuard) buildResponse(to string, from string, message interface{}) string {

	var toMessage contract.MessageInterface
	switch message.(type) {
	case nil:
		return SUCCESS_EMPTY_RESPONSE

	case string:
		strMessage := message.(string)
		if SUCCESS_EMPTY_RESPONSE == strMessage {
			return SUCCESS_EMPTY_RESPONSE
		} else {
			toMessage = messages.NewText(message.(string))
			break
		}
	case int:
		toMessage = messages.NewText(strconv.Itoa(message.(int)))
		break

	case messages.Raw:
		return message.(messages.Raw).Get("content", SUCCESS_EMPTY_RESPONSE).(string)

	case object.HashMap:
		toMessage = messages.NewNews(message.(*object.HashMap))
		break
	case *object.HashMap:
		toMessage = messages.NewNews(message.(*object.HashMap))
		break
	default:

	}

	return serverGuard.buildReply(to, from, toMessage)
}

func (serverGuard *ServerGuard) handleRequest() (*object.HashMap, error) {

	_, msgHeader, decryptedMessage, err := serverGuard.getMessage()
	if err != nil {
		return nil, err
	}

	var messageType = "text"
	if msgHeader.MsgType != "" {
		messageType = msgHeader.MsgType
	}

	response := serverGuard.Dispatch(MESSAGE_TYPE_MAPPING[messageType], msgHeader, decryptedMessage)

	return &object.HashMap{
		"to":       msgHeader.FromUserName,
		"from":     msgHeader.ToUserName,
		"response": response,
	}, nil
}

func (serverGuard *ServerGuard) buildReply(to string, from string, message contract.MessageInterface) (response string) {

	prepends := &object.HashMap{
		"ToUserName":   to,
		"FromUserName": from,
		"CreateTime":   time.Now(),
		"MsgType":      message.GetType(),
	}
	transformedResponse, _ := message.TransformToXml(prepends, false)
	response = transformedResponse.(string)
	if serverGuard.IsSafeMode() {
		// tbd log here
		encryptor := (*serverGuard.App).GetComponent("Encryptor").(*Encryptor)
		encryptedResponse, err := encryptor.Encrypt(response, "", "")
		if err == nil {

			response = string(encryptedResponse)
		} else {
			// tbd log here
			println("encryptor error: ", err.ErrMsg)
		}

	}

	return response
}

func (serverGuard *ServerGuard) signature(params []string) string {
	sort.Strings(params)
	strJoined := strings.Join(params, "")

	hash := sha1.New()
	hash.Write([]byte(strJoined))
	bs := hash.Sum(nil)

	return string(bs)
}

func (serverGuard *ServerGuard) isSafeMode() bool {
	request := (*serverGuard.App).GetExternalRequest()
	if request == nil {
		println("request is invalid")
		return false
	}
	query := request.URL.Query()

	return query.Get("signature") == "" && "aes" == query.Get("encrypt_type")

}

func (serverGuard *ServerGuard) parseMessage(content string) (callback *models.Callback, err error) {

	callback = &models.Callback{}
	err = xml.Unmarshal([]byte(content), callback)

	return callback, err
}

func (serverGuard *ServerGuard) parseMessage2(content string) (dataContent object.HashMap, err error) {

	dataContent = nil
	if content != "" {
		// check xml format
		if content[0] == '<' {
			dataContent = object.HashMap{}
			//err = xml.Unmarshal([]byte(content), &dataContent)
			dataContent, err = object.Xml2Map([]byte(content))
			if err != nil {
				return nil, err
			}
		}
	}

	if dataContent == nil {
		// Handle JSON format.
		dataContent = object.HashMap{}
		err = json.Unmarshal([]byte(content), &dataContent)
		if err != nil {
			return dataContent, nil
		}
	}

	return dataContent, err
}
func (serverGuard *ServerGuard) shouldReturnRawResponse() bool {
	return false
}

func (serverGuard *ServerGuard) decryptMessage(content string) (callbackHeader *models.CallbackMessageHeader, decryptMessage interface{}, err error) {

	encryptor := (*serverGuard.App).GetComponent("Encryptor").(*Encryptor)
	request := (*serverGuard.App).GetExternalRequest()
	query := request.URL.Query()
	buf, cryptErr := encryptor.Decrypt(
		[]byte(content),
		query.Get("msg_signature"),
		query.Get("nonce"),
		query.Get("timestamp"),
	)
	if cryptErr != nil {
		return nil, nil, errors.New(cryptErr.ErrMsg)
	}

	callbackHeader = &models.CallbackMessageHeader{}
	err = xml.Unmarshal(buf, callbackHeader)
	if err != nil {
		return nil, nil, err
	}

	decryptMessage, err = serverGuard.ToCallbackType(callbackHeader, buf)

	return callbackHeader, decryptMessage, err

}
