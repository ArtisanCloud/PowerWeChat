package messages

import (
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/contract"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
)

const VOID = 0
const TEXT = 2
const IMAGE = 4
const VOICE = 8
const VIDEO = 16
const SHORT_VIDEO = 32
const LOCATION = 64
const LINK = 128
const DEVICE_EVENT = 256
const DEVICE_TEXT = 512
const FILE = 1024
const TEXT_CARD = 2048
const TRANSFER = 4096
const EVENT = 1048576
const MINIPROGRAM_PAGE = 2097152
const MINIPROGRAM_NOTICE = 4194304

const ALL = TEXT | IMAGE | VOICE | VIDEO |
	SHORT_VIDEO | LOCATION | LINK | DEVICE_EVENT |
	DEVICE_TEXT | FILE | TEXT_CARD | TRANSFER |
	EVENT | MINIPROGRAM_PAGE | MINIPROGRAM_NOTICE

type Message struct {
	contract.MessageInterface

	*object.Attribute

	Type        string
	ID          int
	To          string
	From        string
	Properties  []string
	JsonAliases *object.HashMap

	ToXmlArray func() *object.HashMap
}

func NewMessage(attributes *power.HashMap) *Message {
	objAttribute, _ := power.PowerHashMapToObjectHashMap(attributes)

	m := &Message{
		Attribute: &object.Attribute{},
	}
	m.SetAttributes(objAttribute)
	m.ToXmlArray = func() *object.HashMap {
		return m.toXmlArray()
	}

	return m
}

func (msg *Message) GetType() string {
	return msg.Type
}

func (msg *Message) SetType(strType string) {
	msg.Type = strType
}

func (msg *Message) TransformForJsonRequestWithoutType(appends *object.HashMap) (*object.HashMap, error) {
	return msg.TransformForJsonRequest(appends, false)
}

func (msg *Message) TransformForJsonRequest(appends *object.HashMap, withType bool) (*object.HashMap, error) {
	if !withType {
		return msg.PropertiesToArray(&object.HashMap{}, msg.JsonAliases)
	}

	messageType := msg.GetType()
	data := object.MergeHashMap(&object.HashMap{"msgtype": messageType}, appends)
	arrayType := &object.HashMap{}
	if (*data)[messageType] != nil {
		arrayType = (*data)[messageType].(*object.HashMap)
	}
	arrayFromProperties, err := msg.PropertiesToArray(&object.HashMap{}, msg.JsonAliases)
	if err != nil {
		return nil, err
	}
	(*data)[messageType] = object.MergeHashMap(arrayType, arrayFromProperties)

	return data, nil
}

// return string or hashmap
func (msg *Message) TransformToXml(appends *object.HashMap, returnAsArray bool) (interface{}, error) {
	data := object.MergeHashMap(&object.HashMap{"MsgType": msg.GetType()}, msg.ToXmlArray(), appends)

	if returnAsArray {
		return data, nil
	} else {
		strXML := object.Map2Xml(data, false)
		return strXML, nil
		//buffer, err := xml2.Marshal(data)
		//if err != nil {
		//	return "", err
		//}
		//return string(buffer), nil
	}
}

func (msg *Message) PropertiesToArray(data *object.HashMap, aliases *object.HashMap) (*object.HashMap, error) {
	err := msg.CheckRequiredAttributes()
	if err != nil {
		return nil, err
	}

	for property, value := range msg.Attributes {
		if value == nil && !msg.IsRequired(property) {
			continue
		}
		has, alias := object.InHash(property, aliases)
		if has {
			(*data)[alias] = msg.Get(property, nil)
		} else {
			(*data)[property] = msg.Get(property, nil)
		}
	}

	return data, nil
}

func (msg *Message) toXmlArray() *object.HashMap {
	return &object.HashMap{}
}
