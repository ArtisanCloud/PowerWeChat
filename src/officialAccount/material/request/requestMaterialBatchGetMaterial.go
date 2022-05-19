package request

type RequestMaterialBatchGetMaterial struct {
	Type   int64 `json:"type"`
	Offset int64 `json:"offset"`
	Count  int64 `json:"count"`
}
