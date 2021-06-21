package server

import (
	"github.com/ArtisanCloud/go-wechat/src/kernel"
)

type Guard struct {
	
}


func NewGuard(app *kernel.ApplicationInterface) *Guard {
	//config := (*app).GetContainer().GetConfig()

	guard := &Guard{
		//HttpRequest: http.NewHttpRequest(config),
		//App:         app,
		//Token:       token,
	}
	return guard

}