package main

import (
	"fmt"
	fmt2 "github.com/ArtisanCloud/PowerLibs/v3/fmt"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openPlatform"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work"
	"os"
	"strconv"
)

func GetWorkConfig() *work.UserConfig {
	agentID, _ := strconv.Atoi(os.Getenv("wecom_agent_id"))
	return &work.UserConfig{
		CorpID:  os.Getenv("corp_id"),
		AgentID: agentID,
		Secret:  os.Getenv("wecom_secret"),

		ResponseType: os.Getenv("array"),
		Log: work.Log{
			Level: "debug",
			File:  "./wechat.log",
			ENV:   os.Getenv("work.env"),
		},

		OAuth: work.OAuth{
			Callback: os.Getenv("app_oauth_callback_url"),
			Scopes:   []string{},
		},
		Cache: kernel.NewRedisClient(&kernel.RedisOptions{
			Addr:     "127.0.0.1:6379",
			Password: "",
			DB:       1,
		}),

		//HttpDebug: true,
		Debug: true,

		// server config
		Token:  os.Getenv("app_message_token"),
		AESKey: os.Getenv("app_message_aes_key"),
	}
}

func GetPaymentConfig() *payment.UserConfig {
	return &payment.UserConfig{
		//"corp_id":        os.Getenv("corp_id"),
		//"secret":         os.Getenv("secret"),
		AppID:              os.Getenv("app_id"),
		MchID:              os.Getenv("mch_id"),
		MchApiV3Key:        os.Getenv("mch_api_v3_key"),
		Key:                os.Getenv("key"),
		CertPath:           os.Getenv("wx_cert_path"),
		KeyPath:            os.Getenv("wx_key_path"),
		SerialNo:           os.Getenv("serial_no"),
		CertificateKeyPath: os.Getenv("certificate_key_path"),
		WechatPaySerial:    os.Getenv("wechat_pay_serial"),
		RSAPublicKeyPath:   os.Getenv("rsa_public_key_path"),

		ResponseType: os.Getenv("array"),
		Log: payment.Log{
			Level: "debug",
			File:  "./wechat.log",
		},
		Http: payment.Http{
			Timeout: 30.0,
			BaseURI: "https://api.mch.weixin.qq.com",
		},

		Cache: kernel.NewRedisClient(&kernel.RedisOptions{
			Addr:     "127.0.0.1:6379",
			Password: "",
			DB:       1,
		}),
		NotifyURL: os.Getenv("notify_url"),
		HttpDebug: true,
		//Debug: true,
		//"sandbox": true,

		// server config
		//Token:            os.Getenv("token"),
		//AESKey:           os.Getenv("aes_key"),

	}
}

func GetMiniProgramConfig() *miniProgram.UserConfig {
	return &miniProgram.UserConfig{

		AppID:  os.Getenv("miniprogram_app_id"), // 小程序、公众号或者企业微信的appid
		Secret: os.Getenv("miniprogram_secret"), // 商户号 appID

		ResponseType: os.Getenv("array"),
		Log: miniProgram.Log{
			Level: "debug",
			File:  "./wechat.log",
		},
		Cache: kernel.NewRedisClient(&kernel.RedisOptions{
			Addr:     "127.0.0.1:6379",
			Password: "",
			DB:       1,
		}),
		HttpDebug: true,
		//Debug: true,
		//"sandbox": true,

	}

}

func GetOpenPlatformConfig() *openPlatform.UserConfig {
	return &openPlatform.UserConfig{
		AppID:        "123",
		Secret:       "321",
		AuthCode:     "123",
		AESKey:       "321",
		ResponseType: os.Getenv("array"),
		Log: openPlatform.Log{
			Level: "debug",
			File:  "./wechat.log",
		},
		Cache: kernel.NewRedisClient(&kernel.RedisOptions{
			Addr:     "127.0.0.1:6379",
			Password: "",
			DB:       1,
		}),
		//OAuth:        "",
		//HttpDebug:    "",
		//Debug:        "",
		//NotifyURL:    "",
		//Sandbox:      "",
	}
}

func main() {

	fmt.Printf("hello Wechat! \n")

	// init wecom app
	configWecom := GetWorkConfig()
	wecomApp, err := work.NewWork(configWecom)
	if err != nil {
		panic(err.Error())
	}
	fmt2.Dump("wecom config:", wecomApp.GetConfig().All())

	// init payment app
	configPayment := GetPaymentConfig()
	paymentApp, err := payment.NewPayment(configPayment)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt2.Dump("payment config:", paymentApp.GetConfig().All())

	// init miniProgram app
	configMiniProgram := GetMiniProgramConfig()
	miniProgramApp, err := miniProgram.NewMiniProgram(configMiniProgram)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt2.Dump("miniprogram config:", miniProgramApp.GetConfig().All())

	// init miniProgram app
	configOpenPlatform := GetOpenPlatformConfig()
	openPlatform, err := openPlatform.NewOpenPlatform(configOpenPlatform)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt2.Dump("openPlatform config:", openPlatform.GetConfig().All())

}
