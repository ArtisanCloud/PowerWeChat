package main

import (
	"fmt"
	fmt2 "github.com/ArtisanCloud/go-libs/fmt"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-wechat/src/work"
	"github.com/ArtisanCloud/go-wechat/src/work/externalContact/request"
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
		"log": object.StringMap{

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
	//fmt2.Dump(app)
	//fmt2.Dump(app.GetConfig())

	//token := app.AccessToken.GetToken()
	//fmt2.Dump(token)

	//cType := reflect.TypeOf((*app.Components)["base"].(*base.Client))
	//fmt.Printf("kind %s \n", cType.Kind())
	//fmt.Printf("type %v \n", cType)

	//ips := app.Base.GetCallbackIp()
	//fmt2.Dump(ips)
	//domainIps := app.Base.GetAPIDomainIP()
	//fmt2.Dump(domainIps)

	//fmt2.Dump(app.Department.List())

	para := &request.RequestAddContactWay{
		Type:          1,
		Scene:         1,
		Style:         1,
		Remark:        "渠道客户",
		SkipVerify:    true,
		State:         "teststate",
		User:          []string{"michaelhu"},
		Party:         []int{2, 3},
		IsTemp:        true,
		ExpiresIn:     86400,
		ChatExpiresIn: 86400,
		UnionID:       "",
		Conclusions: &object.HashMap{
			"text": object.StringMap{
				"content": "bye bye",
			},
		},
	}

	response := app.ContactWay.Create(para)
	fmt2.Dump(response)

}
