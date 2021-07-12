package featureUnit

import (
	"github.com/ArtisanCloud/power-wechat/src/work/message/request"
	"testing"
)

func Test_Message_Send_Text(t *testing.T) {

	para := &request.RequestMessageSendText{
		&request.RequestMessageSend{
			ToUser:                 "michaelhu",
			MsgType:                "text",
			AgentID:                1000004,
			Safe:                   0,
			EnableIDTrans:          0,
			EnableDuplicateCheck:   0,
			DuplicateCheckInterval: 1800,
		},
		&request.RequestText{
			Content: "您已经成功测试收到系统推送消息",
		},
	}

	response := Work.Message.Send(para)

	if response == nil || response.ResponseWX == nil {
		t.Error("response nil")
	} else if response.ErrCode != 0 {
		t.Error("response error message as :", response.ErrMSG)
	}

}
