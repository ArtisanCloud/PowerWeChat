package kernel

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"net/http"
)

type ApplicationPaymentInterface interface {
	GetContainer() *kernel.ServiceContainer
	GetConfig() *kernel.Config
	GetComponent(name string) interface{}

	Scheme(productID string) string
	CodeUrlScheme(codeUrl string) string
	SetSubMerchant(mchId string, appId string) ApplicationPaymentInterface

	HandlePaidNotify(request *http.Request, closure func(message *object.HashMap, content *object.HashMap, fail string) interface{}) (*http.Response, error)
	HandleRefundedNotify(request *http.Request, closure func(message *object.HashMap, content *object.HashMap, fail string) interface{}) (*http.Response, error)
	HandleScannedNotify(request *http.Request, closure func(message *object.HashMap, content *object.HashMap, fail string, alert string) interface{}) (*http.Response, error)

	InSandbox() bool
	GetKey(endpoint string) (string, error)
}
