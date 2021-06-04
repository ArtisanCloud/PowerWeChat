package kernel

type AccessTokenInterface interface {
	GetToken() string
	Refresh() AccessTokenInterface
	
}

type AccessToken struct {

}