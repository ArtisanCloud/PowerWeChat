package featureUnit

import (
	"github.com/ArtisanCloud/go-libs/fmt"
	"github.com/ArtisanCloud/powerwechat/src/kernel/power"
	request2 "github.com/ArtisanCloud/powerwechat/src/work/externalContact/contactWay/request"
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
		Conclusions: &power.HashMap{
			"text": &power.StringMap{
				"content": "bye bye",
			},
		},
	}

	response,_ := Work.ExternalContactContactWay.Add(para)

	if response == nil || response.ResponseWork == nil {
		t.Error("response nil")
	} else if response.ErrCode != 0 {
		t.Error("response error uniformMessage as :", response.ErrMSG)
	}

	fmt.Dump(response)

}
