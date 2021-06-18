package message

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/go-wechat/src/kernel/messages"
)

type Messager struct {
	*messages.Message

	To        *object.StringMap
	AgentID   int
	Secretive bool
	Client    *Client
}
