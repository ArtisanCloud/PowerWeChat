package request

type RequestSet struct {
	ApiName   string   `json:"api_name"`
	Content   string   `json:"content"`
	PicList   []string `json:"pic_list"`
	VideoList []string `json:"video_list"`
	URLList   []string `json:"url_list"`
}
