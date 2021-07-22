package featureUnit

import (
	"github.com/ArtisanCloud/go-libs/fmt"
	"testing"
)

func Test_BridgeConfig(t *testing.T) {

	response, err := Payment.JSSDK.BridgeConfig("wx221017163973503c079bd96aea80e60000", false)
	if err != nil {
		t.Error("err msg:", err.Error())
	}

	if response == nil {

		t.Error("response nil")
	}

	fmt.Dump(response)

}
