# 移动应用微信登录与分享

在微信开放平台注册移动应用并开通相应能力。iOS 和 Android 客户端集成微信 OpenSDK，并按平台要求配置应用标识及回调。Go 后端使用同一移动应用的 AppID 和 AppSecret，AppSecret 只保存在服务端。

## 微信登录

1. App 为本次登录生成随机 `state`，通过 OpenSDK 发起授权，`scope` 使用 `snsapi_userinfo`。iOS 使用 `SendAuthReq`，Android 使用 `SendAuth.Req`。
2. App 收到微信回调后，确认授权成功且 `state` 与本次请求一致，再通过 HTTPS 将一次性 `code` 发给自己的 Go 后端。
3. 后端调用 `ExchangeCode` 获取 `OpenID`，以移动应用的 AppID 和该 OpenID 关联业务用户，并向 App 签发自己的登录会话。需要展示用户资料时，再调用 `GetUserInfo(ctx, token.AccessToken, token.OpenID)`。微信的 access token 和 refresh token 保存在后端。

下面的函数从一次请求中取得微信用户标识。服务启动时创建一次 `client`，在登录请求中传入 `r.Context()` 和客户端提交的 `code`。

```go
package login

import (
	"context"
	"os"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/mobileApp"
)

func NewClient() (*mobileApp.Client, error) {
	return mobileApp.NewMobileApp(&mobileApp.UserConfig{
		AppID:  os.Getenv("WECHAT_MOBILE_APP_ID"),
		Secret: os.Getenv("WECHAT_MOBILE_APP_SECRET"),
	})
}

func OpenID(ctx context.Context, client *mobileApp.Client, code string) (string, error) {
	token, err := client.ExchangeCode(ctx, code)
	if err != nil {
		return "", err
	}
	return token.OpenID, nil
}
```

授权码已使用或过期时，App 需重新向微信发起授权。[微信登录开发指南](https://developers.weixin.qq.com/doc/oplatform/Mobile_App/WeChat_Login/Development_Guide.html) · [授权后接口调用](https://developers.weixin.qq.com/doc/oplatform/Mobile_App/WeChat_Login/Authorized_API_call_UnionID.html)

## 分享到好友或朋友圈

App 使用 OpenSDK 构造文本、图片或网页消息，设置分享目标，再调起微信，由用户确认分享。根据接入的 OpenSDK 版本和应用配置，需要消息签名时，由服务端按微信的规则签名并将签名交给 App。[OpenSDK 分享签名手册](https://developers.weixin.qq.com/doc/oplatform/Mobile_App/Share_and_Favorites/OpenSDK_Signature.html)

| 平台 | 分享请求 | 好友会话 | 朋友圈 |
| --- | --- | --- | --- |
| iOS | `SendMessageToWXReq` | `WXSceneSession` | `WXSceneTimeline` |
| Android | `SendMessageToWX.Req` | `SendMessageToWX.Req.WXSceneSession` | `SendMessageToWX.Req.WXSceneTimeline` |

各类消息的构造和发送示例见微信官方手册。[iOS 分享示例](https://developers.weixin.qq.com/doc/oplatform/Mobile_App/Share_and_Favorites/iOS.html) · [Android 分享示例](https://developers.weixin.qq.com/doc/oplatform/Mobile_App/Share_and_Favorites/Android.html)
