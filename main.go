package main

import (
	"fmt"
	fmt2 "github.com/ArtisanCloud/go-libs/fmt"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-wechat/src/work"
	"github.com/ArtisanCloud/go-wechat/src/work/auth"
)

func main() {

	fmt.Printf("hello Wechat! \n")

	config := object.HashMap{
		"corp_id": "ww454dfb9d6f6d432a",
		//'agent_id' : 100020,
		"secret": "",
	}

	app := work.NewWork(config)
	fmt2.Dump(app)
	components := app.GetComponents()
	token := (*components)["access_token"].(*auth.AccessToken).BaseAccessToken.GetToken()
	fmt2.Dump(token)
	//fmt2.Dump(app.Container.GetConfig())

}
