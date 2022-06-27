package miniProgram

import (
	"github.com/ArtisanCloud/PowerLibs/v2/logger"
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/providers"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/miniProgram"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/openPlatform/authorizer/aggregate"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/openPlatform/authorizer/aggregate/account"
)

type Application struct {
	*miniProgram.MiniProgram

	Account *account.Client
}

func NewApplication(config *miniProgram.UserConfig, extraInfos ...*kernel.ExtraInfo) (*Application, error) {

	var extraInfo = &kernel.ExtraInfo{}
	if len(extraInfos) > 0 {
		extraInfo = extraInfos[0]
	}

	miniProgram, err := miniProgram.NewMiniProgram(config, extraInfo)
	if err != nil {
		return nil, err
	}
	app := &Application{
		MiniProgram: miniProgram,
	}

	//-------------- global app config --------------
	// global app config
	app.Config = providers.RegisterConfigProvider(app)

	//-------------- register Aggregate --------------
	app.Account, err = aggregate.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	app.Logger, err = logger.NewLogger("", &object.HashMap{
		"env":        app.Config.GetString("env", "develop"),
		"outputPath": app.Config.GetString("file", "./wechat.log"),
		"errorPath":  app.Config.GetString("file", "./wechat.log"),
	})
	if err != nil {
		return nil, err
	}

	return app, err
}
