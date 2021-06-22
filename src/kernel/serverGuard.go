package kernel

import (
	"bytes"
	"crypto/sha1"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-wechat/src/kernel/contract"
	"github.com/ArtisanCloud/go-wechat/src/kernel/messages"
	"github.com/ArtisanCloud/go-wechat/src/kernel/support"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

const SUCCESS_EMPTY_RESPONSE = "success"

var MESSAGE_TYPE_MAPPING = map[string]int{
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

	alwaysValidate bool
	App            *ApplicationInterface
}

func NewServerGuard(app *ApplicationInterface) *ServerGuard {

	serverGuard := &ServerGuard{
		Observable: support.NewObservable(),
		App:        app,
	}

	return serverGuard

}

func (serverGuard *ServerGuard) Serve() (response *http.Response, err error) {

	validateGuard, err := serverGuard.validate()
	if err != nil {
		return nil, err
	}

	response, err = validateGuard.resolve()

	return response, err
}

func (serverGuard *ServerGuard) validate() (*ServerGuard, error) {

	if !serverGuard.alwaysValidate && serverGuard.isSafeMode() {
		return serverGuard, nil
	}
	request := (*serverGuard.App).GetComponent("ExternalRequest").(*http.Request)
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
	b, err := io.ReadAll(request.Body)
	if err != nil || b != nil {
		return nil, err
	}

	message, err := serverGuard.parseMessage(string(b))
	if err != nil {
		return nil, err
	}

	if serverGuard.isSafeMode() && (*message)["Encrypt"] != nil {
		decryptMessage := serverGuard.decryptMessage(message)
		err = json.Unmarshal([]byte(decryptMessage), dataset)
		if err == nil && dataset != nil {
			return dataset, err
		}

		err = xml.Unmarshal([]byte(decryptMessage), &dataset)
		if err == nil && dataset != nil {
			return dataset, err
		}
	}
	return nil, nil

}

func (serverGuard *ServerGuard) resolve() (response *http.Response, err error) {
	result, err := serverGuard.handleRequest()
	if err != nil {
		return nil, err
	}

	if serverGuard.shouldReturnRawResponse() {
		response = &http.Response{
			Body: ioutil.NopCloser(bytes.NewBufferString((*result)["response"].(string))),
		}

	} else {
		strBuiltResponse := serverGuard.buildResponse((*result)["to"].(string), (*result)["from"].(string), (*result)["response"].(string))
		header := http.Header{}
		header.Set("Content-Type", "application/xml")
		response = &http.Response{
			Body:       ioutil.NopCloser(bytes.NewBufferString(strBuiltResponse)),
			StatusCode: 200,
			Header:     header,
		}
	}

}

func (serverGuard *ServerGuard) getToken() string {
	return ""
}

func (serverGuard *ServerGuard) buildResponse(to string, from string, message contract.MessageInterface) string {
	return ""
}

func (serverGuard *ServerGuard) handleRequest() (*object.HashMap, error) {

	castedMessage, err := serverGuard.getMessage()
	if err != nil {
		return nil, err
	}

	// tbd
	messageArray := serverGuard.detectAndCastResponseToType()

	response := serverGuard.Dispatch("", castedMessage)

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
	},nil
}

func (serverGuard *ServerGuard) buildReply(to string, from string, message contract.MessageInterface) string {
	return ""
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
	query := request.URL.Query()

	return query.Get("signature") == "" && "aes" == query.Get("encrypt_type")

}

func (serverGuard *ServerGuard) parseMessage(content string) (dataContent *object.HashMap, err error) {

	if content[0] == '<' {
		err = xml.Unmarshal([]byte(content), &dataContent)
		if err != nil {
			return nil, err
		} else {
			// Handle JSON format.
			err = json.Unmarshal([]byte(content), &dataContent)
			if err != nil {
				return nil, err
			}
		}
	}
	return dataContent, err
}
func (serverGuard *ServerGuard) shouldReturnRawResponse() bool {
	return false
}

func (serverGuard *ServerGuard) decryptMessage(message *object.HashMap) (decryptMessage string) {

	encryptor := (*serverGuard.App).GetComponent("Encryptor").(*Encryptor)
	request := (*serverGuard.App).GetComponent("request").(*http.Request)
	query := request.URL.Query()
	buf, err := encryptor.Decrypt(
		[]byte((*message)["Encrypt"].(string)),
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
