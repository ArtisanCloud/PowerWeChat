package oauth

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-socialite/src/providers"
	"github.com/ArtisanCloud/go-wechat/src/kernel"
)

type Manager struct {
	Config   *object.HashMap
	App      *kernel.ApplicationInterface
	Provider *providers.WeCom
}

func NewManager(config *object.HashMap,app *kernel.ApplicationInterface) *Manager {

	manager:= &Manager{
		config,
		app,
		nil,
	}
	return manager

}