package work

import (
	"github.com/ArtisanCloud/PowerLibs/v3/cache"
	"github.com/ArtisanCloud/PowerLibs/v3/logger"
	"github.com/ArtisanCloud/PowerLibs/v3/logger/contract"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/providers"
	miniProgram2 "github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/accountService"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/accountService/customer"
	message3 "github.com/ArtisanCloud/PowerWeChat/v3/src/work/accountService/message"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/accountService/serviceState"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/accountService/servicer"
	tag3 "github.com/ArtisanCloud/PowerWeChat/v3/src/work/accountService/tag"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/agent"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/agent/workbench"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/auth"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/base"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/corpgroup"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/department"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/contactWay"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/customerStrategy"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/groupChat"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/groupWelcomeTemplate"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/messageTemplate"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/moment"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/momentStrategy"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/school"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/statistics"
	tag2 "github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/tag"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/externalContact/transfer"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/groupRobot"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/invoice"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/jssdk"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/media"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/menu"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/message"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/message/appChat"
	externalContact2 "github.com/ArtisanCloud/PowerWeChat/v3/src/work/message/externalContact"
	linkedCorp2 "github.com/ArtisanCloud/PowerWeChat/v3/src/work/message/linkedCorp"
	miniProgram "github.com/ArtisanCloud/PowerWeChat/v3/src/work/miniProgram"
	msgaudit "github.com/ArtisanCloud/PowerWeChat/v3/src/work/msgAudit"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/oa"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/oa/calendar"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/oa/dial"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/oa/journal"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/oa/living"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/oa/meeting"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/oa/meetingroom"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/oa/pstncc"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/oa/schedule"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/oa/webdrive"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/oauth"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/server"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/user"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/user/batchJobs"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/user/exportJobs"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/user/linkedCorp"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/user/tag"
)

type Work struct {
	*kernel.ServiceContainer

	Base        *base.Client
	AccessToken *auth.AccessToken
	OAuth       *oauth.Manager

	Config     *kernel.Config
	Department *department.Client

	JSSDK *jssdk.Client

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

	Logger *logger.Logger
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

	Http Http

	HttpDebug bool
	Debug     bool
	NotifyURL string
	Sandbox   bool
}

type Http struct {
	Timeout float64
	BaseURI string
}

type Log struct {
	Driver contract.LoggerInterface
	Level  string
	File   string
	Error  string
	ENV    string
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

	app.Logger, err = logger.NewLogger(app.Config.Get("log.driver", nil), &object.HashMap{
		"env":        app.Config.GetString("env", "develop"),
		"outputPath": app.Config.GetString("log.file", "./wechat/info.log"),
		"errorPath":  app.Config.GetString("log.error", "./wechat/error.log"),
	})
	if err != nil {
		return nil, err
	}

	//-------------- register auth --------------
	app.AccessToken, err = auth.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	//-------------- register Base --------------
	app.Base, err = base.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register JSSDK --------------
	app.JSSDK, err = jssdk.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register oauth --------------
	app.OAuth, err = oauth.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	//-------------- register agent --------------
	app.Agent,
		app.AgentWorkbench, err = agent.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register Department --------------
	app.Department, err = department.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register Message --------------
	app.Message, app.Messager,
		app.MessageAppChat,
		app.MessageExternalContact,
		app.MessageLinkedCorp, err = message.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register Encryptor --------------
	app.Encryptor, app.Server, err = server.RegisterProvider(app)
	if err != nil {
		return nil, err
	}
	//-------------- register user --------------
	app.User,
		app.UserBatchJobs,
		app.UserExportJobs,
		app.UserLinkedCorp,
		app.UserTag, err = user.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

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
		app.ExternalContactTransfer, err = externalContact.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- register account service --------------
	app.AccountService,
		app.AccountServiceCustomer,
		app.AccountServiceMessage,
		app.AccountServiceServicer,
		app.AccountServiceState,
		app.AccountServiceTag, err = accountService.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- media --------------
	app.Media, err = media.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- menu --------------
	app.Menu, err = menu.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

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
		app.OAWebDrive, err = oa.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- msg audit --------------
	app.MsgAudit, err = msgaudit.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- corp group --------------
	app.CorpGroup, err = corpgroup.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	//-------------- invoice --------------
	app.Invoice, err = invoice.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

	app.GroupRobot, app.GroupRobotMessenger, err = groupRobot.RegisterProvider(app)
	if err != nil {
		return nil, err
	}

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

	case "Logger":
		return app.Logger

	default:
		return nil
	}

}

func MapUserConfig(userConfig *UserConfig) (*object.HashMap, error) {

	baseURI := "https://qyapi.weixin.qq.com/"
	if userConfig.Http.BaseURI != "" {
		baseURI = userConfig.Http.BaseURI
	}
	timeout := 5.0
	if userConfig.Http.Timeout > 0 {
		timeout = userConfig.Http.Timeout
	}
	config := &object.HashMap{
		"corp_id":      userConfig.CorpID,
		"agent_id":     userConfig.AgentID,
		"secret":       userConfig.Secret,
		"token":        userConfig.Token,
		"aes_key":      userConfig.AESKey,
		"callback_url": userConfig.CallbackURL,

		"response_type": userConfig.ResponseType,
		"http": &object.HashMap{
			"timeout":  timeout,
			"base_uri": baseURI,
		},
		"log": &object.HashMap{
			"level": userConfig.Log.Level,
			"file":  userConfig.Log.File,
			"error": userConfig.Log.Error,
			"env":   userConfig.Log.ENV,
		},
		"oauth.callbacks": userConfig.OAuth.Callback,
		"oauth.scopes":    userConfig.OAuth.Scopes,
		"cache":           userConfig.Cache,
		"http_debug":      userConfig.HttpDebug,
		"debug":           userConfig.Debug,
	}

	return config, nil

}

func (app *Work) MiniProgram() (*miniProgram.Application, error) {

	userConfig, err := MapConfigToMiniProgramUserConfig(app)
	if err != nil {
		return nil, err
	}

	miniProgramApp, err := miniProgram.NewApplication(userConfig)
	miniProgramApp.Config = app.Config

	return miniProgramApp, err
}

func MapConfigToMiniProgramUserConfig(app kernel.ApplicationInterface) (userConfig *miniProgram.UserConfig, err error) {

	config := app.GetConfig()
	cache := config.Get("cache", nil).(cache.CacheInterface)
	log := config.Get("log", nil).(*object.HashMap)
	userConfig = &miniProgram.UserConfig{
		MiniProgramUserConfig: &miniProgram2.UserConfig{
			AppID:  config.GetString("corp_id", ""),
			Secret: config.GetString("secret", ""),

			ResponseType: config.GetString("response_type", ""),
			Log: miniProgram2.Log{
				Level: (*log)["level"].(string),
				File:  (*log)["file"].(string),
				ENV:   (*log)["env"].(string),
			},
			Cache:     cache,
			HttpDebug: config.GetBool("http_debug", false),
			Debug:     config.GetBool("debug", false),
		},

		CorpID:      config.GetString("corp_id", ""),
		AgentID:     config.GetInt("agent_id", 0),
		Secret:      config.GetString("secret", ""),
		Token:       config.GetString("token", ""),
		AESKey:      config.GetString("aes_key", ""),
		CallbackURL: config.GetString("callback_url", ""),
		Http:        config.Get("http", nil).(*object.HashMap),
	}

	return userConfig, err
}
