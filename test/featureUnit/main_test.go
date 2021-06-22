package featureUnit

import (
	"github.com/ArtisanCloud/go-libs/fmt"
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-wechat/src/work"
	"log"
	"os"
	"strconv"
	"testing"
)

var Work *work.Work

func TestMain(m *testing.M) {

	log.Println("Before Test: [++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

	// init test app

	exitVal := m.Run()

	log.Println("After Test: ------------------------------------------------------------------]")

	os.Exit(exitVal)

}

func GetConfig() *object.HashMap {
	agentID, _ := strconv.Atoi(os.Getenv("agent_id"))
	fmt.Dump(os.Getenv("secret"))
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

func TestInit(t *testing.T) {
	TestInitWork(t)
}

func TestInitWork(t *testing.T) {
	config := GetConfig()
	Work, _ = work.NewWork(config,nil)

}
