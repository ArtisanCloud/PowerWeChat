package kernel

import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel/contract"
	"github.com/ArtisanCloud/power-wechat/src/kernel/messages"
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

func (serverGuard *ServerGuard) Serve() (response *http.Response, err error) {

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

	request := (*serverGuard.App).GetComponent("ExternalRequest").(*http.Request)
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

func (serverGuard *ServerGuard) getMessage() (dataset *object.HashMap, err error) {
	dataset = &object.HashMap{}
	request := (*serverGuard.App).GetComponent("ExternalRequest").(*http.Request)
	if request == nil {
		return nil, errors.New("request is invalid")
	}
	b, err := io.ReadAll(request.Body)
	if err != nil || b == nil {
		return nil, err
	}

	message, err := serverGuard.parseMessage(string(b))
	if err != nil {
		return nil, err
	}

	if serverGuard.IsSafeMode() && message["Encrypt"] != nil {
		decryptMessage := serverGuard.decryptMessage(string(b), &message)
		err = json.Unmarshal([]byte(decryptMessage), dataset)
		if err == nil && dataset != nil {
			return dataset, err
		}

		//err = xml.Unmarshal([]byte(decryptMessage), &dataset)
		*dataset, err = object.Xml2Map([]byte(decryptMessage))
		if err == nil && dataset != nil {
			return dataset, err
		}
	}

	config := (*serverGuard.App).GetConfig()
	responseType := config.Get("response_type", "array").(string)
	arrayResponse, err := serverGuard.DetectAndCastResponseToType(&message, responseType)
	dataset = arrayResponse.(*object.HashMap)

	return dataset, err

}

func (serverGuard *ServerGuard) resolve() (response *http.Response, err error) {
	result, err := serverGuard.handleRequest()
	if err != nil {
		return nil, err
	}

	if serverGuard.ShouldReturnRawResponse() {
		response = &http.Response{
			Body: ioutil.NopCloser(bytes.NewBufferString((*result)["response"].(string))),
		}

	} else {
		strBuiltResponse := serverGuard.buildResponse((*result)["to"].(string), (*result)["from"].(string), (*result)["response"])
		header := http.Header{}
		header.Set("Content-Type", "application/xml")
		response = &http.Response{
			Body:       ioutil.NopCloser(bytes.NewBufferString(strBuiltResponse)),
			StatusCode: 200,
			Header:     header,
		}
	}
	return response, nil
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

	castedMessage, err := serverGuard.getMessage()
	if err != nil {
		return nil, err
	}

	typeData, err := serverGuard.DetectAndCastResponseToType(castedMessage, "array")
	if err != nil {
		return nil, err
	}
	messageArray := *(typeData.(*object.HashMap))

	var messageType = "text"
	if messageArray["MsgType"] != nil {
		messageType = messageArray["MsgType"].(string)
	}

	response := serverGuard.Dispatch(MESSAGE_TYPE_MAPPING[messageType], castedMessage)

	var (
		strFromUserName string = ""
		strToUserName   string = ""
	)
	if messageArray["FromUserName"] != nil {
		strFromUserName = messageArray["FromUserName"].(string)
	}
	if messageArray["ToUserName"] != nil {
		strToUserName = messageArray["ToUserName"].(string)
	}

	return &object.HashMap{
		"to":       strFromUserName,
		"from":     strToUserName,
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
	request := (*serverGuard.App).GetComponent("ExternalRequest").(*http.Request)
	if request == nil {
		println("request is invalid")
		return false
	}
	query := request.URL.Query()

	return query.Get("signature") == "" && "aes" == query.Get("encrypt_type")

}

func (serverGuard *ServerGuard) parseMessage(content string) (dataContent object.HashMap, err error) {

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

func (serverGuard *ServerGuard) decryptMessage(content string, message *object.HashMap) (decryptMessage string) {

	encryptor := (*serverGuard.App).GetComponent("Encryptor").(*Encryptor)
	request := (*serverGuard.App).GetComponent("ExternalRequest").(*http.Request)
	query := request.URL.Query()
	buf, err := encryptor.Decrypt(
		[]byte(content),
		query.Get("msg_signature"),
		query.Get("nonce"),
		query.Get("timestamp"),
	)
	if err != nil {
		fmt.Println("decrypt message error:", err.ErrMsg)
	}
	decryptMessage = string(buf)

	return decryptMessage
}
