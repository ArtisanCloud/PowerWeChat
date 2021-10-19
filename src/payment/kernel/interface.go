package kernel

import "github.com/ArtisanCloud/PowerWeChat/src/kernel"

type ApplicationPaymentInterface interface {
	kernel.ApplicationInterface
	//GetContainer() *kernel.ServiceContainer
	//GetConfig() *kernel.Config
	//GetComponent(name string) interface{}

	Scheme(productID string) string
	CodeUrlScheme(codeUrl string) string
	SetSubMerchant(mchId string, appId string) ApplicationPaymentInterface

	InSandbox() bool
	GetKey(endpoint string) (string, error)
}
