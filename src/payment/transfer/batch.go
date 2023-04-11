package transfer

import (
	"context"
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	payment "github.com/ArtisanCloud/PowerWeChat/v3/src/payment/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/transfer/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/transfer/response"
	"net/http"
)

type BatchClient struct {
	*payment.BaseClient
	Version string
}

func NewBatchClient(app *payment.ApplicationPaymentInterface) (*BatchClient, error) {
	baseClient, err := payment.NewBaseClient(app)
	if err != nil {
		return nil, err
	}
	return &BatchClient{
		BaseClient: baseClient,
		Version:    "v3",
	}, nil
}

// 发起商家转账API
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter4_3_1.shtml
func (comp *BatchClient) Batch(ctx context.Context, param *request.RequestTransferBatch) (*response.ResponseTrasferBatch, error) {

	result := &response.ResponseTrasferBatch{}

	options, err := object.StructToHashMap(param)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("/%s/transfer/batches", comp.Version)
	_, err = comp.Request(ctx, comp.Wrap(endpoint), nil, http.MethodPost, options, false, nil, result)
	return result, err
}

// 微信批次单号查询批次单AP
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter4_3_2.shtml
func (comp *BatchClient) QueryBatch(ctx context.Context, batchID string, needQueryDetail bool, offset int, limit int, detailStatus string) (*response.ResponseTrasferQueryBatch, error) {

	result := &response.ResponseTrasferQueryBatch{}

	query := &object.StringMap{
		"need_query_detail": fmt.Sprintf("%t", needQueryDetail),
		"offset":            fmt.Sprintf("%d", offset),
		"limit":             fmt.Sprintf("%d", limit),
		"detail_status":     detailStatus,
	}

	endpoint := fmt.Sprintf("/%s/transfer/batches/batch-id/%s", comp.Version, batchID)
	_, err := comp.Request(ctx, comp.Wrap(endpoint), query, http.MethodPost, nil, false, nil, result)
	return result, err
}

// 微信明细单号查询明细单API
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter4_3_3.shtml
func (comp *BatchClient) QueryBatchDetail(ctx context.Context, batchID string, detailID string) (*response.TransferBatch, error) {

	result := &response.TransferBatch{}

	endpoint := fmt.Sprintf("/%s/transfer/batches/batch-id/%s/details/detail-id/%s", comp.Version, batchID, detailID)
	_, err := comp.Request(ctx, comp.Wrap(endpoint), nil, http.MethodPost, nil, false, nil, result)
	return result, err
}

// 商家批次单号查询批次单API
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter4_3_5.shtml
func (comp *BatchClient) QueryOutBatchNO(ctx context.Context, outBatchNO string, needQueryDetail bool, offset int, limit int, detailStatus string) (*response.ResponseTrasferQueryOutBatchNO, error) {

	result := &response.ResponseTrasferQueryOutBatchNO{}

	query := &object.StringMap{
		"need_query_detail": fmt.Sprintf("%t", needQueryDetail),
		"offset":            fmt.Sprintf("%d", offset),
		"limit":             fmt.Sprintf("%d", limit),
		"detail_status":     detailStatus,
	}

	endpoint := fmt.Sprintf("/%s/transfer/batches/out_batch_no/%s", comp.Version, outBatchNO)
	_, err := comp.Request(ctx, comp.Wrap(endpoint), query, http.MethodPost, nil, false, nil, result)
	return result, err
}

// 商家明细单号查询明细单API
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter4_3_6.shtml
func (comp *BatchClient) QueryOutBatchNoDetail(ctx context.Context, outBatchNO string, outDetailNO string) (*response.ResponseOutBatchNODetail, error) {

	result := &response.ResponseOutBatchNODetail{}

	endpoint := fmt.Sprintf("/%s/transfer/batches/out-batch-no/%s/details/out-detail-no/%s", comp.Version, outBatchNO, outDetailNO)
	_, err := comp.Request(ctx, comp.Wrap(endpoint), nil, http.MethodPost, nil, false, nil, result)
	return result, err
}

// 转账电子回单申请受理API
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter4_3_7.shtml
func (comp *BatchClient) BillReceipt(ctx context.Context, outBatchNO string) (*response.ResponseTrasferBillReceipt, error) {

	result := &response.ResponseTrasferBillReceipt{}

	options := &object.HashMap{
		"out_batch_no": outBatchNO,
	}

	endpoint := fmt.Sprintf("/%s/transfer/bill-receipt", comp.Version)
	_, err := comp.Request(ctx, comp.Wrap(endpoint), nil, http.MethodPost, options, false, nil, result)
	return result, err
}

// 查询转账电子回单API
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter4_3_8.shtml
func (comp *BatchClient) QueryBillReceipt(ctx context.Context, outBatchNO string) (*response.ResponseTrasferBillReceipt, error) {

	result := &response.ResponseTrasferBillReceipt{}

	endpoint := fmt.Sprintf("/%s/transfer/bill-receipt/%s", comp.Version, outBatchNO)
	_, err := comp.Request(ctx, comp.Wrap(endpoint), nil, http.MethodPost, nil, false, nil, result)
	return result, err
}

// 转账明细电子回单受理API
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter4_3_9.shtml
func (comp *BatchClient) ElectronicReceipts(ctx context.Context, acceptType string, outBatchNO string, outDetailNO string) (*response.ResponseTrasferElectronicReceipts, error) {

	result := &response.ResponseTrasferElectronicReceipts{}

	options := &object.HashMap{
		"accept_type":   acceptType,
		"out_batch_no":  outBatchNO,
		"out_detail_no": outDetailNO,
	}

	endpoint := fmt.Sprintf("/%s/transfer-detail/electronic-receipts", comp.Version)
	_, err := comp.Request(ctx, comp.Wrap(endpoint), nil, http.MethodPost, options, false, nil, result)
	return result, err
}

// 查询转账明细电子回单受理结果API
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter4_3_10.shtml
func (comp *BatchClient) QueryElectronicReceipts(ctx context.Context, acceptType string, outBatchNO string, outDetailNO string) (*response.ResponseTrasferElectronicReceipts, error) {

	result := &response.ResponseTrasferElectronicReceipts{}

	query := &object.StringMap{
		"accept_type":   acceptType,
		"out_batch_no":  outBatchNO,
		"out_detail_no": outDetailNO,
	}

	endpoint := fmt.Sprintf("/%s/transfer-detail/electronic-receipts", comp.Version)
	_, err := comp.Request(ctx, comp.Wrap(endpoint), query, http.MethodPost, nil, false, nil, result)
	return result, err
}

// 下载电子回单API
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter4_3_11.shtml
func (comp *Client) DownloadSignFile(ctx context.Context, requestDownload *power.RequestDownload, filePath string) (int64, error) {
	return comp.StreamDownload(ctx, requestDownload, filePath)
}
