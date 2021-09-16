package work

import (
	"github.com/ArtisanCloud/go-libs/object"
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/kernel/providers"
	"github.com/ArtisanCloud/power-wechat/src/work/accountService"
	"github.com/ArtisanCloud/power-wechat/src/work/accountService/customer"
	message3 "github.com/ArtisanCloud/power-wechat/src/work/accountService/message"
	"github.com/ArtisanCloud/power-wechat/src/work/accountService/serviceState"
	"github.com/ArtisanCloud/power-wechat/src/work/accountService/servicer"
	tag3 "github.com/ArtisanCloud/power-wechat/src/work/accountService/tag"
	"github.com/ArtisanCloud/power-wechat/src/work/agent"
	"github.com/ArtisanCloud/power-wechat/src/work/agent/workbench"
	"github.com/ArtisanCloud/power-wechat/src/work/auth"
	"github.com/ArtisanCloud/power-wechat/src/work/base"
	"github.com/ArtisanCloud/power-wechat/src/work/corpgroup"
	"github.com/ArtisanCloud/power-wechat/src/work/department"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/contactWay"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/customerStrategy"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/groupChat"
	message2 "github.com/ArtisanCloud/power-wechat/src/work/externalContact/message"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/messageTemplate"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/moment"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/school"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/statistics"
	tag2 "github.com/ArtisanCloud/power-wechat/src/work/externalContact/tag"
	"github.com/ArtisanCloud/power-wechat/src/work/externalContact/transfer"
	"github.com/ArtisanCloud/power-wechat/src/work/groupRobot"
	"github.com/ArtisanCloud/power-wechat/src/work/invoice"
	"github.com/ArtisanCloud/power-wechat/src/work/media"
	"github.com/ArtisanCloud/power-wechat/src/work/menu"
	"github.com/ArtisanCloud/power-wechat/src/work/message"
	msgaudit "github.com/ArtisanCloud/power-wechat/src/work/msgAudit"
	"github.com/ArtisanCloud/power-wechat/src/work/oa"
	"github.com/ArtisanCloud/power-wechat/src/work/oauth"
	"github.com/ArtisanCloud/power-wechat/src/work/server"
	"github.com/ArtisanCloud/power-wechat/src/work/user"
	"github.com/ArtisanCloud/power-wechat/src/work/user/batchJobs"
	"github.com/ArtisanCloud/power-wechat/src/work/user/exportJobs"
	"github.com/ArtisanCloud/power-wechat/src/work/user/linkedCorp"
	"github.com/ArtisanCloud/power-wechat/src/work/user/tag"
	"net/http"
)

type Work struct {
	*kernel.ServiceContainer

	ExternalRequest *http.Request

	Base        *base.Client
	AccessToken *auth.AccessToken
	OAuth       *oauth.Manager

	Config     *kernel.Config
	Department *department.Client

	Agent          *agent.Client
	AgentWorkbench *workbench.Client

	Message  *message.Client
	Messager *message.Messager

	Encryptor *kernel.Encryptor
	Server    *server.Guard

	User           *user.Client
	UserBatchJobs  *batchJobs.Client
	UserExportJobs *exportJobs.Client
	UserLinkedCorp *linkedCorp.Client
	UserTag        *tag.Client

	ExternalContact                 *externalContact.Client
	ExternalContactContactWay       *contactWay.Client
	ExternalContactCustomerStrategy *customerStrategy.Client
	ExternalContactStatistics       *statistics.Client
	ExternalContactMessage          *message2.Client
	ExternalContactSchool           *school.Client
	ExternalContactMoment           *moment.Client
	ExternalContactMessageTemplate  *messageTemplate.Client
	ExternalContactGroupChat        *groupChat.Client
	ExternalContactTag              *tag2.Client
	ExternalContactTransfer         *transfer.Client

	AccountService             *accountService.Client
	AccountServiceCustomer     *customer.Client
	AccountServiceMessage      *message3.Client
	AccountServiceServicer     *servicer.Client
	AccountServiceServiceState *serviceState.Client
	AccountServiceTag          *tag3.Client

	Media *media.Client
	Menu  *menu.Client

	OA *oa.Client

	MsgAudit *msgaudit.Client

	CorpGroup *corpgroup.Client

	Invoice *invoice.Client

	GroupRobot          *groupRobot.Client
	GroupRobotMessenger *groupRobot.Messager
}

type UserConfig struct {
	CorpID           string
	AgentID          int
	Secret           string
	Token            string
	AESKey           string
	AuthCallbackHost string

	ResponseType string
	Log          Log
	OAuth        OAuth
	HttpDebug    bool
	Debug        bool
}

type Log struct {
	Level string
	File  string
}

type OAuth struct {
	Callback string
	Scopes   []string
}

