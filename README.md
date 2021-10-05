# PowerWeChat SDK 介绍

[![Go Build](https://github.com/ArtisanCloud/power-wechat/actions/workflows/go-build.yml/badge.svg?branch=release%2F1.0.0)](https://github.com/ArtisanCloud/power-wechat/actions/workflows/go-build.yml)

[![Go Test](https://github.com/ArtisanCloud/power-wechat/actions/workflows/go-test.yml/badge.svg?branch=release%2F1.0.0)](https://github.com/ArtisanCloud/power-wechat/actions/workflows/go-test.yml)


## 产品介绍


PowerWechat是一款全覆盖微信开发接口，基于Golang的开源项目。您只需安装一次Power WeChat SDK，就可以对接企业微信，小程序，公众号，支付等，微信的开发功能接口。同时我们提供了丰富的文档教程和辅助工具，帮助您轻松使用微信的接口功能。


---

## 产品概述
![Image of Main Page]
(blob:null/ff591363-1799-4516-9f8b-8bb81303fe4f)


### 核心产品
[PowerWeChat SDK](https://github.com/ArtisanCloud/power-wechat) ：是核心的SDK产品，安装后即可开箱即用。
在github上，长期维护的开源项目，可以提Issue在讨论版块。也可以在ArtisanCloud官网上，扫企业微信讨论群，方便用户提问，给宝贵的意见。


#### 简易安装，开箱即用
下载安装
```go
go get github.com/ArtisanCloud/power-wechat
```


示范：初始化实例对象，调用小程序的授权登陆接口
```go
import (
  "github.com/ArtisanCloud/power-wechat/src/miniProgram"
  "os"
)

// 1. 初始化小程序应用实例
app, err := miniProgram.NewMiniProgram(&miniProgram.UserConfig{
    AppID:  os.Getenv("miniprogram_app_id"), // 小程序、公众号或者企业微信的appid
    Secret: os.Getenv("miniprogram_secret"), // 商户号 appID

    ResponseType: os.Getenv("array"),
    HttpDebug: true,
    Debug:     false,
  })


// 2. 调用小程序的授权登陆接口
var code string = "CODE" // 前端小程序登录时，从微信获取的code
rs, err := services.MiniProgramApp.Auth.Session(code)


// 查看获取强类型对象的属性
// 请参考官方文档的返回值
printf(rs.OpenID)
printf(rs.SessionKey)
printf(rs.UnionID)

```

更多实例接口，请打[开官方文档](https://powerwechat.artisan-cloud.com/zh/start/)

---

### 辅助产品

[PowerWeChat Document](https://powerwechat.artisan-cloud.com/zh/start/) ：全面的接口文档，方便用户查找，使用我们开发的sdk功能

![Image of Document Page]
(blob:null/ff591363-1799-4516-9f8b-8bb81303fe4f)



[PowerWeChat Tutorial](https://github.com/ArtisanCloud/power-wechat-tutorial) ：独立的golang项目，提供完整的web接口，让开发者方便调试PowerWeChat 接口实例


PowerWeChat 配置中心客户端/SAAS：如果您有多个微信的开发环境，或者多个应用，可以使用这个配置中心来方便切换账号（此应用暂时内部使用，如需体验，可以联系我们）

![Image of Products]
(blob:null/ff591363-1799-4516-9f8b-8bb81303fe4f)


---


# 产品诞生背景
[ArtisanCloud](https://github.com/ArtisanCloud) 团队也是很多同学一样，从PHP转向Golang，具体为什么，有什么好处，就不用我这里多介绍了吧。 但是现在因为微信的生态做私域化管理是得天独厚，所以我们公司也开发了蛮多企业微信的功能。只是在转型golang的过程中，没有找到像 [Easy Wechat](https://www.easywechat.com)（PHP语言）这样好用的sdk。所以我们就自己想为golang的同学们做一点贡献。产品会长期维护，迭代，希望同学们有兴趣在使用的过程中，多给意见。


---

# 产品特性

## 简易上手，安装一次，全覆盖微信功能接口
## 开源项目，丰富的文档内容，长期维护
## 新增群机器接口和文档
## Golang特性，强类型覆盖
## 完整的测试项目，支持web API测试

---

# 相关资源

## 阅览教程文档
### [官网介绍](https://powerwechat.artisan-cloud.com)
### [使用手册](https://powerwechat.artisan-cloud.com/zh/start/)
### (企业微信群，即将提供)


## Github开源代码
### [SDK源代码](https://github.com/ArtisanCloud/power-wechat)
### [SDK调试项目](https://github.com/ArtisanCloud/power-wechat-tutorial)


## 视频教程
### (策划制作中...)