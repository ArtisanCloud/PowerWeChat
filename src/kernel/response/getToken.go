package response

type ResponseGetToken struct {
	AccessToken string  `json:"access_token"`
	ExpiresIn   float64 `json:"expires_in"`
	*ResponseWork
}
