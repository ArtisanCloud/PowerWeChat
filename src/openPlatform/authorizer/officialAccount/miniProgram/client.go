package miniProgram

import "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"

type Client struct {
	*kernel.BaseClient

	AllowTypes []string
}

func NewClient(app *kernel.ApplicationInterface) (*Client, error) {
	baseClient, err := kernel.NewBaseClient(app, nil)
	if err != nil {
		return nil, err
	}
	return &Client{
		BaseClient: baseClient,
		AllowTypes: []string{"image", "voice", "video", "thumb", "news_image"},
	}, nil
}
