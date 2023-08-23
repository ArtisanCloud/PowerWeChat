package virtualPayment

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
)

type Client struct {
	App *kernel.ApplicationInterface
	*kernel.BaseClient

	appKey  string
	offerId string
}

func NewClient(app *kernel.ApplicationInterface) (*Client, error) {
	config := (*app).GetConfig()

	baseClient, err := kernel.NewBaseClient(app, nil)

	if err != nil {

		return nil, err
	}

	client := &Client{
		BaseClient: baseClient,
		App:        app,
		appKey:     config.GetString("app_key", ""),
		offerId:    config.GetString("offer_id", ""),
	}

	client.OverrideGetMiddlewares()
	client.RegisterHttpMiddlewares()

	return client, nil

}
