package shakeAround

import (
	"errors"
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/src/officialAccount/shakeAround/response"
	"os"
	"strings"
)

type MaterialClient struct {
	*kernel.BaseClient
}

func NewMaterialClient(app kernel.ApplicationInterface) *MaterialClient {
	return &MaterialClient{
		kernel.NewBaseClient(&app, nil),
	}
}

// 上传在摇一摇功能中需使用到的图片素材
// https://developers.weixin.qq.com/doc/offiaccount/Shake_Nearby/Uploading_Image_Assets.html
func (comp *MaterialClient) UploadImage(path string, Type string) (*response.ResponseMaterialUpload, error) {

	result := &response.ResponseMaterialUpload{}

	_, err := os.Stat(path)
	if (err != nil && os.IsExist(err)) && (err != nil && os.IsPermission(err)) {
		return nil, errors.New(fmt.Sprintf("File does not exist, or the file is unreadable: \"%s\"", path))
	}

	file := &object.HashMap{
		"media": path,
	}

	_, err = comp.HttpUpload("shakearound/material/add", file, nil, &object.StringMap{"type": strings.ToLower(Type)}, nil, result)
	return result, err
}
