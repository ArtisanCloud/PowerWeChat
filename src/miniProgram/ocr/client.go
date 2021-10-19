package ocr

import (
	"errors"
	"github.com/ArtisanCloud/PowerLibs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/miniProgram/ocr/response"
)

type Client struct {
	*kernel.BaseClient
}

// 本接口提供基于小程序的银行卡 OCR 识别
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.bankcard.html
func (comp *Client) Bankcard(imgURL string, img *power.HashMap) (*response.ResponseOCRBankcard, error) {

	result := &response.ResponseOCRBankcard{}

	_, err := comp.UploadImage("cv/ocr/bankcard", imgURL, img, result)

	return result, err
}

// 本接口提供基于小程序的营业执照 OCR 识别
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.businessLicense.html
func (comp *Client) BusinessLicense(imgURL string, img *power.HashMap) (*response.ResponseOCRBusinessLicense, error) {

	result := &response.ResponseOCRBusinessLicense{}

	_, err := comp.UploadImage("cv/ocr/bizlicense", imgURL, img, result)

	return result, err
}

// 本接口提供基于小程序的驾驶证 OCR 识别
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.driverLicense.html
func (comp *Client) DriverLicense(imgURL string, img *power.HashMap) (*response.ResponseOCRDrivingLicense, error) {

	result := &response.ResponseOCRDrivingLicense{}

	_, err := comp.UploadImage("cv/ocr/drivinglicense", imgURL, img, result)

	return result, err

}

// 本接口提供基于小程序的身份证 OCR 识别
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.idcard.html
func (comp *Client) IDCard(imgURL string, img *power.HashMap) (*response.ResponseOCRIDCard, error) {

	result := &response.ResponseOCRIDCard{}

	_, err := comp.UploadImage("cv/ocr/idcard", imgURL, img, result)

	return result, err

}

// 本接口提供基于小程序的通用印刷体 OCR 识别
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.printedText.html
func (comp *Client) PrintedText(imgURL string, img *power.HashMap) (*response.ResponseOCRPrintedText, error) {

	result := &response.ResponseOCRPrintedText{}

	_, err := comp.UploadImage("cv/ocr/comm", imgURL, img, result)

	return result, err

}

// 本接口提供基于小程序的行驶证 OCR 识别
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/ocr/ocr.vehicleLicense.html
func (comp *Client) VehicleLicense(mode string, imgURL string, img *power.HashMap) (*response.ResponseOCRVehicleLicense, error) {

	result := &response.ResponseOCRVehicleLicense{}

	data := &object.HashMap{
		"type": mode,
	}

	if imgURL != "" {
		(*data)["img_url"] = imgURL
	} else if img != nil {
		(*data)["img"] = img
	} else {
		return nil, errors.New("Please send image url or image form data ")
	}

	_, err := comp.UploadImage("cv/ocr/driving", imgURL, img, result)

	return result, err

}

func (comp *Client) UploadImage(entryPoint string, imgURL string, img *power.HashMap, result interface{}) (interface{}, error) {

	params := &object.StringMap{}
	var formData *object.HashMap

	if imgURL != "" {
		params = &object.StringMap{
			"img_url": imgURL,
		}
	} else if img != nil {
		formData = &object.HashMap{
			"name":  (*img)["name"],
			"value": (*img)["value"],
		}
	} else {
		return nil, errors.New("Please send image url or image form data ")
	}

	rs, err := comp.HttpUpload(entryPoint, nil, formData, params, nil, result)

	return rs, err
}
