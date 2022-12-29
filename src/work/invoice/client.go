package invoice

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/invoice/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/invoice/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// 查询电子发票
// https://developer.work.weixin.qq.com/document/path/90284
func (comp *Client) GetInvoiceInfo(ctx context.Context, cardID string, encryptCode string) (*response.ResponseInvoiceGetInfo, error) {

	result := &response.ResponseInvoiceGetInfo{}

	data := &object.HashMap{
		"card_id":      cardID,
		"encrypt_code": encryptCode,
	}
	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/card/invoice/reimburse/getinvoiceinfo", data, nil, nil, result)

	return result, err
}

// 批量查询电子发票
// https://developer.work.weixin.qq.com/document/path/90287
func (comp *Client) GetInvoiceInfoBatch(ctx context.Context, invoiceList []*request.RequestCardInvoice) (*response.ResponseInvoiceGetInfoBatch, error) {

	result := &response.ResponseInvoiceGetInfoBatch{}

	data := &object.HashMap{
		"item_list": invoiceList,
	}
	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/card/invoice/reimburse/getinvoiceinfobatch", data, nil, nil, result)

	return result, err
}

// 更新发票状态
// https://developer.work.weixin.qq.com/document/path/90285
func (comp *Client) UpdateInvoiceStatus(ctx context.Context, cardID string, encryptCode string, status string) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	data := &object.HashMap{
		"card_id":          cardID,
		"encrypt_code":     encryptCode,
		"reimburse_status": status,
	}
	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/card/invoice/reimburse/updateinvoicestatus", data, nil, nil, result)

	return result, err
}

// 批量更新发票状态
// https://developer.work.weixin.qq.com/document/path/90286
func (comp *Client) UpdateStatusBatch(ctx context.Context, openid string, status string, invoiceList []*request.RequestCardInvoice) (*response2.ResponseWork, error) {

	result := &response2.ResponseWork{}

	data := &object.HashMap{
		"openid":           openid,
		"reimburse_status": status,
		"invoice_list":     invoiceList,
	}
	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/card/invoice/reimburse/updatestatusbatch", data, nil, nil, result)

	return result, err
}
