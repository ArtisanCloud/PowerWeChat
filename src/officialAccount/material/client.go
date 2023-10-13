package material

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	request2 "github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/material/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/material/response"
	response4 "github.com/ArtisanCloud/PowerWeChat/v3/src/work/media/response"
	"net/http"
	"os"
	"path/filepath"
	"encoding/json"
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

type Client struct {
	*kernel.BaseClient

	AllowTypes []string
}

// 二进制上传素材入参
type UploadMediaRequest struct {
	Type         string `json:"type"`
	Media        []byte `json:"media"`
	FileName     string `json:"fileName"`
	Title        string `json:"title"`
	Introduction string `json:"introduction"`
}

// 二进制上传素材回参
type UploadMediaResponse struct {
	Errcode int         `json:"errcode"`
	MediaId interface{} `json:"media_id"`
	Url     interface{} `json:"url"`
}


// UploadNewsMedia 二进制上传永久图文素材
func (comp *Client) UploadNewsMedia(requests UploadMediaRequest) (data UploadMediaResponse, err error) {
	token, err := comp.Token.GetToken(true)
	if err != nil {
		return nil
	}
	url := "https://api.weixin.qq.com/cgi-bin/media/uploadimg?access_token=" + token
	client := &http.Client{}
	method := "POST"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, err := writer.CreateFormFile("media", requests.FileName)
	file.Write(requests.Media)
	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = json.Unmarshal(body, &data)
	return
}

