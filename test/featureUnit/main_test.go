package featureUnit

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work"
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

func GetWorkConfig() *work.UserConfig {
	agentID, _ := strconv.Atoi(os.Getenv("agent_id"))
	return &work.UserConfig{
		CorpID:  os.Getenv("corp_id"),
		AgentID: agentID,
		Secret:  os.Getenv("secret"),

		ResponseType: os.Getenv("array"),
		Log: work.Log{
			"debug",
			"./wechat.log",
			"develop",
		},

		OAuth: work.OAuth{
			Callback: os.Getenv("app_oauth_callback_url"),
			Scopes:   []string{},
		},
		//HttpDebug: true,
		Debug: true,

		// server config
		Token:  os.Getenv("token"),
		AESKey: os.Getenv("aes_key"),
	}
}

func GetPaymentConfig() *payment.UserConfig {
	return &payment.UserConfig{
		//"corp_id":        os.Getenv("corp_id"),
		//"secret":         os.Getenv("secret"),
		AppID:       os.Getenv("app_id"),
		MchID:       os.Getenv("mch_id"),
		MchApiV3Key: os.Getenv("mch_api_v3_key"),
		Key:         os.Getenv("key"),
		CertPath:    os.Getenv("wx_cert_path"),
		KeyPath:     os.Getenv("wx_key_path"),
		SerialNo:    os.Getenv("serial_no"),

		ResponseType: os.Getenv("array"),
		Log: payment.Log{
			Level: "debug",
			File:  "./wechat.log",
		},
		Http: payment.Http{
			Timeout: 30.0,
			BaseURI: "https://api.mch.weixin.qq.com",
		},

		NotifyURL: os.Getenv("notify_url"),
		HttpDebug: true,
		//Debug: true,
		//"sandbox": true,

		// server config
		//Token:            os.Getenv("token"),
		//AESKey:           os.Getenv("aes_key"),

	}
}

func TestInit(t *testing.T) {
	TestInitWork(t)
	TestInitPayment(t)
}

func TestInitWork(t *testing.T) {
	config := GetWorkConfig()
	Work, _ = work.NewWork(config)

}
func TestInitPayment(t *testing.T) {
	config := GetPaymentConfig()
	Payment, _ = payment.NewPayment(config)
}
