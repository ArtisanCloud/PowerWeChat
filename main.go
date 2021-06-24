package main

import (
	"fmt"
	fmt2 "github.com/ArtisanCloud/go-libs/fmt"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-wechat/src/work"
	"os"
	"strconv"
)

func GetConfig() *object.HashMap {
	agentID, _ := strconv.Atoi(os.Getenv("agent_id"))
	return &object.HashMap{
		"corp_id":  os.Getenv("corp_id"),
		"agent_id": agentID,
		"secret":   os.Getenv("secret"),

		"response_type": os.Getenv("array"),
		"log": &object.StringMap{

			"level": "debug",
			"file":  "./wechat.log",
		},

		"oauth.callback": "https://wechat-work-sso-2.spacecycle.cn/callback/authorized/user",
		"oauth.scopes":   []string{},
		"debug":          true,

		// server config
		"token":   os.Getenv("token"),
		"aes_key": os.Getenv("aes_key"),
	}
}

func main() {

	fmt.Printf("hello Wechat! \n")

	config := GetConfig()

	app, err := work.NewWork(config, nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	response := app.ExternalContact.List("Matt")
	fmt2.Dump(response)

}
