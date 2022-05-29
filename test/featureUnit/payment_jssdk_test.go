package featureUnit

import (
	"github.com/ArtisanCloud/PowerLibs/v2/fmt"
	"testing"
)

func Test_BridgeConfig(t *testing.T) {

	response, err := Payment.JSSDK.BridgeConfig("wx30003559330979ff187dcbf5c4704b0000", false)
	if err != nil {
		t.Error("err msg:", err.Error())
	}

	if response == nil {

		t.Error("response nil")
	}

	fmt.Dump(response)

}
