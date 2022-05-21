package response

import "github.com/ArtisanCloud/PowerWeChat/src/kernel/response"

type DataPage struct {
	PageID int `json:"page_id"`
}

type ResponsePage struct {
	*response.ResponseOfficialAccount

	Data *DataPage `json:"data"`
}


// ---------------------------------------------------

type Page struct {
	Comment     string `json:"comment"`
	Description string `json:"description"`
	IconURL     string `json:"icon_url"`
	PageID      int    `json:"page_id"`
	PageURL     string `json:"page_url"`
	Title       string `json:"title"`
}

type DataPageList struct {
	Pages []*Page`json:"pages"`
	TotalCount int `json:"total_count"`
}

type ResponsePageList struct {
	*response.ResponseOfficialAccount

	Data *DataPageList `json:"data"`
}