package officialAccount

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/providers"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/officialAccount"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/openPlatform/authorizer/officialAccount/account"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/openPlatform/authorizer/officialAccount/miniProgram"
)

type Application struct {
	*officialAccount.OfficialAccount

	Account *account.Client

	MiniProgram *miniProgram.Client
}

func NewApplication(config *officialAccount.UserConfig, extraInfos ...*kernel.ExtraInfo) (*Application, error) {

	var extraInfo, _ = kernel.NewExtraInfo()
	if len(extraInfos) > 0 {
		extraInfo = extraInfos[0]
	}

	officialAccount, err := officialAccount.NewOfficialAccount(config, extraInfo)
	if err != nil {
		return nil, err
	}
	app := &Application{
		OfficialAccount: officialAccount,
	}

	//-------------- global app config --------------
	// global app config
	app.Config = providers.RegisterConfigProvider(app)

	//-------------- register Aggregate --------------
	app.Account, err = account.NewClient(app, nil)
	if err != nil {
		return nil, err
	}

	//-------------- register MiniProgram --------------
	app.MiniProgram, err = miniProgram.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	return app, err
}
