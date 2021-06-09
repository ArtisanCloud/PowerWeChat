package main

import (
	"fmt"
	fmt2 "github.com/ArtisanCloud/go-libs/fmt"
	"github.com/ArtisanCloud/go-wechat/src/work"
)

func main() {

	fmt.Printf("hello Wechat! \n")

	config := GetConfig()

	app := work.NewWork(config)
	//fmt2.Dump(app)

	//token := app.AccessToken.GetToken()
	//fmt2.Dump(token)
	//fmt2.Dump(app.GetConfig())

	//cType := reflect.TypeOf((*app.Components)["base"].(*base.Client))
	//fmt.Printf("kind %s \n", cType.Kind())
	//fmt.Printf("type %v \n", cType)

	ips := app.Base.GetCallbackIp()
	fmt2.Dump(ips)

}