// UploadOtherMedia 二进制上传其它类型永久素材
func (comp *Client) UploadOtherMedia(requests UploadMediaRequest) (data UploadMediaResponse, err error) {
	token, err := comp.Token.GetToken(true)
	if err != nil {
		return nil
	}
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/material/add_material?access_token=%s&type=%s",token, requests.Type)
	client := &http.Client{}
	method := "POST"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	file, err := writer.CreateFormFile("media", requests.FileName)
	file.Write(requests.Media)
	err = writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	if requests.Type == "video" {
		writer.WriteField("description", fmt.Sprintf("{\"title\":\"%s\",\"introduction\":\"%s\"}", requests.Title, requests.Introduction))
	}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = json.Unmarshal(body, &data)
	return
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
func (comp *Client) UploadImage(ctx context.Context, path string) (*response.ResponseMaterialAddMaterial, error) {
	result := &response.ResponseMaterialAddMaterial{}
	_, err := comp.Upload(ctx, "image", path, &object.StringMap{}, result)
	return result, err
}

// 上传永久语音素材
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Adding_Permanent_Assets.html
func (comp *Client) UploadVoice(ctx context.Context, path string) (*response.ResponseMaterialAddMaterial, error) {
	result := &response.ResponseMaterialAddMaterial{}
	_, err := comp.Upload(ctx, "voice", path, &object.StringMap{}, result)
	return result, err
}

// 上传永久缩略图素材
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Adding_Permanent_Assets.html
func (comp *Client) UploadThumb(ctx context.Context, path string) (*response.ResponseMaterialAddMaterial, error) {
	result := &response.ResponseMaterialAddMaterial{}
	_, err := comp.Upload(ctx, "thumb", path, &object.StringMap{}, result)
	return result, err
}

// 上传永久视频素材
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Adding_Permanent_Assets.html
func (comp *Client) UploadVideo(ctx context.Context, path string, title string, description string) (*response.ResponseMaterialAddMaterial, error) {

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

	_, err = comp.Upload(ctx, "video", path, params, result)
	return result, err
}

// 新增永久素材
// https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html
func (comp *Client) UploadArticle(ctx context.Context, articles request2.RequestAddArticles) (*response.ResponseMaterialAddNews, error) {

	result := &response.ResponseMaterialAddNews{}

	//params, err := object.StructToHashMapWithTag(articles, "json")
	params, err := object.StructToHashMap(articles)
	if err != nil {
		return nil, err
	}

	_, err = comp.BaseClient.HttpPostJson(ctx, "cgi-bin/material/add_news", params, nil, nil, result)
	return result, err
}

// 上传永久视频素材
// https://developers.weixin.qq.com/doc/offiaccount/Comments_management/Image_Comments_Management_Interface.html
func (comp *Client) UpdateArticle(ctx context.Context, mediaID string, articles request2.RequestAddArticles, index int) (response.ResponseMaterialAddNews, error) {
	result := response.ResponseMaterialAddNews{}

	params := &object.HashMap{
		"media_id": mediaID,
		"index":    index,
		"articles": articles,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/material/update_news", params, nil, nil, result)
	return result, err
}

// 上传图文消息内的图片获取URL
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Adding_Permanent_Assets.html
func (comp *Client) UploadArticleImage(ctx context.Context, path string) (*response.ResponseMaterialAddMaterial, error) {
	result := &response.ResponseMaterialAddMaterial{}
	_, err := comp.Upload(ctx, "news_image", path, &object.StringMap{}, result)
	return result, err
}

// 获取永久素材图片
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Getting_Permanent_Assets.html
func (comp *Client) Get(ctx context.Context, mediaID string) (*http.Response, error) {

	header := &response4.ResponseHeaderMedia{}
	response, err := comp.BaseClient.RequestRaw(ctx, "cgi-bin/material/get_material", http.MethodPost, &object.HashMap{
		"form_params": &object.HashMap{
			"media_id": mediaID,
		},
	}, header, nil)

	return response, err
}

// 获取永久视频消息素材
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Getting_Permanent_Assets.html
func (comp *Client) GetVideo(ctx context.Context, mediaID string) (*response.ResponseMaterialGetVideo, error) {

	result := &response.ResponseMaterialGetVideo{}

	options := &object.HashMap{
		"media_id": mediaID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/material/get_material", options, nil, nil, result)

	return result, err
}

// 获取永久图文素材
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Getting_Permanent_Assets.html
func (comp *Client) GetNews(ctx context.Context, mediaID string) (*response.ResponseMaterialGetNews, error) {

	result := &response.ResponseMaterialGetNews{}

	options := &object.HashMap{
		"media_id": mediaID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/material/get_material", options, nil, nil, result)

	return result, err
}

// 删除永久素材
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Deleting_Permanent_Assets.html
func (comp *Client) Delete(ctx context.Context, mediaID string) (*response2.ResponseOfficialAccount, error) {

	result := &response2.ResponseOfficialAccount{}

	options := &object.HashMap{
		"media_id": mediaID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/material/del_material", options, nil, nil, result)

	return result, err
}

// 获取素材列表
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Get_materials_list.html
func (comp *Client) List(ctx context.Context, options *request2.RequestMaterialBatchGetMaterial) (*response.ResponseMaterialBatchGetMaterial, error) {

	result := &response.ResponseMaterialBatchGetMaterial{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/material/batchget_material", options, nil, nil, result)

	return result, err
}

// 获取素材总数
// https://developers.weixin.qq.com/doc/offiaccount/Asset_Management/Get_the_total_of_all_materials.html
func (comp *Client) Stats(ctx context.Context) (*response.ResponseMaterialGetMaterialCount, error) {

	result := &response.ResponseMaterialGetMaterialCount{}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/material/get_materialcount", nil, nil, nil, result)

	return result, err

}

func (comp *Client) Upload(ctx context.Context, Type string, path string, query *object.StringMap, result interface{}) (interface{}, error) {

	_, err := os.Stat(path)
	if (err != nil && os.IsExist(err)) && (err != nil && os.IsPermission(err)) {
		return "", err
	}

	var files *object.HashMap
	if path != "" {
		files = &object.HashMap{
			"media": path,
		}
	}

	(*query)["type"] = Type

	form := &kernel.UploadForm{
		FileName: filepath.Base(path),
	}

	return comp.BaseClient.HttpUpload(ctx, comp.getApiByType(Type), files, form, query, nil, result)
}

func (comp *Client) getApiByType(Type string) string {

	switch Type {
	case "news_image":
		return "cgi-bin/media/uploadimg"
	default:
		return "cgi-bin/material/add_material"
	}

}
