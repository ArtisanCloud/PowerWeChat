package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponsePluginGetPluginList struct {
	*response.ResponseMiniProgram
	PluginList []*power.HashMap `json:"plugin_list"`
}
