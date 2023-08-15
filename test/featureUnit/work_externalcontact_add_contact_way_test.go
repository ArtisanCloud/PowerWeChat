package featureUnit

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/fmt"
	request2 "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/contactWay/request"
	"testing"
)

func Test_ExternalContact_Add_Contact_Way(t *testing.T) {

	para := &request2.RequestAddContactWay{
		Type:          1,
		Scene:         1,
		Style:         1,
		Remark:        "渠道客户",
		SkipVerify:    true,
		State:         "",
		User:          []string{"WangChaoYi"},
		Party:         []int{1},
		IsTemp:        true,
		ExpiresIn:     86400,
		ChatExpiresIn: 86400,
		UnionID:       "",
		Conclusions:   nil,
	}

	response, _ := Work.ExternalContactContactWay.Add(context.Background(), para)

	if response == nil || response.ResponseWork.ErrCode != 0 {
		t.Error("response nil")
	} else if response.ErrCode != 0 {
		t.Error("response error uniformMessage as :", response.ErrMsg)
	}

	fmt.Dump(response)

}
