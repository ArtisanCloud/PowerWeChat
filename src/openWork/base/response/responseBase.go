package response

import "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"

type FuncScopeCategory struct {
	ID int `json:"id"`
}

type ConfirmInfo struct {
	NeedConfirm    int `json:"need_confirm"`
	AlreadyConfirm int `json:"already_confirm"`
	CanConfirm     int `json:"can_confirm"`
}

type FuncInfo struct {
	FuncScopeCategory *FuncScopeCategory `json:"funcscope_category"`

	ConfirmInfo *ConfirmInfo `json:"confirm_info,omitempty"`
}

type AuthorizationInfo struct {
	AuthorizerAppid        string     `json:"authorizer_appid"`
	AuthorizerAccessToken  string     `json:"authorizer_access_token"`
	ExpiresIn              int        `json:"expires_in"`
	AuthorizerRefreshToken string     `json:"authorizer_refresh_token"`
	FuncInfos              []FuncInfo `json:"func_info"`
}

type ResponseHandleAuthorize struct {
	response.ResponseOpenPlatform
	
	AuthorizationInfo *AuthorizationInfo `json:"authorization_info"`
}

// ------------------------------------------------------------------------------------------------------------------------

type ServiceTypeInfo struct {
	ID int `json:"id"`
}

type VerifyTypeInfo struct {
	ID int `json:"id"`
}

type BusinessInfo struct {
	OpenStore int `json:"open_store"`
	OpenScan  int `json:"open_scan"`
	OpenPay   int `json:"open_pay"`
	OpenCard  int `json:"open_card"`
	OpenShake int `json:"open_shake"`
}

type Network struct {
	RequestDomain        []string `json:"RequestDomain"`
	WsRequestDomain      []string `json:"WsRequestDomain"`
	UploadDomain         []string `json:"UploadDomain"`
	DownloadDomain       []string `json:"DownloadDomain"`
	BizDomain            []string `json:"BizDomain"`
	UDPDomain            []string `json:"UDPDomain"`
	TCPDomain            []string `json:"TCPDomain"`
	PrefetchDNSDomain    []string `json:"PrefetchDNSDomain"`
	NewRequestDomain     []string `json:"NewRequestDomain"`
	NewWsRequestDomain   []string `json:"NewWsRequestDomain"`
	NewUploadDomain      []string `json:"NewUploadDomain"`
	NewDownloadDomain    []string `json:"NewDownloadDomain"`
	NewBizDomain         []string `json:"NewBizDomain"`
	NewUDPDomain         []string `json:"NewUDPDomain"`
	NewTCPDomain         []string `json:"NewTCPDomain"`
	NewPrefetchDNSDomain []string `json:"NewPrefetchDNSDomain"`
}

type Category struct {
	First  string `json:"first"`
	Second string `json:"second"`
}

type MiniProgramInfo struct {
	Network     *Network   `json:"network"`
	Categories  []Category `json:"categories"`
	VisitStatus int        `json:"visit_status"`
}

type BasicConfig struct {
	IsPhoneConfigured bool `json:"is_phone_configured"`
	IsEmailConfigured bool `json:"is_email_configured"`
}

type AuthorizerInfo struct {
	NickName        string           `json:"nick_name"`
	HeadImg         string           `json:"head_img"`
	ServiceTypeInfo *ServiceTypeInfo `json:"service_type_info"`
	VerifyTypeInfo  *VerifyTypeInfo  `json:"verify_type_info"`
	UserName        string           `json:"user_name"`
	PrincipalName   string           `json:"principal_name"`
	BusinessInfo    *BusinessInfo    `json:"business_info"`
	Alias           string           `json:"alias"`
	QrcodeURL       string           `json:"qrcode_url"`
	AccountStatus   int              `json:"account_status"`

	// used for MiniProgram
	Idc             int              `json:"idc"`
	Signature       string           `json:"signature"`
	MiniProgramInfo *MiniProgramInfo `json:"MiniProgramInfo"`
	RegisterType    int              `json:"register_type"`
	BasicConfig     BasicConfig      `json:"basic_config"`
}

type ResponseGetAuthorizer struct {
	response.ResponseOpenPlatform

	AuthorizerInfo    *AuthorizerInfo    `json:"authorizer_info"`
	AuthorizationInfo *AuthorizationInfo `json:"authorization_info"`
}

// ------------------------------------------------------------------------------------------------------------------------

type ResponseGetAuthorizerOption struct {
	response.ResponseOpenPlatform

	AuthorizerAppid string `json:"authorizer_appid"`
	OptionName      string `json:"option_name"`
	OptionValue     string `json:"option_value"`
}

// ------------------------------------------------------------------------------------------------------------------------

type ResponseCreatePreAuthorizationCode struct {
	response.ResponseOpenPlatform

	PreAuthCode string `json:"pre_auth_code"`
	ExpiresIn   int    `json:"expires_in"`
}
