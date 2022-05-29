package broadcasting

import (
	"errors"
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/contract"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/officialAccount/broadcasting/request"
)

type MessageBuilder struct {
	To *object.HashMap

	Message contract.MessageInterface

	Attribute *object.Attribute
}

func NewMessengerBuilder() *MessageBuilder {
	return &MessageBuilder{
		Attribute: object.NewAttribute(&object.HashMap{}),
	}
}

func (comp *MessageBuilder) SetMessage(message contract.MessageInterface) *MessageBuilder {
	comp.Message = message

	return comp
}

func (comp *MessageBuilder) SetTo(to *power.HashMap) *MessageBuilder {

	objTo, _ := power.PowerHashMapToObjectHashMap(to)

	comp.To = objTo

	return comp
}

func (comp *MessageBuilder) ToTag(tagID int) *MessageBuilder {
	comp.SetTo(&power.HashMap{
		"filter": &power.HashMap{
			"is_to_all": false,
			"tag_id":    tagID,
		},
	})
	return comp
}

func (comp *MessageBuilder) ToUsers(openIDs []string) *MessageBuilder {
	comp.SetTo(&power.HashMap{
		"touser": openIDs,
	})
	return comp
}

func (comp *MessageBuilder) ToAll() *MessageBuilder {
	comp.SetTo(&power.HashMap{
		"filter": &power.HashMap{
			"is_to_all": true,
		},
	})
	return comp
}

func (comp *MessageBuilder) With(attributes *power.HashMap) (*MessageBuilder, error) {

	objAttributes, _ := power.PowerHashMapToObjectHashMap(attributes)

	comp.Attribute.SetAttributes(objAttributes)
	return comp, nil
}

func (comp *MessageBuilder) Build(prepends *object.HashMap) (*object.HashMap, error) {

	if comp.Message == nil {
		return nil, errors.New("no message content to send")
	}

	content, err := comp.Message.TransformForJsonRequest(&object.HashMap{}, true)
	if err != nil {
		return nil, err
	}

	if prepends == nil {
		prepends = comp.To

	}

	message := object.MergeHashMap(prepends, content, comp.Attribute.GetAttributes())

	return message, nil
}

func (comp *MessageBuilder) BuildForPreview(by string, user *request.Reception) (*object.HashMap, error) {
	return comp.Build(&object.HashMap{
		by: user,
	})
}
