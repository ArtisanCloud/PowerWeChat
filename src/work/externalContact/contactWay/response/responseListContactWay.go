package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
)

// ------------------------------------------------------------------------------------
type ContactWayID struct {
	ConfigID string `json:"config_id"`
}

type ResponseListContactWay struct {
	response.ResponseWork

	ContactWayIDs []*ContactWayID `json:"contact_way"`
	NextCursor    string          `json:"next_cursor"`
}
