package provider

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	suite "github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/suitAuth"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

func NewClient(app *kernel.ApplicationInterface) (*Client, error) {
	token := (*app).GetComponent("ProviderAccessToken").(*suite.AccessToken)
	baseClient, err := kernel.NewBaseClient(app, token.AccessToken)
	if err != nil {
		return nil, err
	}
	return &Client{
		baseClient,
	}, nil
}
