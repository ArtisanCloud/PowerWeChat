package merchantService

import (
	"context"
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	payment "github.com/ArtisanCloud/PowerWeChat/v3/src/payment/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/merchantService/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/merchantService/response"
	"net/http"
)

type Client struct {
	*payment.BaseClient
}

func NewClient(app *payment.ApplicationPaymentInterface) (*Client, error) {
	baseClient, err := payment.NewBaseClient(app)
	if err != nil {
		return nil, err
	}
	return &Client{
		baseClient,
	}, nil
}

// 查询投诉单列表
// https://pay.weixin.qq.com/docs/merchant/apis/consumer-complaint/complaints/list-complaints-v2.html
func (comp *Client) Complaints(ctx context.Context, params *request.RequestComplaints) (*response.ResponseComplaints, error) {

	result := &response.ResponseComplaints{}

	data, _ := object.StructToStringMap(params)

	endpoint := "/v3/merchant-service/complaints-v2"
	_, err := comp.Request(ctx, endpoint, data, http.MethodGet, &object.HashMap{}, false, nil, result)

	return result, err
}

// 查询投诉单详情
// https://pay.weixin.qq.com/docs/merchant/apis/consumer-complaint/complaints/query-complaint-v2.html
func (comp *Client) QueryComplaint(ctx context.Context, complaintId string) (*response.ResponseQueryComplaint, error) {

	result := &response.ResponseQueryComplaint{}

	endpoint := fmt.Sprintf("/v3/merchant-service/complaints-v2/%s", complaintId)
	_, err := comp.Request(ctx, endpoint, nil, http.MethodGet, &object.HashMap{}, false, nil, result)

	return result, err
}

// 查询投诉单协商历史
// https://pay.weixin.qq.com/docs/merchant/apis/consumer-complaint/complaints/query-negotiation-history-v2.html
func (comp *Client) QueryNegotiationHistoriesByComplaint(ctx context.Context, complaintId string, limit int8, offset int8) (*response.ResponseQueryNegotiationHistoriesByComplaint, error) {

	result := &response.ResponseQueryNegotiationHistoriesByComplaint{}

	data := &object.StringMap{
		"limit":  fmt.Sprintf("%d", limit),
		"offset": fmt.Sprintf("%d", offset),
	}

	endpoint := fmt.Sprintf("/v3/merchant-service/complaints-v2/%s/negotiation-historys", complaintId)
	_, err := comp.Request(ctx, endpoint, data, http.MethodGet, &object.HashMap{}, false, nil, result)

	return result, err
}

// 创建投诉通知回调
// https://pay.weixin.qq.com/docs/merchant/apis/consumer-complaint/complaint-notifications/create-complaint-notifications.html
func (comp *Client) CreateComplaintNotifications(ctx context.Context, url string) (*response.ResponseCreateComplaintNotifications, error) {

	result := &response.ResponseCreateComplaintNotifications{}

	data := &object.HashMap{
		"url": url,
	}

	endpoint := "/v3/merchant-service/complaint-notifications"
	_, err := comp.Request(ctx, endpoint, nil, http.MethodPost, data, false, nil, result)

	return result, err
}

// 查询投诉通知回调
// https://pay.weixin.qq.com/docs/merchant/apis/consumer-complaint/complaint-notifications/query-complaint-notifications.html
func (comp *Client) QueryComplaintNotifications(ctx context.Context, complaintId string, limit int8, offset int8) (*response.ResponseCreateComplaintNotifications, error) {

	result := &response.ResponseCreateComplaintNotifications{}

	endpoint := "/v3/merchant-service/complaint-notifications"
	_, err := comp.Request(ctx, endpoint, nil, http.MethodGet, &object.HashMap{}, false, nil, result)

	return result, err
}

// 更新投诉通知回调
// https://pay.weixin.qq.com/docs/merchant/apis/consumer-complaint/complaint-notifications/update-complaint-notifications.html
func (comp *Client) UpdateComplaintNotifications(ctx context.Context, url string) (*response.ResponseCreateComplaintNotifications, error) {

	result := &response.ResponseCreateComplaintNotifications{}

	data := &object.HashMap{
		"url": url,
	}

	endpoint := "/v3/merchant-service/complaint-notifications"
	_, err := comp.Request(ctx, endpoint, nil, http.MethodPut, data, false, nil, result)

	return result, err
}

// 删除投诉通知回调
// https://pay.weixin.qq.com/docs/merchant/apis/consumer-complaint/complaint-notifications/delete-complaint-notifications.html
func (comp *Client) DeleteComplaintNotifications(ctx context.Context, url string) (*response2.ResponsePayment, error) {

	result := &response2.ResponsePayment{}

	endpoint := "/v3/merchant-service/complaint-notifications"
	_, err := comp.Request(ctx, endpoint, nil, http.MethodDelete, &object.HashMap{}, false, nil, result)

	return result, err
}

// 回复用户
// https://pay.weixin.qq.com/docs/merchant/apis/consumer-complaint/complaints/response-complaint-v2.html
func (comp *Client) ReplyToUser(ctx context.Context, complaintId string, params *request.RequestReplyToUser) (*response2.ResponsePayment, error) {

	result := &response2.ResponsePayment{}

	data, _ := object.StructToHashMap(params)

	endpoint := fmt.Sprintf("/v3/merchant-service/complaints-v2/%s/response", complaintId)
	_, err := comp.Request(ctx, endpoint, nil, http.MethodPost, data, false, nil, result)

	return result, err
}

// 反馈处理完成
// https://pay.weixin.qq.com/docs/merchant/apis/consumer-complaint/complaints/complete-complaint-v2.html
func (comp *Client) CompleteFeedback(ctx context.Context, complaintId string) (*response2.ResponsePayment, error) {

	result := &response2.ResponsePayment{}

	endpoint := fmt.Sprintf("/v3/merchant-service/complaints-v2/%s/complete", complaintId)
	_, err := comp.Request(ctx, endpoint, nil, http.MethodPost, &object.HashMap{}, false, nil, result)

	return result, err
}

// 更新退款审批结果
// https://pay.weixin.qq.com/docs/merchant/apis/consumer-complaint/complaints/complete-complaint-v2.html
func (comp *Client) UpdateFeedback(ctx context.Context, complaintId string, params *request.RequestUpdateFeedback) (*response2.ResponsePayment, error) {

	result := &response2.ResponsePayment{}

	data, _ := object.StructToHashMap(params)

	endpoint := fmt.Sprintf("/v3/merchant-service/complaints-v2/%s/update-refund-progress", complaintId)
	_, err := comp.Request(ctx, endpoint, nil, http.MethodPost, data, false, nil, result)

	return result, err
}

// 图片上传接口
// https://pay.weixin.qq.com/docs/merchant/apis/consumer-complaint/images/create-images.html
func (comp *Client) UploadImg(ctx context.Context, params *request.RequestMediaUpload) (*response.ResponseMediaUpload, error) {

	result := &response.ResponseMediaUpload{}

	var files *object.HashMap
	if params.File != "" {
		files = &object.HashMap{
			"file": params.File,
		}
	}

	var formData *kernel.UploadForm
	if params.Meta != nil {
		formData = &kernel.UploadForm{
			Contents: []*kernel.UploadContent{
				&kernel.UploadContent{
					Name:  "file",
					Value: params.Meta.Filename,
				},
			},
		}
	}
	options, _ := object.StructToHashMap(params.Meta)

	_, err := comp.BaseClient.HttpUploadJson(ctx, "/v3/merchant-service/images/upload", files, formData, options, nil, nil, &result)
	return result, err
}
