package request

type Filter struct {
	IsToAll bool `json:"is_to_all"`
	TagID   int  `json:"tag_id"`
}

type Reception struct {
	ToUser []string `json:"touser"`
	Filter *Filter  `json:"filter"`
}

// ---------------------------------------------
