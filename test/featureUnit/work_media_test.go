package featureUnit

import (
	"github.com/ArtisanCloud/go-libs/fmt"
	"testing"
)

//func Test_Media_Upload_Image(t *testing.T) {
//
//	wxResponse := response2.ResponseUploadImage{}
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

//func Test_Media_Upload_Temp_Image(t *testing.T) {
//
//	wxResponse := response2.ResponseUploadMedia{}
//	homePath, _ := os.UserHomeDir()
//	response := Work.Media.UploadTempImage(homePath+"/Desktop/123.jpg", nil, &wxResponse)
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
	response, _ := Work.Media.Get("3i12vSikgRGsE9Xbs3e1CZH4b_X6cREZtuHe2nh8VZ4w")

	//if response == nil {
	//	t.Error("response nil")
	//}

	fmt.Dump(response)
}
