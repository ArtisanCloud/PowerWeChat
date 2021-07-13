package response

type ResponseHeaderMedia struct {
	Connection         string `header:"Connection"`          //: close
	ContentType        string `header:"Content-Type"`        //: image/jpeg
	ContentDisposition string `header:"Content-disposition"` //: attachment; filename="MEDIA_ID.jpg"
	Date               string `header:"Date"`                //: Sun, 06 Jan 2013 10:20:18 GMT
	CacheControl       string `header:"Cache-Control"`       //: no-cache, must-revalidate
	ContentLength      string `header:"Content-Length"`      //: 339721
	Content            string `header:""`
}
