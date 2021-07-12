package featureUnit

import (
	"github.com/ArtisanCloud/go-libs/fmt"
	"testing"
)

//func Test_Media_Upload_Image(t *testing.T) {
//
//	wxResponse := response.ResponseWX{}
//	homePath, _ := os.UserHomeDir()
//	response := Work.Media.UploadImage(homePath+"/Desktop/123.jpg", nil, &wxResponse)
//
//	if response == nil {
//		t.Error("response nil")
//	} else if wxResponse.ErrCode != 0 {
//		t.Error("response error message as :", wxResponse.ErrMSG)
//	}
//
//	fmt.Dump(wxResponse, response)
//
//}

func Test_Media_Get(t *testing.T) {
	response, _:=Work.Media.Get("3qtl2wILoNc9nPJuKzNxWyp5H8YKu2NfZytgIgzDhxLBKFUb-6m6kijaUMggqklQW")

	if response == nil {
		t.Error("response nil")
	}

	fmt.Dump(response, response)
}
