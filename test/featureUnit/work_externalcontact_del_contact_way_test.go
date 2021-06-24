package featureUnit

import (
	"github.com/ArtisanCloud/go-libs/fmt"
	"testing"
)

func Test_ExternalContact_Del_Contact_Way(t *testing.T) {

	response := Work.ContactWay.Delete("f3626f74a7f94784115b0b8a729c471f")

	if response == nil {
		t.Error("response nil")
	} else if response.ErrCode != 0 {
		t.Error("response error message as :", response.ErrMSG)
	}

	fmt.Dump(response)

}
