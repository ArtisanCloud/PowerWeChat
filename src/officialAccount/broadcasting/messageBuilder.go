package broadcasting

import (
	"github.com/ArtisanCloud/PowerLibs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/contract"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
)

type MessageBuilder struct {
	To *power.HashMap

	Message contract.MessageInterface

	Attribute *object.Attribute
}

func NewMessengerBuilder() *MessageBuilder {
	return &MessageBuilder{}
}

func (comp *MessageBuilder) SetMessage(message contract.MessageInterface) *MessageBuilder {
	comp.Message = message

	return comp
}

func (comp *MessageBuilder) SetTo(to *power.HashMap) *MessageBuilder {
	comp.To = to
	return comp
}

func (comp *MessageBuilder) ToTag(tagID int) {
	//comp.SetTo(
	//	)
}
