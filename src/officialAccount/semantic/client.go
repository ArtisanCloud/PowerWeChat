package semantic

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

// 语义接口
// https://developers.weixin.qq.com/doc/offiaccount/Intelligent_Interface/Natural_Language_Processing.html
func (comp *Client) Query(ctx context.Context, keyword string, categories string, optional *power.HashMap) (interface{}, error) {

	config := (*comp.BaseClient.App).GetConfig()
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

	return comp.BaseClient.HttpPostJson(ctx, "semantic/semproxy/search", params, nil, nil, nil)

}
