package main

import (
	"fmt"
	fmt2 "github.com/ArtisanCloud/go-libs/fmt"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-wechat/src/work"
)

func main() {

	fmt.Printf("hello Wechat! \n")

	config := &object.HashMap{
		"corp_id": "ww454dfb9d6f6d432a",
		//'agent_id' : 100020,
		"secret": "9iZxAhtyKC9OB0ofN8Qrvjy4Kk80nuoJ7nIbocE5j8M",
	}

	app := work.NewWork(config)
	//fmt2.Dump(app)
	//components := app.GetComponents()
	//token := (*components)["access_token"].(*auth.AccessToken).GetToken()
	//fmt2.Dump(token)
	fmt2.Dump(app.GetConfig())

	//cType := reflect.TypeOf((*app.Components)["base"].(*base.Client))
	//fmt.Printf("kind %s \n", cType.Kind())
	//fmt.Printf("type %v \n", cType)

	//ips :=(*app.Components)["base"].(*base.Client).GetCallbackIp()
	//fmt2.Dump(ips)

}
