package response

import "github.com/ArtisanCloud/PowerWeChat/src/kernel/response"

type GroupData struct {
	GroupID   int    `json:"group_id"`
	GroupName string `json:"group_name"`
}

type ResponseGroup struct {
	*response.ResponseOfficialAccount

	Data *GroupData `json:"data"`
}

type GroupList struct {
	TotalCount int          `json:"total_count"`
	Groups     []*GroupData `json:"groups"`
}

// ----------------------------------------------

type ResponseGroupList struct {
	*response.ResponseOfficialAccount

	Data *GroupList `json:"data"`
}

// ----------------------------------------------

type GroupDetail struct {
	GroupID    int       `json:"group_id"`
	GroupName  string    `json:"group_name"`
	TotalCount int       `json:"total_count"`
	Devices    []*Device `json:"devices"`
}

type ResponseGroupDetail struct {
	*response.ResponseOfficialAccount

	Data *GroupDetail `json:"data"`
}

// ----------------------------------------------

type DataGroupAddDevice struct {
	GroupID    int      `json:"group_id"`
	GroupName  string   `json:"group_name"`
	TotalCount int      `json:"total_count"`
	Devices    []Device `json:"devices"`
}

type ResponseGroupAddDevices struct {
	*response.ResponseOfficialAccount

	Data *DataGroupAddDevice `json:"data"`
}


// ----------------------------------------------



// ----------------------------------------------
