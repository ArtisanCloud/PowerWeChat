package main

import (
	"context"
	"fmt"
	fmt2 "github.com/ArtisanCloud/PowerLibs/v3/fmt"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/officialAccount/templateMessage/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openPlatform"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/trace"
	"os"
	"strconv"
)

func GetOfficialConfig() *officialAccount.UserConfig {
	return &officialAccount.UserConfig{
		AppID:  "", // 小程序、公众号或者企业微信的appid
		Secret: "", // 商户号 appID

		Token:  "",
		AESKey: "",
		Log: officialAccount.Log{
			Level:  "debug",
			Stdout: true,
		},

		// ResponseType: os.Getenv("response_type"),
		Cache: kernel.NewRedisClient(&kernel.UniversalOptions{
			Addrs:    []string{"127.0.0.1:6379"},
			Password: "",
			DB:       1,
		}),
		HttpDebug: false,
		Debug:     false,
	}

}

func GetWorkConfig() *work.UserConfig {
	agentID, _ := strconv.Atoi(os.Getenv("wecom_agent_id"))
	return &work.UserConfig{
		CorpID:  os.Getenv("corp_id"),
		AgentID: agentID,
		Secret:  os.Getenv("wecom_secret"),

		ResponseType: os.Getenv("array"),
		Log: work.Log{
			Level: "debug",
			File:  "./wechat/info.log",
			Error: "./wechat/error.log",
			ENV:   os.Getenv("work.env"),
		},

		OAuth: work.OAuth{
			Callback: os.Getenv("app_oauth_callback_url"),
			Scopes:   []string{},
		},
		Cache: kernel.NewRedisClient(&kernel.UniversalOptions{
			Addrs:    []string{"127.0.0.1:6379"},
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
			File:  "./wechat/info.log",
			Error: "./wechat/error.log",
		},
		Http: payment.Http{
			Timeout: 30.0,
			BaseURI: "https://api.mch.weixin.qq.com",
		},

		Cache: kernel.NewRedisClient(&kernel.UniversalOptions{
			Addrs:    []string{"127.0.0.1:6379"},
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
			File:  "./wechat/info.log",
			Error: "./wechat/error.log",
		},
		Cache: kernel.NewRedisClient(&kernel.UniversalOptions{
			Addrs:    []string{"127.0.0.1:6379"},
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
			File:  "./wechat/info.log",
			Error: "./wechat/error.log",
		},
		Cache: kernel.NewRedisClient(&kernel.UniversalOptions{
			Addrs:    []string{"127.0.0.1:6379"},
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

func initTracer() {
	tp := trace.NewTracerProvider()
	// Set Global Tracer Provider
	otel.SetTracerProvider(tp)
}

func init() {
	initTracer()
}

func main() {

	fmt.Printf("hello Wechat! \n")

	tracer := otel.Tracer("example-tracer")
	ctx, span := tracer.Start(context.Background(), "test")
	defer span.End()

	// init officialAccount app
	configOfficialAccount := GetOfficialConfig()
	officialAccountApp, err := officialAccount.NewOfficialAccount(configOfficialAccount)
	if err != nil {
		fmt.Println(err.Error())
	}
	//officialAccountApp.Logger.Driver.Info("custom info log")
	//officialAccountApp.Logger.Driver.Error("custom error log")
	//officialAccountApp.Logger.Driver.Warn("custom warn log")

	officialAccountApp.TemplateMessage.Send(ctx, &request.RequestTemlateMessage{
		ToUser:     "",
		TemplateID: "",
		MiniProgram: &request.MiniProgram{ // 上线后的小程序才可以使用
			AppID:    "",
			PagePath: fmt.Sprintf("pages/order/details?id=%v", 111111),
		},
		Data: &power.HashMap{
			"character_string1": &power.HashMap{
				"value": "111111111",
				"color": "#173177",
			},
			"thing3": &power.HashMap{
				"value": "可口可乐大瓶装",
				"color": "#173177",
			},
			"const2": &power.HashMap{
				"value": "订单超时未支付",
				"color": "#173177",
			},
		},
	})
	fmt2.Dump("officialAccount config:", officialAccountApp.GetConfig().All())
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
