package response

type ResponseGetToken struct {
	ResponseWork

	AccessToken            string  `json:"access_token,omitempty"`
	ComponentAccessToken   string  `json:"component_access_token,omitempty"`
	AuthorizerAccessToken  string  `json:"authorizer_access_token,omitempty"`
	AuthorizerRefreshToken string  `json:"authorizer_refresh_token,omitempty"`
	SuiteAccessToken       string  `json:"suite_access_token,omitempty"`
	ProviderAccessToken    string  `json:"provider_access_token,omitempty"`
	ExpiresIn              float64 `json:"expires_in"`
}
