package kernel

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel"
)

type ApplicationPaymentInterface interface {

	GetContainer() *kernel.ServiceContainer
	GetConfig() *kernel.Config
	GetComponent(name string) interface{}
	Scheme(productID string) string
	CodeUrlScheme(codeUrl string) string
	SetSubMerchant(mchId string, appId string) ApplicationPaymentInterface
	HandlePaidNotify(closure func(payload ...interface{}) interface{})
	HandleRefundedNotify(closure func(payload ...interface{}) interface{})
	HandleScannedNotify(closure func(payload ...interface{}) interface{})
	InSandbox() bool
	GetKey(endpoint string) (string, error)
}
