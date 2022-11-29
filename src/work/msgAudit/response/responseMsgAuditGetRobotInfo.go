package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

type RobotData struct {
	RobotID       string `json:"robot_id"`
	Name          string `json:"name"`
	CreatorUserID string `json:"creator_userid"`
}

type ResponseMsgAuditGetRobotInfo struct {
	*response.ResponseWork

	Data *RobotData `json:"data"`
}
