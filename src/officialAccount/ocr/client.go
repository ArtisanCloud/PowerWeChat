package ocr

import (
	"context"
	"errors"
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/ocr/response"
)

type Client struct {
	BaseClient *kernel.BaseClient

	AllowTypes []string
}

func NewClient(app kernel.ApplicationInterface) (*Client, error) {
	baseClient, err := kernel.NewBaseClient(&app, nil)
	if err != nil {
		return nil, err
	}
	return &Client{
		BaseClient: baseClient,
		AllowTypes: []string{"photo", "scan"},
	}, nil
}

// 身份证 OCR 识别接口
// https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html
func (comp *Client) IDCard(ctx context.Context, path string, IDType string) (*response.ResponseOCRIDCard, error) {

	if !object.ContainsString(comp.AllowTypes, IDType) {
		return nil, errors.New(fmt.Sprintf("Unsupported media type: '%s'", IDType))
	}

	result := &response.ResponseOCRIDCard{}

	params := &object.HashMap{
		"type":    IDType,
		"img_url": path,
	}

	_, err := comp.BaseClient.HttpPost(ctx, "cv/ocr/idcard", params, nil, result)

	return result, err
}

// 银行卡 OCR 识别接口
// https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html
func (comp *Client) BankCard(ctx context.Context, path string) (*response.ResponseOCRBankCard, error) {

	result := &response.ResponseOCRBankCard{}

	params := &object.HashMap{
		"img_url": path,
	}

	_, err := comp.BaseClient.HttpPost(ctx, "cv/ocr/bankcard", params, nil, result)

	return result, err
}

// 驾驶证 OCR 识别接口
// https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html
func (comp *Client) VehicleLicense(ctx context.Context, path string) (*response.ResponseOCRVehicleLicense, error) {

	result := &response.ResponseOCRVehicleLicense{}

	params := &object.HashMap{
		"img_url": path,
	}

	_, err := comp.BaseClient.HttpPost(ctx, "cv/ocr/drivinglicense", params, nil, result)

	return result, err
}

// 驾驶证 OCR 识别接口
// https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html
func (comp *Client) Driving(ctx context.Context, path string) (*response.ResponseOCRDriving, error) {

	result := &response.ResponseOCRDriving{}

	params := &object.HashMap{
		"img_url": path,
	}

	_, err := comp.BaseClient.HttpPost(ctx, "cv/ocr/driving", params, nil, result)

	return result, err
}

// 营业执照 OCR 识别接口
// https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html
func (comp *Client) BizLicense(ctx context.Context, path string) (*response.ResponseOCRBizLicense, error) {

	result := &response.ResponseOCRBizLicense{}

	params := &object.HashMap{
		"img_url": path,
	}

	_, err := comp.BaseClient.HttpPost(ctx, "cv/ocr/bizlicense", params, nil, result)

	return result, err
}

// 通用印刷体 OCR 识别接口
// https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html
func (comp *Client) Common(ctx context.Context, path string) (*response.ResponseOCRCommon, error) {

	result := &response.ResponseOCRCommon{}

	params := &object.HashMap{
		"img_url": path,
	}

	_, err := comp.BaseClient.HttpPost(ctx, "cv/ocr/comm", params, nil, result)

	return result, err
}

// 车牌识别接口
// https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/OCR.html
func (comp *Client) PlateNumber(ctx context.Context, path string) (*response.ResponseOCRPlateNumber, error) {

	result := &response.ResponseOCRPlateNumber{}

	params := &object.HashMap{
		"img_url": path,
	}

	_, err := comp.BaseClient.HttpPost(ctx, "cv/ocr/platenum", params, nil, result)

	return result, err
}
