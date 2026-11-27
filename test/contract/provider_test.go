package contract_test

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"github.com/ArtisanCloud/PowerLibs/v3/http/contract"
	"github.com/ArtisanCloud/PowerLibs/v3/http/helper"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	suite "github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/suitAuth"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/suitAuth/request"
	auth "github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/suitAuth/response"
	users "github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/user"
	templates "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/messageTemplate"
	message "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/messageTemplate/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/oa/wedrive"
	drive "github.com/ArtisanCloud/PowerWeChat/v3/src/work/oa/wedrive/request"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testApp struct{ kernel.ApplicationInterface }

func (testApp) GetConfig() *kernel.Config { return kernel.NewConfig(&object.HashMap{}) }

func post(t *testing.T, path, request, reply string) *kernel.BaseClient {
	t.Helper()
	var calls atomic.Int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls.Add(1)
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/cgi-bin/"+path, r.URL.Path)
		body, err := io.ReadAll(r.Body)
		assert.NoError(t, err)
		assert.JSONEq(t, request, string(body))
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, reply)
	}))
	t.Cleanup(func() { server.Close(); assert.EqualValues(t, 1, calls.Load()) })
	h, err := helper.NewRequestHelper(&helper.Config{BaseUrl: server.URL + "/", ClientConfig: &contract.ClientConfig{}})
	require.NoError(t, err)
	return &kernel.BaseClient{HttpHelper: h, App: testApp{}}
}

func TestSetSessionInfo(t *testing.T) {
	const path = "service/set_session_info"
	const reply = `{"errcode":0,"errmsg":"ok"}`

	t.Run("default authorization", func(t *testing.T) {
		client := suite.Client{BaseClient: post(t, path, `{"pre_auth_code":"code","session_info":{}}`, reply)}
		err := client.SetSessionInfo(context.Background(), &request.SetSessionInfoRequest{PreAuthCode: "code"})
		require.NoError(t, err)
	})

	t.Run("test authorization for selected apps", func(t *testing.T) {
		client := suite.Client{BaseClient: post(t, path, `{"pre_auth_code":"code","session_info":{"appid":[1,2,3],"auth_type":1}}`, reply)}
		err := client.SetSessionInfo(context.Background(), &request.SetSessionInfoRequest{
			PreAuthCode: "code",
			SessionInfo: request.SessionInfo{AppID: []uint64{1, 2, 3}, AuthType: 1},
		})
		require.NoError(t, err)
	})
}

func TestProviderResponses(t *testing.T) {
	ctx := context.Background()
	const legacy = `{"permanent_code":"p","access_token":"token","expires_in":7200,"auth_info":{"agent":[{"agentid":1}]},"dealer_corp_info":{"corpid":"dealer"}}`
	client := suite.Client{BaseClient: post(t, "service/get_permanent_code", `{"auth_code":"code"}`, legacy)}
	result, err := client.GetPermanentCode(ctx, "code")
	require.NoError(t, err)
	require.Equal(t, "token", result.AccessToken)
	require.EqualValues(t, 7200, result.ExpiresIn)
	require.Equal(t, "dealer", result.DealerCorpInfo.CorpId)
	require.Equal(t, 1, result.AuthInfo.Agent[0].AgentID)

	client.BaseClient = post(t, "service/get_auth_info", `{"auth_corpid":"corp","permanent_code":"p"}`, legacy)
	result, err = client.GetAuthInfo(ctx, "corp", "p")
	require.NoError(t, err)
	require.Equal(t, 1, result.AuthInfo.Agent[0].AgentID)

	client.BaseClient = post(t, "service/v2/get_auth_info", `{"auth_corpid":"corp","permanent_code":"p"}`, `{"auth_info":{"agent":[{"auth_from_thirdapp":true,"shared_from":{"corpid":"source"}}]},"auth_corp_info":{"corp_ex_name":{"name_list":["Example"]}}}`)
	var existingSignature func(context.Context, string, string) (*auth.GetPermanentCodeResponse, error) = client.GetAuthInfoV2
	result, err = existingSignature(ctx, "corp", "p")
	require.NoError(t, err)
	require.True(t, result.AuthInfo.Agent[0].AuthFromThirdApp)
	require.Equal(t, "source", result.AuthInfo.Agent[0].SharedFrom.CorpID)
	require.Equal(t, []string{"Example"}, result.AuthCorpInfo.CorpExName.NameList)

	for _, reply := range []string{`{"gender":"1","qr_code":"qr"}`, `{"gender":1,"qrcode":"qr"}`} {
		userClient := users.Client{BaseClient: post(t, "service/auth/getuserdetail3rd", `{"user_ticket":"ticket"}`, reply)}
		member, err := userClient.GetUserInfo3rdByUserTicket(ctx, "ticket")
		require.NoError(t, err)
		require.Equal(t, 1, member.Gender)
		require.Equal(t, "qr", member.Qrcode)
		require.Empty(t, member.Mobile)
		require.Empty(t, member.Email)
	}
}

