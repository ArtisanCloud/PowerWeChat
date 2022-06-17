package material

import (
	"github.com/ArtisanCloud/PowerLibs/v2/http/contract"
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
	request2 "github.com/ArtisanCloud/PowerWeChat/v2/src/officialAccount/material/request"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/officialAccount/material/response"
	response4 "github.com/ArtisanCloud/PowerWeChat/v2/src/work/media/response"
	"os"
	"path/filepath"
)

type Client struct {
	*kernel.BaseClient

	AllowTypes []string
}

func NewClient(app kernel.ApplicationInterface) (*Client, error) {
	baseClient, err := kernel.NewBaseClient(&app, nil)
	if err != nil {
		return nil, err
	}
	return &Client{
		BaseClient: baseClient,
		AllowTypes: []string{"image", "voice", "video", "thumb", "news_image"},
	}, nil
}

// 上传永久图片素材
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Adding_Permanent_Assets.html
func (comp *Client) UploadImage(path string) (*response.ResponseMaterialAddMaterial, error) {
	result := &response.ResponseMaterialAddMaterial{}
	_, err := comp.Upload("image", path, &object.StringMap{}, result)
	return result, err
}

// 上传永久语音素材
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Adding_Permanent_Assets.html
func (comp *Client) UploadVoice(path string) (*response.ResponseMaterialAddMaterial, error) {
	result := &response.ResponseMaterialAddMaterial{}
	_, err := comp.Upload("voice", path, &object.StringMap{}, result)
	return result, err
}

// 上传永久缩略图素材
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Adding_Permanent_Assets.html
func (comp *Client) UploadThumb(path string) (*response.ResponseMaterialAddMaterial, error) {
	result := &response.ResponseMaterialAddMaterial{}
	_, err := comp.Upload("thumb", path, &object.StringMap{}, result)
	return result, err
}

// 上传永久视频素材
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Adding_Permanent_Assets.html
func (comp *Client) UploadVideo(path string, title string, description string) (*response.ResponseMaterialAddMaterial, error) {

	result := &response.ResponseMaterialAddMaterial{}

	jsonDescription, err := object.JsonEncode(&object.HashMap{
		"title":        title,
		"introduction": description,
	})
	if err != nil {
		return nil, err
	}

	params := &object.StringMap{
		"Description": jsonDescription,
	}

	_, err = comp.Upload("video", path, params, result)
	return result, err
}

// 新增永久素材
// https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html
func (comp *Client) UploadArticle(articles request2.RequestAddArticles) (*response.ResponseMaterialAddNews, error) {

	result := &response.ResponseMaterialAddNews{}

	//params, err := object.StructToHashMapWithTag(articles, "json")
	params, err := object.StructToHashMap(articles)
	if err != nil {
		return nil, err
	}

	_, err = comp.HttpPostJson("cgi-bin/material/add_news", params, nil, nil, result)
	return result, err
}

// 上传永久视频素材
// https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html
func (comp *Client) UpdateArticle(mediaID string, articles request2.RequestAddArticles, index int) (response.ResponseMaterialAddNews, error) {
	result := response.ResponseMaterialAddNews{}

	params := &object.HashMap{
		"media_id": mediaID,
		"index":    index,
		"articles": articles,
	}

	_, err := comp.HttpPostJson("cgi-bin/material/update_news", params, nil, nil, result)
	return result, err
}

// 上传图文消息内的图片获取URL
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Adding_Permanent_Assets.html
func (comp *Client) UploadArticleImage(path string) (*response.ResponseMaterialAddMaterial, error) {
	result := &response.ResponseMaterialAddMaterial{}
	_, err := comp.Upload("news_image", path, &object.StringMap{}, result)
	return result, err
}

// 获取永久素材图片
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Getting_Permanent_Assets.html
func (comp *Client) Get(mediaID string) (contract.ResponseInterface, error) {

	result := ""
	header := &response4.ResponseHeaderMedia{}
	response, err := comp.RequestRaw("cgi-bin/material/get_material", "POST", &object.HashMap{
		"form_params": &object.HashMap{
			"media_id": mediaID,
		},
	}, header, &result)

	return response.(contract.ResponseInterface), err
}

// 获取永久视频消息素材
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Getting_Permanent_Assets.html
func (comp *Client) GetVideo(mediaID string) (*response.ResponseMaterialGetVideo, error) {

	result := &response.ResponseMaterialGetVideo{}

	options := &object.HashMap{
		"media_id": mediaID,
	}

	_, err := comp.HttpPostJson("cgi-bin/material/get_material", options, nil, nil, result)

	return result, err
}

// 获取永久图文素材
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Getting_Permanent_Assets.html
func (comp *Client) GetNews(mediaID string) (*response.ResponseMaterialGetNews, error) {

	result := &response.ResponseMaterialGetNews{}

	options := &object.HashMap{
		"media_id": mediaID,
	}

	_, err := comp.HttpPostJson("cgi-bin/material/get_material", options, nil, nil, result)

	return result, err
}

// 删除永久素材
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Deleting_Permanent_Assets.html
func (comp *Client) Delete(mediaID int) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	options := &object.HashMap{
		"media_id": mediaID,
	}

	_, err := comp.HttpPostJson("cgi-bin/material/del_material", options, nil, nil, result)

	return result, err
}

// 获取素材列表
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Get_materials_list.html
func (comp *Client) List(options *request2.RequestMaterialBatchGetMaterial) (*response.ResponseMaterialBatchGetMaterial, error) {

	result := &response.ResponseMaterialBatchGetMaterial{}

	_, err := comp.HttpPostJson("cgi-bin/material/batchget_material", options, nil, nil, result)

	return result, err
}

// 获取素材总数
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Get_the_total_of_all_materials.html
func (comp *Client) Stats() (*response.ResponseMaterialGetMaterialCount, error) {

	result := &response.ResponseMaterialGetMaterialCount{}

	_, err := comp.HttpPostJson("cgi-bin/material/get_materialcount", nil, nil, nil, result)

	return result, err

}

func (comp *Client) Upload(Type string, path string, query *object.StringMap, result interface{}) (interface{}, error) {

	_, err := os.Stat(path)
	if (err != nil && os.IsExist(err)) && (err != nil && os.IsPermission(err)) {
		return "", err
	}
	file := &object.HashMap{
		"media": path,
	}

	(*query)["type"] = Type

	form := &object.HashMap{
		"filename": filepath.Base(path),
	}

	return comp.HttpUpload(comp.getApiByType(Type), file, form, query, nil, result)
}

func (comp *Client) getApiByType(Type string) string {

	switch Type {
	case "news_image":
		return "cgi-bin/media/uploadimg"
	default:
		return "cgi-bin/material/add_material"
	}

}
