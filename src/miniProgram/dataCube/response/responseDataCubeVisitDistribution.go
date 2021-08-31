package response

import (
	"github.com/ArtisanCloud/go-libs/object"
)

type ResponseDataCubeVisit struct {
	RefDate string            `json:"ref_date"`
	List    []*object.HashMap `json:"list"`
}
