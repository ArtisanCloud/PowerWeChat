package response

type ResponseGetToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	*ResponseWX
}