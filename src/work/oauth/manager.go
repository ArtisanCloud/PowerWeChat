package oauth

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-socialite/src/providers"
	"github.com/ArtisanCloud/go-wechat/src/kernel"
	"net/http"
)

type Manager struct {
	Config   *object.HashMap
	App      *kernel.ApplicationInterface
	Provider *providers.WeCom
}

func NewManager(config *object.HashMap, providerConfig *object.HashMap, app *kernel.ApplicationInterface) *Manager {

	manager := &Manager{
		(*config)["wecom"].(*object.HashMap),
		app,
		providers.NewWeCom(providerConfig),
	}
	return manager

}

func (manager *Manager) Redirect(redirect string) {
	redirectURL := manager.Provider.Redirect(redirect)

	http.Redirect(nil, nil, redirectURL, http.StatusFound)
}
