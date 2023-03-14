package qrCode

import (
	"context"
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/basicService/qrCode/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/basicService/qrCode/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"math"
	"net/url"
)

const DAY int = 86400
const SCENE_MAX_VALUE int = 100000
const SCENE_QR_CARD string = "QR_CARD"
const SCENE_QR_TEMPORARY string = "QR_SCENE"
const SCENE_QR_TEMPORARY_STR string = "QR_STR_SCENE"
const SCENE_QR_FOREVER string = "QR_LIMIT_SCENE"
const SCENE_QR_FOREVER_STR string = "QR_LIMIT_STR_SCENE"

type Client struct {
	BaseClient *kernel.BaseClient
}

func NewClient(app *kernel.ApplicationInterface) (*Client, error) {
	baseClient, err := kernel.NewBaseClient(app, nil)
	if err != nil {
		return nil, err
	}
	client := &Client{
		BaseClient: baseClient,
	}

	client.BaseClient.BaseURI = "https://api.weixin.qq.com/"
	return client, nil
}

// 生成永久二维码请求
// https://developers.weixin.qq.com/doc/offiaccount/Account_Management/Generating_a_Parametric_QR_Code.html
func (comp *Client) Forever(ctx context.Context, sceneValue interface{}) (*response.ResponseQRCodeCreate, error) {

	data := &request.RequestQRCodeCreate{
		ActionInfo: &request.ActionInfo{},
	}

	switch sceneValue.(type) {
	case int:
		value := sceneValue.(int)
		if value > 0 && value < SCENE_MAX_VALUE {
			data.ActionName = SCENE_QR_FOREVER
			data.ActionInfo.Scene = &object.HashMap{
				"scene_id": sceneValue,
			}
		}
		break
	case string:
		data.ActionName = SCENE_QR_FOREVER_STR
		data.ActionInfo.Scene = &object.HashMap{
			"scene_str": sceneValue,
		}
		break
	default:
	}

	return comp.Create(ctx, data, false, 9999000)

}

// 生成临时二维码请求
// https://developers.weixin.qq.com/doc/offiaccount/Account_Management/Generating_a_Parametric_QR_Code.html
func (comp *Client) Temporary(ctx context.Context, sceneValue interface{}, expireSeconds int) (*response.ResponseQRCodeCreate, error) {

	data := &request.RequestQRCodeCreate{
		ActionInfo: &request.ActionInfo{},
	}

	switch sceneValue.(type) {
	case int:
		value := sceneValue.(int)
		if value > 0 && value < SCENE_MAX_VALUE {
			data.ActionName = SCENE_QR_TEMPORARY
			data.ActionInfo.Scene = &object.HashMap{
				"scene_id": sceneValue,
			}
		}
		break
	case string:
		data.ActionName = SCENE_QR_TEMPORARY_STR
		data.ActionInfo.Scene = &object.HashMap{
			"scene_str": sceneValue,
		}
		break
	default:
	}

	return comp.Create(ctx, data, true, expireSeconds)

}

// 生成二维码ticket
// https://developers.weixin.qq.com/doc/offiaccount/Account_Management/Generating_a_Parametric_QR_Code.html
func (comp *Client) URL(ticket string) string {
	return fmt.Sprintf("https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=%s", url.QueryEscape(ticket))
}

// 生成带参数的二维码
// https://developers.weixin.qq.com/doc/offiaccount/Account_Management/Generating_a_Parametric_QR_Code.html
func (comp *Client) Create(ctx context.Context, data *request.RequestQRCodeCreate, temporary bool, expireSecond int) (*response.ResponseQRCodeCreate, error) {

	result := &response.ResponseQRCodeCreate{}

	//params, err := object.StructToHashMapWithTag(data, "json")
	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	if temporary {
		(*params)["expire_seconds"] = int(math.Min(float64(expireSecond), float64(30*DAY)))
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, "cgi-bin/qrcode/create", params, nil, nil, result)

	return result, err

}
