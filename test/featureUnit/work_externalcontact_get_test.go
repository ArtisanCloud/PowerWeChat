package featureUnit

import (
	"github.com/ArtisanCloud/go-libs/fmt"
	"testing"
)

func Test_ExternalContact_Get_Batch_ByUser(t *testing.T) {

	userID := "siyuan-notify"
	cursor := ""
	limit := "100"

	response := Work.ExternalContact.BatchGet(userID, cursor, limit)

	if response == nil {
		t.Error("response nil")
	} else if response.ErrCode != 0 {
		t.Error("response error message as :", response.ErrMSG)
	}

	fmt.Dump(response)

}

func Test_ExternalContact_Get_Unassigned(t *testing.T) {

	pageId := 1
	pageSize := 50

	response := Work.ExternalContact.GetUnassigned(pageId, pageSize)

	if response == nil {
		t.Error("response nil")
	} else if response.ErrCode != 0 {
		t.Error("response error message as :", response.ErrMSG)
	}

	fmt.Dump(response)

}

func Test_ExternalContact_Transfer(t *testing.T) {

	handoverUserId := ""
	takeoverUserId := ""
	externalUserId := []string{""}

	response := Work.ExternalContact.Transfer(externalUserId, handoverUserId, takeoverUserId)

	if response == nil {
		t.Error("response nil")
	} else if response.ErrCode != 0 {
		t.Error("response error message as :", response.ErrMSG)
	}

	fmt.Dump(response)

}
