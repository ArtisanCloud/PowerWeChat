package request

type RequestPageInfo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	PageUrl     string `json:"page_url"`
	Comment     string `json:"comment"`
	IconUrl     string `json:"icon_url"`
}

type RequestPageUpdate struct {

	*RequestPageInfo

	PageID       int `json:"page_id"`

}