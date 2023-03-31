package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type ResponseHeaderMedia struct {
	response.ResponseOfficialAccount

	Connection         []string `header:"Connection" json:"Connection"`                   //: close
	ContentType        []string `header:"Content-Type" json:"Content-Type"`               //: image/jpeg
	ContentDisposition []string `header:"Content-disposition" json:"Content-disposition"` //: attachment; filename="MEDIA_ID.jpg"
	Date               []string `header:"Date" json:"Date"`                               //: Sun, 06 Jan 2013 10:20:18 GMT
	CacheControl       []string `header:"Cache-Control" json:"Cache-Control"`             //: no-cache, must-revalidate
	ContentLength      []string `header:"Content-Length" json:"Content-Length"`           //: 339721
	Content            []string `header:"Content" json:"content"`
}
