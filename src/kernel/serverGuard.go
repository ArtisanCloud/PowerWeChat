package kernel

import (
	"github.com/ArtisanCloud/go-wechat/src/kernel/messages"
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
	alwaysValidate bool
	App   *ApplicationInterface
}

func NewServerGuard(app *ApplicationInterface) *ServerGuard {
	//config := (*app).GetContainer().GetConfig()
	//
	//client := &BaseClient{
	//	HttpRequest: http.NewHttpRequest(config),
	//	App:         app,
	//	//Token:       token,
	//}
	//return client
	return nil

}