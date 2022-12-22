package wifi

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/wifi/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/wifi/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) (*Client, error) {
	baseClient, err := kernel.NewBaseClient(&app, nil)
	if err != nil {
		return nil, err
	}
	return &Client{
		baseClient,
	}, nil
}

// 查询一定时间范围内的 WiFi 连接总人数、微信方式连 Wi-Fi 人数、商家主页访问人数、连网后消息发送人数、新增公众号关注人数和累计公众号关注人数。
// https://developers.weixin.qq.com/doc/offiaccount/WiFi_via_WeChat/Wi-Fi_data_statistics.html
func (comp *Client) Summary(beginDate string, endDate string, shopID int) (*response.ResponseWifiSummary, error) {

	result := &response.ResponseWifiSummary{}

	params := &object.HashMap{
		"begin_date": beginDate,
		"end_date":   endDate,
		"shop_id":    shopID,
	}

	_, err := comp.HttpPostJson("bizwifi/statistics/list", params, nil, nil, result)

	return result, err

}

// 添加设备后，通过此接口可以获取物料。
// https://developers.weixin.qq.com/doc/offiaccount/WiFi_via_WeChat/Connecting/Get_materials_QR_Code.html
func (comp *Client) GetQrCodeUrl(bShopID int, ssid string, ImageID int) (*response.ResponseWifiQrCodeURL, error) {

	result := &response.ResponseWifiQrCodeURL{}

	params := &object.HashMap{
		"shop_id": bShopID,
		"ssid":    ssid,
		"img_id":  ImageID,
	}

	_, err := comp.HttpPostJson("bizwifi/qrcode/get", params, nil, nil, result)

	return result, err

}

// 当顾客使用微信连 Wi-Fi 方式连网成功时，点击页面右上角“完成”按钮，即可进入已设置的连网完成页。
// https://developers.weixin.qq.com/doc/offiaccount/WiFi_via_WeChat/Home_Page_management/Set_up_the_Wi-Fi_Connection_Completion_page.html
func (comp *Client) SetFinishPage(bShopID int, finishPageURL string) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	params := &object.HashMap{
		"shop_id":        bShopID,
		"finishpage_url": finishPageURL,
	}

	_, err := comp.HttpPostJson("bizwifi/finishpage/set", params, nil, nil, result)

	return result, err

}

// 设置商户主页后，点击微信聊首页欢迎语，即可进入设置的商户主页。
// https://developers.weixin.qq.com/doc/offiaccount/WiFi_via_WeChat/Home_Page_management/Set_up_a_merchant_home_page.html
func (comp *Client) SetHomePage(data *request.RequestWifiSetHomePage) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	_, err := comp.HttpPostJson("bizwifi/homepage/set", data, nil, nil, result)

	return result, err

}
