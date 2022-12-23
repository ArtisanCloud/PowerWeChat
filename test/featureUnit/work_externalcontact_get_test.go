package featureUnit

import (
	"github.com/ArtisanCloud/PowerLibs/v3/fmt"
	"testing"
)

func Test_ExternalContact_Get_Batch_ByUser(t *testing.T) {

	userID := []string{"siyuan-notify"}
	cursor := ""
	limit := 100

	response, _ := Work.ExternalContact.BatchGet(userID, cursor, limit)

	if response == nil {
		t.Error("response nil")
	} else if response.ErrCode != 0 {
		t.Error("response error uniformMessage as :", response.ErrMSG)
	}

	fmt.Dump(response)

}

func Test_ExternalContact_Get_Unassigned(t *testing.T) {

	pageID := 1
	pageSize := 50

	response, _ := Work.ExternalContact.GetUnassigned(pageID, pageSize)

	if response == nil {
		t.Error("response nil")
	} else if response.ErrCode != 0 {
		t.Error("response error uniformMessage as :", response.ErrMSG)
	}

	fmt.Dump(response)

}

func Test_ExternalContact_Transfer(t *testing.T) {

	handoverUserID := ""
	takeoverUserID := ""
	externalUserID := []string{""}

	response, _ := Work.ExternalContact.Transfer(externalUserID, handoverUserID, takeoverUserID)

	if response == nil {
		t.Error("response nil")
	} else if response.ErrCode != 0 {
		t.Error("response error uniformMessage as :", response.ErrMSG)
	}

	fmt.Dump(response)

}
