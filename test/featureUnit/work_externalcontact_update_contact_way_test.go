package featureUnit

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/fmt"
	request2 "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/contactWay/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/messageTemplate/request"
	"testing"
)

func Test_ExternalContact_Update_Contact_Way(t *testing.T) {

	para := &request2.RequestUpdateContactWay{

		Style:      1,
		Remark:     "渠道客户",
		SkipVerify: true,
		State:      "",
		User:       []string{"WangChaoYi"},
		Party:      []int{1},

		ExpiresIn:     86400,
		ChatExpiresIn: 86400,
		UnionID:       "",
		Conclusions: &request2.Conclusions{
			Text: &request.TextOfMessage{
				Content: "bye bye",
			},
		},
	}

	response, _ := Work.ExternalContactContactWay.Update(context.Background(), para)

	if response == nil || response.ResponseBase.Code == "0" {
		t.Error("response nil")
	} else if response.ErrCode != 0 {
		t.Error("response error uniformMessage as :", response.ErrMsg)
	}

	fmt.Dump(response)

}
