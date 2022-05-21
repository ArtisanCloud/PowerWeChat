package response

import "github.com/ArtisanCloud/PowerWeChat/src/kernel/response"

type Inner struct {
	Name string `json:"name"`
}

type Exter struct {
	InnerList []*Inner `json:"inner_list"`
}

type Qualify struct {
	ExterList []*Exter `json:"exter_list"`
}

type Category struct {
	ID            int      `json:"id"`
	Name          string   `json:"name"`
	Level         int      `json:"level"`
	Children      []int    `json:"children"`
	Father        int      `json:"father,omitempty"`
	Qualify       *Qualify `json:"qualify,omitempty"`
	Scene         int      `json:"scene,omitempty"`
	SensitiveType int      `json:"sensitive_type,omitempty"`
}

type AllCategoryInfo struct {
	Categories []*Category `json:"categories"`
}

type Store struct {
	AllCategoryInfo *AllCategoryInfo `json:"all_category_info"`
}

type ResponseStoreCategory struct {
	*response.ResponseOfficialAccount

	Data Store `json:"data"`
}
