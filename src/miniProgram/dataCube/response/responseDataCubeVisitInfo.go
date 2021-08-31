package response

import "github.com/ArtisanCloud/go-libs/object"

type ResponseDataCubeVisitInfo struct {
	RefDate    string            `json:"ref_date"`
	VisitUVNew []*object.HashMap `json:"visit_uv_new"`
	VisitUV    []*object.HashMap `json:"visit_uv"`
}
