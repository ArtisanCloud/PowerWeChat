# 企业微信智能机器人 HTTP 回调

在企业自建智能机器人的 API 模式中选择「设置接收消息回调地址」。每个机器人的回调路径注册一个处理器，使用该机器人的 Token 和 EncodingAESKey。企微验证 URL 时发起 GET，用户消息、事件和流式消息刷新通过 POST 到达同一路径。[企微接入说明](https://developer.work.weixin.qq.com/document/path/101039)

下面的处理器对文本消息回复一条已完成的流式消息。返回 `nil, nil` 会确认收到回调。需要读取图片、模板卡片或事件的其他参数时，可从 `msg.Raw` 解码完整的已验签消息。

```go
package bot

import (
	"context"
	"net/http"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/aibot"
)

func Register(mux *http.ServeMux, token, encodingAESKey string) error {
	h, err := aibot.NewCallbackHandler(token, encodingAESKey, func(ctx context.Context, msg *aibot.CallbackMessage) (any, error) {
		if msg.MsgType != "text" {
			return nil, nil
		}
		return map[string]any{
			"msgtype": "stream",
			"stream": map[string]any{
				"id": msg.MsgID, "finish": true, "content": "收到：" + msg.Text.Content,
			},
		}, nil
	})
	if err != nil {
		return err
	}
	mux.Handle("/wecom/bot", h)
	return nil
}
```

将该路径通过 HTTPS 提供给企微。处理器接收最大 10 MiB 的 JSON 信封，验证签名、解密消息并加密被动回复。业务处理返回错误时，处理器响应 HTTP 500；收到重复消息时，可按 `msg.MsgID` 对业务操作去重。持续生成内容时，首次回复设置 `finish: false`，在后续 `msgtype: stream` 的刷新回调中按 `msg.Stream.ID` 查找进度，最后以同一 ID 返回 `finish: true`。[加解密协议](https://developer.work.weixin.qq.com/document/path/101033) · [被动回复消息](https://developer.work.weixin.qq.com/document/path/101031)
