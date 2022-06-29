package response

type ResponseGetToken struct {
	ResponseWork

	AccessToken          string  `json:"access_token,omitempty"`
	ComponentAccessToken string  `json:"component_access_token,omitempty"`
	ExpiresIn            float64 `json:"expires_in"`
}
