package featureUnit

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/fmt"
	"testing"
)

func Test_ExternalContact_Get_Contact_Way(t *testing.T) {

	response, _ := Work.ExternalContactContactWay.Get(context.Background(), "008dc067bf677e5f03df89ce49bea25a")

	if response == nil {
		t.Error("response nil")
	} else if response.ErrCode != 0 {
		t.Error("response error uniformMessage as :", response.ErrMsg)
	}

	fmt.Dump(response)

}
