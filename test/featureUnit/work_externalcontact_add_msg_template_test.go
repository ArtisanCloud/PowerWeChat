package featureUnit

import (
	"github.com/ArtisanCloud/PowerLibs/fmt"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/work/externalContact/messageTemplate/request"
	"testing"
)

func Test_ExternalContact_Add_Msg_Template(t *testing.T) {

	msg := &request.RequestAddMsgTemplate{
		ChatType: "",
		ExternalUserID: []string{
			"wm_ViZBwAApoZUCOn3JeqlfW1YUme5pg",
		},
		Sender: "WangChaoYi",
		Text: request.TextOfMessage{
			Content: "新的活动现在开始啦！",
		},
		Attachments: []*power.HashMap{},
	}

	response, err := Work.ExternalContactMessageTemplate.AddMsgTemplate(msg)

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
