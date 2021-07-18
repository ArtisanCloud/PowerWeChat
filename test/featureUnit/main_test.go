package featureUnit

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/payment"
	"github.com/ArtisanCloud/power-wechat/src/work"
	"log"
	"os"
	"strconv"
	"testing"
)

var Work *work.Work
var Payment *payment.Payment

func TestMain(m *testing.M) {

	log.Println("Before Test: [++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

	// init test app

	exitVal := m.Run()

	log.Println("After Test: ------------------------------------------------------------------]")

	os.Exit(exitVal)

}

func GetWorkConfig() *object.HashMap {
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

		"oauth.callback": os.Getenv("oauth_callback"),
		"oauth.scopes":   []string{},
		"debug":          true,

		// server config
		"token":   os.Getenv("token"),
		"aes_key": os.Getenv("aes_key"),
	}
}

func GetPaymentConfig() *object.HashMap {
	return &object.HashMap{
		"corp_id":  os.Getenv("corp_id"),
		"secret":   os.Getenv("secret"),
		"app_id":   os.Getenv("app_id"),
		"mch_id":   os.Getenv("mch_id"),
		"key":      os.Getenv("key"),
		"cert_pat": os.Getenv("cert_pat"),
		"key_path": os.Getenv("key_path"),

		"response_type": os.Getenv("array"),
		"log": &object.StringMap{

			"level": "debug",
			"file":  "./wechat.log",
		},

		"http": object.HashMap{
			"timeout":  30.0,
			"base_uri": "https://api.mch.weixin.qq.com/v3",
		},

		"notify_url": os.Getenv("notify_url"),
		"debug":      true,
		//"sandbox": true,

		// server config
		"token":   os.Getenv("token"),
		"aes_key": os.Getenv("aes_key"),
	}
}

func TestInit(t *testing.T) {
	TestInitWork(t)
	TestInitPayment(t)
}

func TestInitWork(t *testing.T) {
	config := GetWorkConfig()
	Work, _ = work.NewWork(config, nil)

}
func TestInitPayment(t *testing.T) {
	config := GetPaymentConfig()
	Payment, _ = payment.NewPayment(config, nil)
}
