package miniProgram

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/providers"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/miniProgram"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/officialAccount/material"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/openPlatform/authorizer/miniProgram/account"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/openPlatform/authorizer/miniProgram/code"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/openPlatform/authorizer/miniProgram/domain"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/openPlatform/authorizer/miniProgram/privacy"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/openPlatform/authorizer/miniProgram/setting"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/openPlatform/authorizer/miniProgram/tester"
)

type Application struct {
	*miniProgram.MiniProgram

	Account  *account.Client
	Code     *code.Client
	Domain   *domain.Client
	Material *material.Client
	Privacy  *privacy.Client
	Setting  *setting.Client
	Tester   *tester.Client
}

func NewApplication(config *miniProgram.UserConfig, extraInfos ...*kernel.ExtraInfo) (*Application, error) {

	var extraInfo, _ = kernel.NewExtraInfo()
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
	app.Account, err = account.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register Code --------------
	app.Code, err = code.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register Domain --------------
	app.Domain, err = domain.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register Material --------------
	app.Material, err = material.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register Privacy --------------
	app.Privacy, err = privacy.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register Setting --------------
	app.Setting, err = setting.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register Tester --------------
	app.Tester, err = tester.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	return app, err
}
