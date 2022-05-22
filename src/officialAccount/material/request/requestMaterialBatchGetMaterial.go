package request

type RequestMaterialBatchGetMaterial struct {
	Type   string `json:"type"`
	Offset int64 `json:"offset"`
	Count  int64 `json:"count"`
}
