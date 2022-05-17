package work

import (
	"github.com/ArtisanCloud/PowerLibs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/providers"
	"github.com/ArtisanCloud/PowerWeChat/src/work/accountService"
	"github.com/ArtisanCloud/PowerWeChat/src/work/accountService/customer"
	message3 "github.com/ArtisanCloud/PowerWeChat/src/work/accountService/message"
	"github.com/ArtisanCloud/PowerWeChat/src/work/accountService/serviceState"
	"github.com/ArtisanCloud/PowerWeChat/src/work/accountService/servicer"
	tag3 "github.com/ArtisanCloud/PowerWeChat/src/work/accountService/tag"
	"github.com/ArtisanCloud/PowerWeChat/src/work/agent"
	"github.com/ArtisanCloud/PowerWeChat/src/work/agent/workbench"
	"github.com/ArtisanCloud/PowerWeChat/src/work/auth"
	"github.com/ArtisanCloud/PowerWeChat/src/work/base"
	"github.com/ArtisanCloud/PowerWeChat/src/work/corpgroup"
	"github.com/ArtisanCloud/PowerWeChat/src/work/department"
	"github.com/ArtisanCloud/PowerWeChat/src/work/externalContact"
	"github.com/ArtisanCloud/PowerWeChat/src/work/externalContact/contactWay"
	"github.com/ArtisanCloud/PowerWeChat/src/work/externalContact/customerStrategy"
	"github.com/ArtisanCloud/PowerWeChat/src/work/externalContact/groupChat"
	"github.com/ArtisanCloud/PowerWeChat/src/work/externalContact/groupWelcomeTemplate"
	"github.com/ArtisanCloud/PowerWeChat/src/work/externalContact/messageTemplate"
	"github.com/ArtisanCloud/PowerWeChat/src/work/externalContact/moment"
	"github.com/ArtisanCloud/PowerWeChat/src/work/externalContact/momentStrategy"
	"github.com/ArtisanCloud/PowerWeChat/src/work/externalContact/school"
	"github.com/ArtisanCloud/PowerWeChat/src/work/externalContact/statistics"
	tag2 "github.com/ArtisanCloud/PowerWeChat/src/work/externalContact/tag"
	"github.com/ArtisanCloud/PowerWeChat/src/work/externalContact/transfer"
	"github.com/ArtisanCloud/PowerWeChat/src/work/groupRobot"
	"github.com/ArtisanCloud/PowerWeChat/src/work/invoice"
	"github.com/ArtisanCloud/PowerWeChat/src/work/media"
	"github.com/ArtisanCloud/PowerWeChat/src/work/menu"
	"github.com/ArtisanCloud/PowerWeChat/src/work/message"
	"github.com/ArtisanCloud/PowerWeChat/src/work/message/appChat"
	externalContact2 "github.com/ArtisanCloud/PowerWeChat/src/work/message/externalContact"
	linkedCorp2 "github.com/ArtisanCloud/PowerWeChat/src/work/message/linkedCorp"
	msgaudit "github.com/ArtisanCloud/PowerWeChat/src/work/msgAudit"
	"github.com/ArtisanCloud/PowerWeChat/src/work/oa"
	"github.com/ArtisanCloud/PowerWeChat/src/work/oa/calendar"
	"github.com/ArtisanCloud/PowerWeChat/src/work/oa/dial"
	"github.com/ArtisanCloud/PowerWeChat/src/work/oa/journal"
	"github.com/ArtisanCloud/PowerWeChat/src/work/oa/living"
	"github.com/ArtisanCloud/PowerWeChat/src/work/oa/meeting"
	"github.com/ArtisanCloud/PowerWeChat/src/work/oa/meetingroom"
	"github.com/ArtisanCloud/PowerWeChat/src/work/oa/pstncc"
	"github.com/ArtisanCloud/PowerWeChat/src/work/oa/schedule"
	"github.com/ArtisanCloud/PowerWeChat/src/work/oa/webdrive"
	"github.com/ArtisanCloud/PowerWeChat/src/work/oauth"
	"github.com/ArtisanCloud/PowerWeChat/src/work/server"
	"github.com/ArtisanCloud/PowerWeChat/src/work/user"
	"github.com/ArtisanCloud/PowerWeChat/src/work/user/batchJobs"
	"github.com/ArtisanCloud/PowerWeChat/src/work/user/exportJobs"
	"github.com/ArtisanCloud/PowerWeChat/src/work/user/linkedCorp"
	"github.com/ArtisanCloud/PowerWeChat/src/work/user/tag"
	"net/http"
)

