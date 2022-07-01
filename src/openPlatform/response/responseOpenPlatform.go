package response

import "encoding/xml"

type ResponseVerifyTicket struct {
	XMLName               xml.Name `xml:"xml"`
	Text                  string   `xml:",chardata"`
	AppID                 string   `xml:"AppId"`
	CreateTime            string   `xml:"CreateTime"`
	InfoType              string   `xml:"InfoType"`
	ComponentVerifyTicket string   `xml:"ComponentVerifyTicket"`
}
