package featureUnit

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/fmt"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/messageTemplate/request"
	"testing"
)

func Test_ExternalContact_Add_Msg_Template(t *testing.T) {

	msg := &request.RequestAddMsgTemplate{
		ChatType: "",
		ExternalUserID: []string{
			"wm_ViZBwAApoZUCOn3JeqlfW1YUme5pg",
		},
		Sender: "WangChaoYi",
		Text: &request.TextOfMessage{
			Content: "新的活动现在开始啦！",
		},
		Attachments: nil,
	}

	response, err := Work.ExternalContactMessageTemplate.AddMsgTemplate(context.Background(), msg)

	if err != nil {
		t.Error("error:", err.Error())
	}

	if response == nil || response.ResponseWork.ErrCode != 0 {
		t.Error("response nil")
	} else if response.ErrCode != 0 {
		t.Error("response error uniformMessage as :", response.ErrMsg)
	}

	fmt.Dump(response)
}
