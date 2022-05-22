package material

import (
	"github.com/ArtisanCloud/PowerLibs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	response2 "github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
	request2 "github.com/ArtisanCloud/PowerWeChat/src/officialAccount/material/request"
	response3 "github.com/ArtisanCloud/PowerWeChat/src/officialAccount/material/response"
	"os"
)

type Client struct {
	*kernel.BaseClient

	AllowTypes []string
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		BaseClient: kernel.NewBaseClient(&app, nil),
		AllowTypes: []string{"image", "voice", "video", "thumb", "news_image"},
	}
}

// 上传永久图片素材
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Adding_Permanent_Assets.html
func (comp *Client) UploadImage(path string) (interface{}, error) {
	return comp.Upload("image", path, nil)
}

// 上传永久语音素材
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Adding_Permanent_Assets.html
func (comp *Client) UploadVoice(path string) (interface{}, error) {
	return comp.Upload("voice", path, nil)
}

// 上传永久缩略图素材
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Adding_Permanent_Assets.html
func (comp *Client) UploadThumb(path string) (interface{}, error) {
	return comp.Upload("thumb", path, nil)
}

// 上传永久视频素材
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Adding_Permanent_Assets.html
func (comp *Client) UploadVideo(path string, title string, description string) (interface{}, error) {

	jsonDescription, err := object.JsonEncode(&object.HashMap{
		"title":        title,
		"introduction": description,
	})
	if err != nil {
		return nil, err
	}

	params := &power.HashMap{
		"Description": jsonDescription,
	}

	return comp.Upload("video", path, params)
}

// 新增永久素材
// https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html
func (comp *Client) UploadArticle(articles request2.RequestAddArticles) (*response3.ResponseMaterialAddNews, error) {

	result := &response3.ResponseMaterialAddNews{}

	params, err := object.StructToHashMapWithTag(articles, "json")
	if err != nil {
		return nil, err
	}

	_, err = comp.HttpPostJson("cgi-bin/material/add_news", params, nil, nil, result)
	return result, err
}

// 上传永久视频素材
// https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html
func (comp *Client) UpdateArticle(mediaID string, articles request2.RequestAddArticles, index int) (response3.ResponseMaterialAddNews, error) {
	result := response3.ResponseMaterialAddNews{}

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
func (comp *Client) UploadArticleImage(path string) (interface{}, error) {
	return comp.Upload("news_image", path, nil)
}

// 获取永久素材
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Getting_Permanent_Assets.html
func (comp *Client) Get(mediaID string) (*response3.ResponseMaterialGetMaterial, error) {

	result := &response3.ResponseMaterialGetMaterial{}

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
func (comp *Client) List(options *request2.RequestMaterialBatchGetMaterial) (*response3.ResponseMaterialBatchGetMaterial, error) {

	result := &response3.ResponseMaterialBatchGetMaterial{}

	_, err := comp.HttpPostJson("cgi-bin/material/batchget_material", options, nil, nil, result)

	return result, err
}

// 获取素材总数
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Get_the_total_of_all_materials.html
func (comp *Client) Stats() (*response3.ResponseMaterialGetMaterialCount, error) {

	result := &response3.ResponseMaterialGetMaterialCount{}

	_, err := comp.HttpPostJson("cgi-bin/material/get_materialcount", nil, nil, nil, result)

	return result, err

}

func (comp *Client) Upload(Type string, path string, form *power.HashMap) (interface{}, error) {

	_, err := os.Stat(path)
	if (err != nil && os.IsExist(err)) && (err != nil && os.IsPermission(err)) {
		return "", err
	}
	file := &object.HashMap{
		"file": path,
	}

	(*form)["type"] = Type

	objForm, err := power.PowerHashMapToObjectHashMap(form)

	return comp.HttpUpload(comp.getApiByType(Type), file, objForm, nil, nil, nil)
}

func (comp *Client) getApiByType(Type string) string {

	switch Type {
	case "news_image":
		return "cgi-bin/media/uploadimg"
	default:
		return "cgi-bin/material/add_material"
	}

}
