package contract_test

import (
	"bytes"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/ArtisanCloud/PowerLibs/v3/logger"
	"github.com/ArtisanCloud/PowerLibs/v3/logger/drivers"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/contract"
	kernelModels "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/models"
	server "github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/server"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/server/models"
	"github.com/stretchr/testify/require"
)

type callbackApp struct {
	testApp
	encryptor *kernel.Encryptor
	logger    *logger.Logger
}

func (a callbackApp) GetComponent(name string) interface{} {
	switch name {
	case "Encryptor":
		return a.encryptor
	case "Logger":
		return a.logger
	default:
		panic("unexpected callback component: " + name)
	}
}

type eventHandler func(contract.EventInterface, interface{})

func (h eventHandler) Handle(_ *http.Request, header contract.EventInterface, payload interface{}) interface{} {
	h(header, payload)
	return nil
}

func callbackRequest(t *testing.T, enc *kernel.Encryptor, body string) *http.Request {
	t.Helper()
	envelope, cryptErr := enc.Encrypt(body, "nonce", "123")
	require.Nil(t, cryptErr)
	var reply kernel.WeComReplyMsg
	require.NoError(t, xml.Unmarshal(envelope, &reply))
	query := url.Values{"msg_signature": {reply.Signature.Value}, "nonce": {reply.Nonce.Value}, "timestamp": {reply.Timestamp}}
	return httptest.NewRequest(http.MethodPost, "/callback?"+query.Encode(), bytes.NewReader(envelope))
}

func TestProviderEventDelivery(t *testing.T) {
	enc, err := kernel.NewEncryptor("test-suite", "test-token", base64.StdEncoding.EncodeToString(bytes.Repeat([]byte{1}, 32)))
	require.NoError(t, err)
	log, err := logger.NewLogger(&drivers.DummyLogger{}, nil)
	require.NoError(t, err)
	guard := server.NewGuard(callbackApp{encryptor: enc, logger: log})
	type fixture struct {
		event int
		body  string
	}
	contact := func(id int, change string) fixture {
		return fixture{id, fmt.Sprintf(`<xml><InfoType>change_external_contact</InfoType><ChangeType>%s</ChangeType><AuthCorpId>corp</AuthCorpId><ExternalUserID>external</ExternalUserID></xml>`, change)}
	}
	cases := []fixture{
		{server.EVENT_WEDRIVE_SPACE_CHANGE, `<xml><MsgType>event</MsgType><Event>wedrive_space_change</Event><ChangeType>dismiss_space</ChangeType><SpaceId>s</SpaceId></xml>`},
		contact(server.EVENT_ADD_EXTERNAL_CONTACT, "add_external_contact"),
		contact(server.EVENT_EDIT_EXTERNAL_CONTACT, "edit_external_contact"),
		contact(server.EVENT_DEL_EXTERNAL_CONTACT, "del_external_contact"),
	}
	counts := map[int]int{}
	checkPayload := func(id int, payload interface{}) {
		t.Helper()
		if id == server.EVENT_WEDRIVE_SPACE_CHANGE {
			require.IsType(t, &models.EventWedriveSpaceChange{}, payload)
			require.Equal(t, "s", payload.(*models.EventWedriveSpaceChange).SpaceId)
		} else {
			require.IsType(t, &models.ChangeExternalContact{}, payload)
			ev := payload.(*models.ChangeExternalContact)
			require.Equal(t, "corp", ev.AuthCorpId)
			require.Equal(t, "external", ev.ExternalUserID)
			require.Equal(t, id, server.GetEventType(ev.GetInfoType(), ev.GetChangeType()))
		}
	}
	for _, entry := range cases {
		guard.On(entry.event, eventHandler(func(header contract.EventInterface, payload interface{}) {
			require.NotNil(t, header)
			checkPayload(entry.event, payload)
			counts[entry.event]++
		}))
	}
	for _, entry := range cases {
		res, err := guard.Serve(callbackRequest(t, enc, entry.body))
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, res.StatusCode)
		body, err := io.ReadAll(res.Body)
		require.NoError(t, err)
		res.Body.Close()
		require.Equal(t, "success", string(body))

		calls := 0
		res, err = guard.Notify(callbackRequest(t, enc, entry.body), func(_ *kernelModels.Callback, ev models.IEvent, _ interface{}) interface{} {
			calls++
			checkPayload(entry.event, ev)
			return "success"
		})
		require.NoError(t, err)
		res.Body.Close()
		require.Equal(t, 1, calls)
	}
	for _, entry := range cases {
		require.Equal(t, 1, counts[entry.event])
	}
	invalid := callbackRequest(t, enc, cases[0].body)
	query := invalid.URL.Query()
	query.Set("msg_signature", "invalid")
	invalid.URL.RawQuery = query.Encode()
	_, err = guard.Serve(invalid)
	require.Error(t, err)
	require.Equal(t, 1, counts[cases[0].event])

}
