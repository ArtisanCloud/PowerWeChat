package contract

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-wechat/src/kernel"
)

type ApplicationInterface interface {

	GetConfig() []object.HashMap

	GetComponents() *object.HashMap

	GetContainer() *kernel.ServiceContainer
}
