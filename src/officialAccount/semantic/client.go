package semantic

import (
	"github.com/ArtisanCloud/PowerLibs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
)

type Client struct {
	*kernel.BaseClient
}

// 语义接口
// https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/Natural_Language_Processing.html
func (comp *Client) Query(keyword string, categories string, optional *power.HashMap) (interface{}, error) {

	config := (*comp.App).GetConfig()
	params := &object.HashMap{
		"query":    keyword,
		"category": categories,
		"appid":    config.GetString("app_id", ""),
	}

	objOptional, err := power.PowerHashMapToObjectHashMap(optional)
	if err != nil {
		return nil, err
	}
	params = object.MergeHashMap(params, objOptional)

	return comp.HttpPostJson("semantic/semproxy/search", params, nil, nil, nil)

}
