package request

type RequestApplyForBusiness struct {
	BusinessCode string `json:"business_code"`
	ContactInfo  struct {
		BusinessAuthorizationLetter string `json:"business_authorization_letter"`
		ContactEmail                string `json:"contact_email"`
		ContactIdDocCopy            string `json:"contact_id_doc_copy"`
		ContactIdDocCopyBack        string `json:"contact_id_doc_copy_back"`
		ContactIdDocType            string `json:"contact_id_doc_type"`
		ContactIdNumber             string `json:"contact_id_number"`
		ContactName                 string `json:"contact_name"`
		ContactPeriodBegin          string `json:"contact_period_begin"`
		ContactPeriodEnd            string `json:"contact_period_end"`
		ContactType                 string `json:"contact_type"`
		MobilePhone                 string `json:"mobile_phone"`
		Openid                      string `json:"openid"`
	} `json:"contact_info"`
	SubjectInfo struct {
		SubjectType         string `json:"subject_type"`
		FinanceInstitution  bool   `json:"finance_institution"`
		BusinessLicenseInfo struct {
			LegalPerson    string `json:"legal_person"`
			LicenseAddress string `json:"license_address"`
			LicenseCopy    string `json:"license_copy"`
			LicenseNumber  string `json:"license_number"`
			MerchantName   string `json:"merchant_name"`
			PeriodBegin    string `json:"period_begin"`
			PeriodEnd      string `json:"period_end"`
		} `json:"business_license_info"`
		CertificateLetterCopy  string `json:"certificate_letter_copy"`
		FinanceInstitutionInfo struct {
			FinanceLicensePics []string `json:"finance_license_pics"`
			FinanceType        string   `json:"finance_type"`
		} `json:"finance_institution_info"`
		IdentityInfo struct {
			AuthorizeLetterCopy string `json:"authorize_letter_copy"`
			IdCardInfo          struct {
				CardPeriodBegin string `json:"card_period_begin"`
				CardPeriodEnd   string `json:"card_period_end"`
				IdCardAddress   string `json:"id_card_address"`
				IdCardCopy      string `json:"id_card_copy"`
				IdCardName      string `json:"id_card_name"`
				IdCardNational  string `json:"id_card_national"`
				IdCardNumber    string `json:"id_card_number"`
			} `json:"id_card_info"`
			IdDocInfo struct {
				IdDocCopy      string `json:"id_doc_copy"`
				IdDocCopyBack  string `json:"id_doc_copy_back"`
				IdDocName      string `json:"id_doc_name"`
				IdDocNumber    string `json:"id_doc_number"`
				IdDocAddress   string `json:"id_doc_address"`
				DocPeriodBegin string `json:"doc_period_begin"`
				DocPeriodEnd   string `json:"doc_period_end"`
			} `json:"id_doc_info"`
			IdDocType    string `json:"id_doc_type"`
			IdHolderType string `json:"id_holder_type"`
			Owner        bool   `json:"owner"`
		} `json:"identity_info"`
		UboInfoList []struct {
			UboIdDocAddress  string `json:"ubo_id_doc_address"`
			UboIdDocCopy     string `json:"ubo_id_doc_copy"`
			UboIdDocCopyBack string `json:"ubo_id_doc_copy_back"`
			UboIdDocName     string `json:"ubo_id_doc_name"`
			UboIdDocNumber   string `json:"ubo_id_doc_number"`
			UboIdDocType     string `json:"ubo_id_doc_type"`
			UboPeriodBegin   string `json:"ubo_period_begin"`
			UboPeriodEnd     string `json:"ubo_period_end"`
		} `json:"ubo_info_list"`
	} `json:"subject_info"`
	BusinessInfo struct {
		MerchantShortname string `json:"merchant_shortname"`
		SalesInfo         struct {
			AppInfo struct {
				AppAppid string   `json:"app_appid"`
				AppPics  []string `json:"app_pics"`
			} `json:"app_info"`
			BizStoreInfo struct {
				BizAddressCode   string   `json:"biz_address_code"`
				BizStoreAddress  string   `json:"biz_store_address"`
				BizStoreName     string   `json:"biz_store_name"`
				BizSubAppid      string   `json:"biz_sub_appid"`
				IndoorPic        []string `json:"indoor_pic"`
				StoreEntrancePic []string `json:"store_entrance_pic"`
			} `json:"biz_store_info"`
			MiniProgramInfo struct {
				MiniProgramAppid    string   `json:"mini_program_appid"`
				MiniProgramSubAppid string   `json:"mini_program_sub_appid"`
				MiniProgramPics     []string `json:"mini_program_pics"`
			} `json:"mini_program_info"`
			MpInfo struct {
				MpAppid string   `json:"mp_appid"`
				MpPics  []string `json:"mp_pics"`
			} `json:"mp_info"`
			SalesScenesType []string `json:"sales_scenes_type"`
			WebInfo         struct {
				Domain           string `json:"domain"`
				WebAppid         string `json:"web_appid"`
				WebAuthorisation string `json:"web_authorisation"`
			} `json:"web_info"`
			WeworkInfo struct {
				CorpId     string   `json:"corp_id"`
				SubCorpId  string   `json:"sub_corp_id"`
				WeworkPics []string `json:"wework_pics"`
			} `json:"wework_info"`
		} `json:"sales_info"`
		ServicePhone string `json:"service_phone"`
	} `json:"business_info"`
	SettlementInfo struct {
		ActivitiesAdditions []string `json:"activities_additions"`
		ActivitiesId        string   `json:"activities_id"`
		ActivitiesRate      string   `json:"activities_rate"`
		QualificationType   string   `json:"qualification_type"`
		Qualifications      []string `json:"qualifications"`
		SettlementId        string   `json:"settlement_id"`
	} `json:"settlement_info"`
	BankAccountInfo struct {
		AccountBank     string `json:"account_bank"`
		AccountName     string `json:"account_name"`
		AccountNumber   string `json:"account_number"`
		BankAccountType string `json:"bank_account_type"`
		BankAddressCode string `json:"bank_address_code"`
		BankBranchId    string `json:"bank_branch_id"`
	} `json:"bank_account_info"`
	AdditionInfo struct {
		LegalPersonCommitment string   `json:"legal_person_commitment"`
		LegalPersonVideo      string   `json:"legal_person_video"`
		BusinessAdditionPics  []string `json:"business_addition_pics"`
		BusinessAdditionMsg   string   `json:"business_addition_msg"`
	} `json:"addition_info"`
}
