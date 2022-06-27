package miniProgram

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/miniProgram"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/work/auth"
	auth2 "github.com/ArtisanCloud/PowerWeChat/v2/src/work/miniProgram/auth"
)

type Application struct {
	*miniProgram.MiniProgram

	AccessToken *auth.AccessToken
	Auth        *auth2.Client
}

func NewApplication(config *miniProgram.UserConfig, extraInfos ...*kernel.ExtraInfo) (*Application, error) {

	var extraInfo = &kernel.ExtraInfo{}
	if len(extraInfos) > 0 {
		extraInfo = extraInfos[0]
	}

	(*extraInfo.Prepends)["AccessToken"] = func(app *kernel.ApplicationInterface) (*auth.AccessToken, error) {
		return auth.NewAccessToken(app)
	}
	(*extraInfo.Prepends)["auth"] = func(app *kernel.ApplicationInterface) (*auth2.Client, error) {
		return auth2.NewClient(*app)
	}

	miniProgram, err := miniProgram.NewMiniProgram(config, extraInfo)
	if err != nil {
		return nil, err
	}
	app := &Application{
		MiniProgram: miniProgram,
	}

	return app, err
}
