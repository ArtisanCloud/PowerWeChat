package card

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/officialAccount/card/request"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/officialAccount/card/response"
)

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

func (comp *Client) Categories() {

}

// 创建卡券
// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Create_a_Coupon_Voucher_or_Card.html
func (comp *Client) Create(param *request.RequestCardCreate) (*response.ResponseCardCreate, error) {
	result := &response.ResponseCardCreate{}

	_, err := comp.HttpPostJson("card/create", param, nil, nil, result)

	return result, err

}

// 查看卡券详情
// https://developers.weixin.qq.com/doc/offiaccount/Cards_and_Offer/Managing_Coupons_Vouchers_and_Cards.html#2
func (comp *Client) Get(cardID string) (*response.ResponseCardGet, error) {
	result := &response.ResponseCardGet{}

	param := object.HashMap{
		"card_id": cardID,
	}

	_, err := comp.HttpPostJson("card/get", param, nil, nil, result)

	return result, err

}