func TestProviderMessageAttachments(t *testing.T) {
	const items = `[{"msgtype":"image","image":{"media_id":"a","pic_url":""}},{"msgtype":"image","image":{"media_id":"b","pic_url":""}}]`
	images := []message.MessageTemplateInterface{
		&message.ImageOfMessage{MsgType: "image", Image: &message.Image{MediaID: "a"}},
		&message.Attachment{MsgType: "image", Image: &message.Image{MediaID: "b"}},
	}
	c := templates.Client{BaseClient: post(t, "externalcontact/add_msg_template", `{"chat_type":"single","external_userid":null,"chat_id_list":null,"tag_filter":{"group_list":null},"allow_select":false,"attachments":`+items+`}`, "{}")}
	_, err := c.AddMsgTemplate(context.Background(), &message.RequestAddMsgTemplate{ChatType: "single", Attachments: images})
	require.NoError(t, err)
	c.BaseClient = post(t, "externalcontact/send_welcome_msg", `{"welcome_code":"w","text":null,"attachments":`+items+`}`, "{}")
	_, err = c.SendWelcomeMsg(context.Background(), &message.RequestSendWelcomeMsg{WelcomeCode: "w", Attachments: images})
	require.NoError(t, err)
}

func TestWeDriveLegacyAndPartialUpdates(t *testing.T) {
	ctx := context.Background()
	c := wedrive.Client{BaseClient: post(t, "wedrive/file_delete", `{"userid":"u","fileid":["f"]}`, "{}")}
	_, err := c.FileDelete(ctx, &drive.RequestWeDriveFileDelete{UserID: "u", FileID: "f"})
	require.NoError(t, err)
	c.BaseClient = post(t, "wedrive/file_delete", `{"fileid":["a","b"]}`, "{}")
	_, err = c.FileDeleteBatch(ctx, &drive.RequestWeDriveFileDeleteBatch{FileIDs: []string{"a", "b"}})
	require.NoError(t, err)
	c.BaseClient = post(t, "wedrive/space_setting", `{"userid":"u","spaceid":"s","enable_watermark":false,"add_member_only_admin":false,"enable_share_url":false,"share_url_no_approve":false,"share_url_no_approve_default_auth":0}`, "{}")
	_, err = c.SpaceSetting(ctx, &drive.RequestWeDriveSpaceSetting{UserID: "u", SpaceID: "s"})
	require.NoError(t, err)
	off, on := false, true
	c.BaseClient = post(t, "wedrive/space_setting", `{"spaceid":"s","share_url_no_approve":false,"enable_watermark":true}`, "{}")
	_, err = c.SpaceSettingPatch(ctx, &drive.RequestWeDriveSpaceSettingPatch{SpaceID: "s", ShareUrlNoApprove: &off, EnableWatermark: &on})
	require.NoError(t, err)
}

func TestWeDriveUpload(t *testing.T) {
	ctx := context.Background()
	c := wedrive.Client{BaseClient: post(t, "wedrive/file_upload", `{"userid":"u","spaceid":"s","fatherid":"f","file_name":"x","file_base64_content":"QQ=="}`, "{}")}
	upload := drive.RequestWeDriveFileUpload{UserID: "u", SpaceID: "s", FatherID: "f", FileName: "x", FileBase64Content: "QQ=="}
	_, err := c.FileUpload(ctx, &upload)
	require.NoError(t, err)
	c.BaseClient = post(t, "wedrive/file_upload", `{"selected_ticket":"t","file_name":"x","file_base64_content":"QQ=="}`, "{}")
	upload.SelectedTicket = "t"
	_, err = c.FileUpload(ctx, &upload)
	require.NoError(t, err)
	require.Equal(t, "s", upload.SpaceID)
	c.BaseClient = post(t, "wedrive/file_download", `{"selected_ticket":"t"}`, "{}")
	_, err = c.FileDownload(ctx, &drive.RequestWeDriveFileDownload{UserID: "u", FileID: "f", SelectedTicket: "t"})
	require.NoError(t, err)

	c.BaseClient = post(t, "wedrive/file_upload_init", `{"spaceid":"s","fatherid":"f","file_name":"x","size":1,"block_sha":["sha"]}`, `{"hit_exist":false,"upload_key":"k"}`)
	init := drive.RequestWeDriveFileUploadInit{SpaceID: "s", FatherID: "f", FileName: "x", Size: 1, BlockSHA: []string{"sha"}}
	started, err := c.FileUploadInit(ctx, &init)
	require.NoError(t, err)
	require.False(t, started.HitExist)
	c.BaseClient = post(t, "wedrive/file_upload_part", `{"upload_key":"k","index":1,"file_base64_content":"QQ=="}`, "{}")
	_, err = c.FileUploadPart(ctx, &drive.RequestWeDriveFileUploadPart{UploadKey: started.UploadKey, Index: 1, FileBase64Content: "QQ=="})
	require.NoError(t, err)
	c.BaseClient = post(t, "wedrive/file_upload_finish", `{"upload_key":"k"}`, `{"fileid":"done"}`)
	finished, err := c.FileUploadFinish(ctx, &drive.RequestWeDriveFileUploadFinish{UploadKey: started.UploadKey})
	require.NoError(t, err)
	require.Equal(t, "done", finished.Fileid)
	init.SelectedTicket = "t"
	c.BaseClient = post(t, "wedrive/file_upload_init", `{"selected_ticket":"t","file_name":"x","size":1,"block_sha":["sha"]}`, `{"hit_exist":true,"fileid":"already"}`)
	hit, err := c.FileUploadInit(ctx, &init)
	require.NoError(t, err)
	require.True(t, hit.HitExist)
	require.Equal(t, "already", hit.FileID)
}
