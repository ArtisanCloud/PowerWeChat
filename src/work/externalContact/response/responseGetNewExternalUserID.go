package response

import "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"

type ResponseGetNewExternalUserID struct {
	response.ResponseWork
	Items []NewExternalUserID `json:"items,omitempty"`
}

type NewExternalUserID struct {
	ExternalUserID    string `json:"external_userid,omitempty"`
	NewExternalUserID string `json:"new_external_userid,omitempty"`
}
