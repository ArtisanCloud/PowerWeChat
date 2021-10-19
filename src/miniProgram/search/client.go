package search

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	response2 "github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/src/miniProgram/search/response"
)

type Client struct {
	*kernel.BaseClient
}

// 本接口提供基于小程序的站内搜商品图片搜索能力
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/search/search.imageSearch.html
func (comp *Client) ImageSearch(img []*power.HashMap) (*response.ResponseSearchImageSearch, error) {

	result := &response.ResponseSearchImageSearch{}

	data := &object.HashMap{
		"img": img,
	}

	_, err := comp.HttpPostJson("wxa/imagesearch", data, nil, nil, result)

	return result, err
}

// 小程序内部搜索API提供针对页面的查询能力
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/search/search.siteSearch.html
func (comp *Client) SiteSearch(keyword string, nextPageInfo string) (*response.ResponseSearchSiteSearch, error) {

	result := &response.ResponseSearchSiteSearch{}

	data := &object.HashMap{
		"keyword":        keyword,
		"next_page_info": nextPageInfo,
	}

	_, err := comp.HttpPostJson("wxa/sitesearch", data, nil, nil, result)

	return result, err
}

// 小程序开发者可以通过本接口提交小程序页面url及参数信息(不要推送webview页面)
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/search/search.submitPages.html
func (comp *Client) SubmitPages(pages []*power.HashMap) (*response2.ResponseMiniProgram, error) {

	result := &response2.ResponseMiniProgram{}

	data := &object.HashMap{
		"pages":        pages,
	}

	_, err := comp.HttpPostJson("wxa/search/wxaapi_submitpages", data, nil, nil, result)

	return result, err
}