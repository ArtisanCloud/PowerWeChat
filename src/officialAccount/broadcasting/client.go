package broadcasting

import "github.com/ArtisanCloud/PowerWeChat/src/kernel"

const BROADCASTING_PREVIEW_BY_OPENID string = "touser"
const BROADCASTING_PREVIEW_BY_NAME string = "towxname"

type Client struct {
	*kernel.BaseClient
}

// 根据标签进行群发
// https://developers.weixin.qq.com/doc/offiaccount/Message_Management/Batch_Sends_and_Originality_Checks.html
//func (comp *Client)Send(message )  {
//
//}
