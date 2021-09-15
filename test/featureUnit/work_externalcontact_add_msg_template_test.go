package featureUnit

import (
	"github.com/ArtisanCloud/go-libs/fmt"
	"github.com/ArtisanCloud/go-libs/object"
	request2 "github.com/ArtisanCloud/power-wechat/src/work/externalContact/message/request"
	"testing"
)

func Test_ExternalContact_Add_Msg_Template(t *testing.T) {

	msg := &object.HashMap{
		"external_userid": []string{
			"wm_ViZBwAApoZUCOn3JeqlfW1YUme5pg",
		},
		"sender": "WangChaoYi",
		//"sender": "Matt",
		"text": request2.TextOfMessage{
			Content: "新的活动现在开始啦！",
		},
		"attachments": []interface{}{
			//request.ImageOfMessage{
			//	MsgType: "image",
			//	Image: request.Image{
			//		MediaID: "",
			//		PicURL:  "",
			//	},
			//},
			request2.LinkOfMessage{
				MsgType: "link",
				Link: request2.Link{
					Title:  "活动地址",
					PicURL: "https://baidu.com",
					Desc:   "好奖品等你来拿",
					URL:    "https://baidu.com",
				},
			},
			//request.MiniProgramOfMessage{
			//	MsgType: "miniprogram",
			//	MiniProgram: request.MiniProgram{
			//		Title:      "消息标题",
			//		PicMediaID: "MEDIA_ID",
			//		AppID:      "wx8bd80126147df384",
			//		Page:       "/path/index",
			//	},
			//},
			//request.VideoOfMessage{
			//	MsgType: "video",
			//	Video: request.Video{
			//		MediaID: "MEDIA_ID",
			//	},
			//},
		},
	}

	response, err := Work.ExternalContactMessage.Submit(msg)

	if err != nil {
		t.Error("error:", err.Error())
	}

	if response == nil || response.ResponseWork == nil {
		t.Error("response nil")
	} else if response.ErrCode != 0 {
		t.Error("response error uniformMessage as :", response.ErrMSG)
	}

	fmt.Dump(response)

}