func NewWork(config *UserConfig) (*Work, error) {
	var err error

	userConfig, err := MapUserConfig(config)
	if err != nil {
		return nil, err
	}

	// init an app container
	container := &kernel.ServiceContainer{
		UserConfig: userConfig,
		DefaultConfig: &object.HashMap{
			"http": object.HashMap{
				"base_uri": "https://qyapi.weixin.qq.com/",
			},
		},
	}
	container.GetConfig()

	// init app
	app := &Work{
		ServiceContainer: container,
	}

	//-------------- global app config --------------
	// global app config
	app.Config = providers.RegisterConfigProvider(app)

	//-------------- register Auth --------------
	app.AccessToken = auth.RegisterProvider(app)
	//-------------- register Base --------------
	app.Base = base.RegisterProvider(app)

	//-------------- register oauth --------------
	app.OAuth, err = oauth.RegisterProvider(app)

	//-------------- register agent --------------
	app.Agent,
		app.AgentWorkbench = agent.RegisterProvider(app)

	//-------------- register Department --------------
	app.Department = department.RegisterProvider(app)

	//-------------- register Message --------------
	app.Message, app.Messager = message.RegisterProvider(app)

	//-------------- register Encryptor --------------
	app.Encryptor, app.Server = server.RegisterProvider(app)

	//-------------- register user --------------
	app.User,
		app.UserBatchJobs,
		app.UserExportJobs,
		app.UserLinkedCorp,
		app.UserTag = user.RegisterProvider(app)

	//-------------- register external contact --------------
	app.ExternalContact,
		app.ExternalContactContactWay,
		app.ExternalContactCustomerStrategy,
		app.ExternalContactGroupChat,
		app.ExternalContactMessage,
		app.ExternalContactMessageTemplate,
		app.ExternalContactMoment,
		app.ExternalContactSchool,
		app.ExternalContactStatistics,
		app.ExternalContactTag,
		app.ExternalContactTransfer = externalContact.RegisterProvider(app)

	//-------------- register account service --------------
	app.AccountService,
		app.AccountServiceCustomer,
		app.AccountServiceMessage,
		app.AccountServiceServicer,
		app.AccountServiceServiceState,
		app.AccountServiceTag = accountService.RegisterProvider(app)

	//-------------- media --------------
	app.Media = media.RegisterProvider(app)

	//-------------- menu --------------
	app.Menu = menu.RegisterProvider(app)

	//-------------- oa --------------
	app.OA = oa.RegisterProvider(app)

	//-------------- msg audit --------------
	app.MsgAudit = msgaudit.RegisterProvider(app)

	//-------------- corp group --------------
	app.CorpGroup = corpgroup.RegisterProvider(app)

	//-------------- invoice --------------
	app.Invoice = invoice.RegisterProvider(app)

	app.GroupRobot, app.GroupRobotMessenger = groupRobot.RegisterProvider(app)

	return app, err
}

func (app *Work) GetContainer() *kernel.ServiceContainer {
	return app.ServiceContainer
}

func (app *Work) GetAccessToken() *kernel.AccessToken {
	return app.AccessToken.AccessToken
}

func (app *Work) GetConfig() *kernel.Config {
	return app.Config
}

func (app *Work) GetComponent(name string) interface{} {

	switch name {
	case "Base":
		return app.Base
	case "AccessToken":
		return app.AccessToken
	case "OAuth":
		return app.OAuth
	case "Config":
		return app.Config
	case "Department":
		return app.Department

	case "Message":
		return app.Message
	case "Messager":
		return app.Messager

	case "Encryptor":
		return app.Encryptor
	case "Server":
		return app.Server

	case "UserClient":
		return app.User
	case "UserBatchJobsClient":
		return app.UserBatchJobs
	case "UserExportJobs":
		return app.UserExportJobs
	case "UserLinkedCorpClient":
		return app.UserLinkedCorp
	case "UserTagClient":
		return app.UserTag

	case "ExternalContact":
		return app.ExternalContact
	case "ExternalContactContactWay":
		return app.ExternalContactContactWay
	case "ExternalContactStatistics":
		return app.ExternalContactStatistics
	case "ExternalContactMessage":
		return app.ExternalContactMessage
	case "ExternalContactSchool":
		return app.ExternalContactSchool
	case "ExternalContactMoment":
		return app.ExternalContactMoment
	case "ExternalContactMessageTemplate":
		return app.ExternalContactMessageTemplate

	case "Menu":
		return app.Menu
	case "OA":
		return app.OA
	case "MsgAudit":
		return app.MsgAudit
	case "CorpGroup":
		return app.CorpGroup
	case "Invoice":
		return app.Invoice

	case "GroupRobot":
		return app.GroupRobot
	case "GroupRobotMessenger":
		return app.GroupRobotMessenger

	default:
		return nil
	}

}

func MapUserConfig(userConfig *UserConfig) (*object.HashMap, error) {

	config := &object.HashMap{
		"corp_id":            userConfig.CorpID,
		"agent_id":           userConfig.AgentID,
		"secret":             userConfig.Secret,
		"token":              userConfig.Token,
		"aes_key":            userConfig.AESKey,
		"auth_callback_host": userConfig.AuthCallbackHost,

		"response_type": userConfig.ResponseType,
		"log": object.StringMap{
			"level": userConfig.Log.Level,
			"file":  userConfig.Log.File,
		},
		"oauth.callback": userConfig.OAuth.Callback,
		"oauth.scopes":   userConfig.OAuth.Scopes,
		"http_debug":     userConfig.HttpDebug,
		"debug":          userConfig.Debug,
	}

	return config, nil

}
