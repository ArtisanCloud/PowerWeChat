package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type ResponseHeaderMedia struct {
	response.ResponseWork

	Connection         []string `header:"Connection,omitempty" json:"Connection,omitempty"`                   //: close
	ContentType        []string `header:"Content-Type,omitempty" json:"Content-Type,omitempty"`               //: image/jpeg
	ContentDisposition []string `header:"Content-disposition,omitempty" json:"Content-disposition",omitempty` //: attachment; filename="MEDIA_ID.jpg"
	Date               []string `header:"Date,omitempty" json:"Date"`                                         //: Sun, 06 Jan 2013 10:20:18 GMT
	CacheControl       []string `header:"Cache-Control,omitempty" json:"Cache-Control,omitempty"`             //: no-cache, must-revalidate
	ContentLength      []string `header:"Content-Length,omitempty" json:"Content-Length,omitempty"`           //: 339721
	Content            []string `header:"Content,omitempty" json:"content,omitempty"`
}