type Work struct {
	*kernel.ServiceContainer

	Base        *base.Client
	AccessToken *auth.AccessToken
	OAuth       *oauth.Manager

	Config     *kernel.Config
	Department *department.Client

	Agent          *agent.Client
	AgentWorkbench *workbench.Client

	Message                *message.Client
	Messager               *message.Messager
	MessageAppChat         *appChat.Client
	MessageExternalContact *externalContact2.Client
	MessageLinkedCorp      *linkedCorp2.Client

	Encryptor *kernel.Encryptor
	Server    *server.Guard

	User           *user.Client
	UserBatchJobs  *batchJobs.Client
	UserExportJobs *exportJobs.Client
	UserLinkedCorp *linkedCorp.Client
	UserTag        *tag.Client

	ExternalContact                     *externalContact.Client
	ExternalContactContactWay           *contactWay.Client
	ExternalContactCustomerStrategy     *customerStrategy.Client
	ExternalContactStatistics           *statistics.Client
	ExternalContactGroupWelcomeTemplate *groupWelcomeTemplate.Client
	ExternalContactSchool               *school.Client
	ExternalContactMoment               *moment.Client
	ExternalContactMomentStrategy       *momentStrategy.Client
	ExternalContactMessageTemplate      *messageTemplate.Client
	ExternalContactGroupChat            *groupChat.Client
	ExternalContactTag                  *tag2.Client
	ExternalContactTransfer             *transfer.Client

	AccountService         *accountService.Client
	AccountServiceCustomer *customer.Client
	AccountServiceMessage  *message3.Client
	AccountServiceServicer *servicer.Client
	AccountServiceState    *serviceState.Client
	AccountServiceTag      *tag3.Client

	Media *media.Client
	Menu  *menu.Client

	OA            *oa.Client
	OACalendar    *calendar.Client
	OADial        *dial.Client
	OAJournal     *journal.Client
	OALiving      *living.Client
	OAMeeting     *meeting.Client
	OAMeetingRoom *meetingroom.Client
	OAPSTNCC      *pstncc.Client
	OASchedule    *schedule.Client
	OAWebDrive    *webdrive.Client

	MsgAudit *msgaudit.Client

	CorpGroup *corpgroup.Client

	Invoice *invoice.Client

	GroupRobot          *groupRobot.Client
	GroupRobotMessenger *groupRobot.Messager
}

type UserConfig struct {
	CorpID      string
	AgentID     int
	Secret      string
	Token       string
	AESKey      string
	CallbackURL string

	ResponseType string
	Log          Log
	OAuth        OAuth
	Cache        kernel.CacheInterface
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
			"http": &object.HashMap{
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
	app.Message, app.Messager,
		app.MessageAppChat,
		app.MessageExternalContact,
		app.MessageLinkedCorp = message.RegisterProvider(app)

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
		app.ExternalContactGroupWelcomeTemplate,
		app.ExternalContactMessageTemplate,
		app.ExternalContactMoment,
		app.ExternalContactMomentStrategy,
		app.ExternalContactSchool,
		app.ExternalContactStatistics,
		app.ExternalContactTag,
		app.ExternalContactTransfer = externalContact.RegisterProvider(app)

	//-------------- register account service --------------
	app.AccountService,
		app.AccountServiceCustomer,
		app.AccountServiceMessage,
		app.AccountServiceServicer,
		app.AccountServiceState,
		app.AccountServiceTag = accountService.RegisterProvider(app)

	//-------------- media --------------
	app.Media = media.RegisterProvider(app)

	//-------------- menu --------------
	app.Menu = menu.RegisterProvider(app)

	//-------------- oa --------------
	app.OA,
		app.OACalendar,
		app.OADial,
		app.OAJournal,
		app.OALiving,
		app.OAMeeting,
		app.OAMeetingRoom,
		app.OAPSTNCC,
		app.OASchedule,
		app.OAWebDrive = oa.RegisterProvider(app)

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
	case "MessageAppChat":
		return app.MessageAppChat
	case "MessageExternalContact":
		return app.MessageExternalContact
	case "MessageLinkedCorp":
		return app.MessageLinkedCorp

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
	case "ExternalContactGroupWelcomeTemplate":
		return app.ExternalContactGroupWelcomeTemplate
	case "ExternalContactSchool":
		return app.ExternalContactSchool
	case "ExternalContactMoment":
		return app.ExternalContactMoment
	case "ExternalContactMomentStrategy":
		return app.ExternalContactMomentStrategy
	case "ExternalContactMessageTemplate":
		return app.ExternalContactMessageTemplate

	case "AccountService":
		return app.AccountService
	case "AccountServiceCustomer":
		return app.AccountServiceCustomer
	case "AccountServiceMessage":
		return app.AccountServiceMessage
	case "AccountServiceServicer":
		return app.AccountServiceServicer
	case "AccountServiceState":
		return app.AccountServiceState
	case "AccountServiceTag":
		return app.AccountServiceTag

	case "Menu":
		return app.Menu
	case "OA":
		return app.OA
	case "OACalendar":
		return app.OACalendar
	case "OADial":
		return app.OADial
	case "OAJournal":
		return app.OAJournal
	case "OALiving":
		return app.OALiving
	case "OAMeeting":
		return app.OAMeeting
	case "OAMeetingRoom":
		return app.OAMeetingRoom
	case "OAPSTNCC":
		return app.OAPSTNCC
	case "OASchedule":
		return app.OASchedule
	case "OAWebDrive":
		return app.OAWebDrive

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

func (app *Work) SetExternalRequest(r *http.Request) {
	app.Base.ExternalRequest = r
}

func (app *Work) GetExternalRequest() (r *http.Request) {
	return app.Base.ExternalRequest
}

func MapUserConfig(userConfig *UserConfig) (*object.HashMap, error) {

	config := &object.HashMap{
		"corp_id":      userConfig.CorpID,
		"agent_id":     userConfig.AgentID,
		"secret":       userConfig.Secret,
		"token":        userConfig.Token,
		"aes_key":      userConfig.AESKey,
		"callback_url": userConfig.CallbackURL,

		"response_type": userConfig.ResponseType,
		"log": &object.StringMap{
			"level": userConfig.Log.Level,
			"file":  userConfig.Log.File,
		},
		"oauth.callback": userConfig.OAuth.Callback,
		"oauth.scopes":   userConfig.OAuth.Scopes,
		"cache":          userConfig.Cache,
		"http_debug":     userConfig.HttpDebug,
		"debug":          userConfig.Debug,
	}

	return config, nil

}
