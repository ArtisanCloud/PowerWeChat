package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
)

type ResponseGetUnassignedList struct {
	response.ResponseWork

	ExternalContactList []*ResponseUnassignedInfo `json:"info"`
	IsLast              bool                      `json:"is_last"`
	NextCursor          string                    `json:"next_cursor"`
}

type ResponseUnassignedInfo struct {
	HandoverUserid string `json:"handover_userid"` // "zhangsan",
	ExternalUserid string `json:"external_userid"` // "woAJ2GCAAAd4uL12hdfsdasassdDmAAAAA",
	DimissionTime  int    `json:"dimission_time"`  // 1550838571
}
