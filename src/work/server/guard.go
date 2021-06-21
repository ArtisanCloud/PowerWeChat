package server

import "github.com/ArtisanCloud/go-libs/http"

type Guard struct {
	
}


func NewGuard(app *ApplicationInterface) *Guard {
	config := (*app).GetContainer().GetConfig()


	guard := &Guard{
		HttpRequest: http.NewHttpRequest(config),
		App:         app,
		Token:       token,
	}
	return guard

}