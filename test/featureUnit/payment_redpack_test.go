package featureUnit

import (
	"github.com/ArtisanCloud/go-libs/fmt"
	"testing"
)

func Test_Redpack_Send(t *testing.T) {

	fmt.Dump(Payment)
	//para := &request.RequestMessageSendText{
	//	&request.RequestMessageSend{
	//		ToUser:                 "michaelhu",
	//		MsgType:                "text",
	//		AgentID:                1000004,
	//		Safe:                   0,
	//		EnableIDTrans:          0,
	//		EnableDuplicateCheck:   0,
	//		DuplicateCheckInterval: 1800,
	//	},
	//	&request.RequestText{
	//		Content: "您已经成功测试收到系统推送消息",
	//	},
	//}
	//
	//response := Work.Message.Send(para)
	//
	//if response == nil || response.ResponseWork == nil {
	//	t.Error("response nil")
	//} else if response.ErrCode != 0 {
	//	t.Error("response error message as :", response.ErrMSG)
	//}

}
