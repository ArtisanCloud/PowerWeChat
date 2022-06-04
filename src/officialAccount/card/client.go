package card

import "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"

type Client struct {
	*kernel.BaseClient

	URL               string
	TicketCacheKey    string
	TicketCachePrefix string
}

func NewClient(app kernel.ApplicationInterface) *Client {
	client := &Client{
		BaseClient: kernel.NewBaseClient(&app, nil),
	}

	client.TicketCachePrefix = "powerwechat.official_account.card.api_ticket."

	return client
}

func (comp *Client) Colors() {

}
