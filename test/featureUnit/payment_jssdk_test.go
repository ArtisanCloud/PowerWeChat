package featureUnit

import (
	"github.com/ArtisanCloud/go-libs/fmt"
	"github.com/ArtisanCloud/go-libs/object"
	"testing"
)

func Test_BridgeConfig(t *testing.T) {


	response,err:= Payment.JSSDK.BridgeConfig("123" , false)
	if err!=nil{
		t.Error("err msg:",err.Error())
	}

	fmt.Dump(response.(*object.StringMap))
	t.Error("response nil")
	//if response == nil || response.ResponseWork == nil {
	//	t.Error("response nil")
	//} else if response.ErrCode != 0 {
	//	t.Error("response error message as :", response.ErrMSG)
	//}

}
