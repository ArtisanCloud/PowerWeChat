package contract

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-wechat/src/kernel"
)

type ApplicationInterface interface {

	GetComponents() *object.HashMap

	GetContainer() *kernel.ServiceContainer
}
