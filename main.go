package main

import (
	"fmt"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel/power"
	"os"
	"strconv"
)

func GetConfig() *power.HashMap {
	agentID, _ := strconv.Atoi(os.Getenv("agent_id"))
	return &power.HashMap{
		"corp_id":  os.Getenv("corp_id"),
		"agent_id": agentID,
		"secret":   os.Getenv("secret"),

		"response_type": os.Getenv("array"),
		"log": &object.StringMap{

			"level": "debug",
			"file":  "./wechat.log",
		},

		"oauth.callback": os.Getenv("oauth_callback"),
		"oauth.scopes":   []string{},
		"debug":          true,

		// server config
		"token":   os.Getenv("token"),
		"aes_key": os.Getenv("aes_key"),
	}
}

func main() {

	fmt.Printf("hello Wechat! \n")

	//config := GetConfig()
	//
	//app, err := work.NewWork(config, nil)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}

	//response := app.ExternalContact.List("Matt")
	//fmt2.Dump(response)

	//response, err := app.Server.Serve()
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt2.Dump(response)

}
