package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
	"github.com/ArtisanCloud/powerwechat/src/work/menu/request"
)

type ResponseMenuGet struct {
	*response.ResponseWork

	// 企业微信官方：返回结果与请参考菜单创建接口
	// https://open.work.weixin.qq.com/api/doc/90000/90135/90232
	request.RequestMenuSet
}
