package request

type RequestListLink struct {
	Cursor string `json:"cursor"`
	Limit  int    `json:"limit"`
}
