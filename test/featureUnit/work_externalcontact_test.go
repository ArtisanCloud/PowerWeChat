package featureUnit

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-wechat/src/work/externalContact/request"
	"testing"
)

func Test_ExternalContact_Add_Contact_Way(t *testing.T) {

	para := &request.RequestAddContactWay{
		Type:          1,
		Scene:         1,
		Style:         1,
		Remark:        "渠道客户",
		SkipVerify:    true,
		State:         "teststate",
		User:          []string{"michaelhu"},
		Party:         []int{2, 3},
		IsTemp:        true,
		ExpiresIn:     86400,
		ChatExpiresIn: 86400,
		UnionID:       "",
		Conclusions: &object.HashMap{
			"text": object.StringMap{
				"content": "bye bye",
			},
		},
	}

	response := Work.ContactWay.Create(para)

	if response == nil || response.ResponseWX == nil {
		t.Error("response nil")
	} else if response.ErrCode != 0 {
		t.Error("response error message as :", response.ErrMSG)
	}

}
