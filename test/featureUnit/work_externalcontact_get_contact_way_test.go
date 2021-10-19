package featureUnit

import (
	"github.com/ArtisanCloud/PowerLibs/fmt"
	"testing"
)

func Test_ExternalContact_Get_Contact_Way(t *testing.T) {

	response,_ := Work.ExternalContactContactWay.Get("008dc067bf677e5f03df89ce49bea25a")

	if response == nil || response.ResponseWork == nil {
		t.Error("response nil")
	} else if response.ErrCode != 0 {
		t.Error("response error uniformMessage as :", response.ErrMSG)
	}

	fmt.Dump(response)

}
