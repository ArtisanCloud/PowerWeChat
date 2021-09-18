package kernel

import (
	"github.com/ArtisanCloud/go-libs/http/response"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"net/http"
)

type ApplicationPaymentInterface interface {
	GetContainer() *kernel.ServiceContainer
	GetConfig() *kernel.Config
	GetComponent(name string) interface{}

	Scheme(productID string) string
	CodeUrlScheme(codeUrl string) string
	SetSubMerchant(mchId string, appId string) ApplicationPaymentInterface

	HandlePaidNotify(request *http.Request, closure func(message *power.HashMap, content *power.HashMap, fail func(message string)) interface{}) (*response.HttpResponse, error)
	HandleRefundedNotify(request *http.Request, closure func(message *power.HashMap, content *power.HashMap, fail func(message string)) interface{}) (*response.HttpResponse, error)
	HandleScannedNotify(request *http.Request, closure func(message *power.HashMap, content *power.HashMap, fail func(message string), alert func(message string)) interface{}) (*response.HttpResponse, error)

	InSandbox() bool
	GetKey(endpoint string) (string, error)
}
