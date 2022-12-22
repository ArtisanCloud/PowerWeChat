package miniProgram

import (
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/providers"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/auth"
	auth2 "github.com/ArtisanCloud/PowerWeChat/v3/src/work/miniProgram/auth"
)

type Application struct {
	*miniProgram.MiniProgram

	AccessToken *auth.AccessToken
	Auth        *auth2.Client
}

type UserConfig struct {
	MiniProgramUserConfig *miniProgram.UserConfig
	CorpID                string
	AgentID               int
	Secret                string
	Token                 string
	AESKey                string
	CallbackURL           string
	Http                  *object.HashMap
}

func NewApplication(config *UserConfig, extraInfos ...*kernel.ExtraInfo) (*Application, error) {

	var extraInfo, _ = kernel.NewExtraInfo()
	if len(extraInfos) > 0 {
		extraInfo = extraInfos[0]
	}

	miniProgram, err := miniProgram.NewMiniProgram(config.MiniProgramUserConfig, extraInfo)
	if err != nil {
		return nil, err
	}
	app := &Application{
		MiniProgram: miniProgram,
	}

	// ****** manually set corp_id and secret  for miniProgram NewAccessToken
	userConfig := (*app.MiniProgram.ServiceContainer.UserConfig)
	userConfig["corp_id"] = config.CorpID
	userConfig["secret"] = config.Secret
	userConfig["http"] = config.Http
	// global app config
	app.Config = providers.RegisterConfigProvider(app)

	app.AccessToken, err = auth.NewAccessToken(app)
	if err != nil {
		return nil, err
	}

	app.Auth, err = auth2.NewClient(app)
	if err != nil {
		return nil, err
	}

	return app, err
}

func (app *Application) GetContainer() *kernel.ServiceContainer {
	return app.ServiceContainer
}

func (app *Application) GetAccessToken() *kernel.AccessToken {
	return app.AccessToken.AccessToken
}

func (app *Application) GetConfig() *kernel.Config {
	return app.Config
}

func (app *Application) GetComponent(name string) interface{} {

	switch name {
	case "Base":
		return app.Base
	case "AccessToken":
		return app.AccessToken
	case "Auth":
		return app.Auth
	case "Config":
		return app.Config

	default:
		return nil
	}

}
