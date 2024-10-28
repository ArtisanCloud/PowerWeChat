package request

type RequestApplyForBusiness struct {
	BusinessCode string `json:"business_code"`
	ContactInfo  struct {
		BusinessAuthorizationLetter string `json:"business_authorization_letter,omitempty"`
		ContactEmail                string `json:"contact_email,omitempty"`
		ContactIdDocCopy            string `json:"contact_id_doc_copy,omitempty"`
		ContactIdDocCopyBack        string `json:"contact_id_doc_copy_back,omitempty"`
		ContactIdDocType            string `json:"contact_id_doc_type,omitempty"`
		ContactIdNumber             string `json:"contact_id_number,omitempty"`
		ContactName                 string `json:"contact_name,omitempty"`
		ContactPeriodBegin          string `json:"contact_period_begin,omitempty"`
		ContactPeriodEnd            string `json:"contact_period_end,omitempty"`
		ContactType                 string `json:"contact_type,omitempty"`
		MobilePhone                 string `json:"mobile_phone,omitempty"`
		Openid                      string `json:"openid,omitempty"`
	} `json:"contact_info,omitempty"`
	SubjectInfo struct {
		SubjectType         string `json:"subject_type,omitempty"`
		FinanceInstitution  bool   `json:"finance_institution"`
		BusinessLicenseInfo struct {
			LegalPerson    string `json:"legal_person,omitempty"`
			LicenseAddress string `json:"license_address,omitempty"`
			LicenseCopy    string `json:"license_copy,omitempty"`
			LicenseNumber  string `json:"license_number,omitempty"`
			MerchantName   string `json:"merchant_name,omitempty"`
			PeriodBegin    string `json:"period_begin,omitempty"`
			PeriodEnd      string `json:"period_end,omitempty"`
		} `json:"business_license_info,omitempty"`
		CertificateLetterCopy  string `json:"certificate_letter_copy,omitempty"`
		FinanceInstitutionInfo *struct {
			FinanceLicensePics []string `json:"finance_license_pics,omitempty"`
			FinanceType        string   `json:"finance_type,omitempty"`
		} `json:"finance_institution_info,omitempty"`
		IdentityInfo struct {
			AuthorizeLetterCopy string `json:"authorize_letter_copy,omitempty"`
			IdCardInfo          struct {
				CardPeriodBegin string `json:"card_period_begin,omitempty"`
				CardPeriodEnd   string `json:"card_period_end,omitempty"`
				IdCardAddress   string `json:"id_card_address,omitempty"`
				IdCardCopy      string `json:"id_card_copy,omitempty"`
				IdCardName      string `json:"id_card_name,omitempty"`
				IdCardNational  string `json:"id_card_national,omitempty"`
				IdCardNumber    string `json:"id_card_number,omitempty"`
			} `json:"id_card_info,omitempty"`
			IdDocInfo *struct {
				IdDocCopy      string `json:"id_doc_copy,omitempty"`
				IdDocCopyBack  string `json:"id_doc_copy_back,omitempty"`
				IdDocName      string `json:"id_doc_name,omitempty"`
				IdDocNumber    string `json:"id_doc_number,omitempty"`
				IdDocAddress   string `json:"id_doc_address,omitempty"`
				DocPeriodBegin string `json:"doc_period_begin,omitempty"`
				DocPeriodEnd   string `json:"doc_period_end,omitempty"`
			} `json:"id_doc_info,omitempty"`
			IdDocType    string `json:"id_doc_type,omitempty"`
			IdHolderType string `json:"id_holder_type,omitempty"`
			Owner        bool   `json:"owner"`
		} `json:"identity_info,omitempty"`
		UboInfoList []struct {
			UboIdDocAddress  string `json:"ubo_id_doc_address,omitempty"`
			UboIdDocCopy     string `json:"ubo_id_doc_copy,omitempty"`
			UboIdDocCopyBack string `json:"ubo_id_doc_copy_back,omitempty"`
			UboIdDocName     string `json:"ubo_id_doc_name,omitempty"`
			UboIdDocNumber   string `json:"ubo_id_doc_number,omitempty"`
			UboIdDocType     string `json:"ubo_id_doc_type,omitempty"`
			UboPeriodBegin   string `json:"ubo_period_begin,omitempty"`
			UboPeriodEnd     string `json:"ubo_period_end,omitempty"`
		} `json:"ubo_info_list,omitempty"`
	} `json:"subject_info,omitempty"`
	BusinessInfo struct {
		MerchantShortname string `json:"merchant_shortname,omitempty"`
		SalesInfo         struct {
			AppInfo *struct {
				AppAppid string   `json:"app_appid,omitempty"`
				AppPics  []string `json:"app_pics,omitempty"`
			} `json:"app_info,omitempty"`
			BizStoreInfo *struct {
				BizAddressCode   string   `json:"biz_address_code,omitempty"`
				BizStoreAddress  string   `json:"biz_store_address,omitempty"`
				BizStoreName     string   `json:"biz_store_name,omitempty"`
				BizSubAppid      string   `json:"biz_sub_appid,omitempty"`
				IndoorPic        []string `json:"indoor_pic,omitempty"`
				StoreEntrancePic []string `json:"store_entrance_pic,omitempty"`
			} `json:"biz_store_info,omitempty"`
			MiniProgramInfo struct {
				MiniProgramAppid    string   `json:"mini_program_appid,omitempty"`
				MiniProgramSubAppid string   `json:"mini_program_sub_appid,omitempty"`
				MiniProgramPics     []string `json:"mini_program_pics,omitempty"`
			} `json:"mini_program_info,omitempty"`
			MpInfo *struct {
				MpAppid string   `json:"mp_appid,omitempty"`
				MpPics  []string `json:"mp_pics,omitempty"`
			} `json:"mp_info,omitempty"`
			SalesScenesType []string `json:"sales_scenes_type,omitempty"`
			WebInfo         *struct {
				Domain           string `json:"domain,omitempty"`
				WebAppid         string `json:"web_appid,omitempty"`
				WebAuthorisation string `json:"web_authorisation,omitempty"`
			} `json:"web_info,omitempty"`
			WeworkInfo *struct {
				CorpId     string   `json:"corp_id,omitempty"`
				SubCorpId  string   `json:"sub_corp_id,omitempty"`
				WeworkPics []string `json:"wework_pics,omitempty"`
			} `json:"wework_info,omitempty"`
		} `json:"sales_info,omitempty"`
		ServicePhone string `json:"service_phone,omitempty"`
	} `json:"business_info,omitempty"`
	SettlementInfo struct {
		ActivitiesAdditions []string  `json:"activities_additions,omitempty"`
		ActivitiesId        string    `json:"activities_id,omitempty"`
		ActivitiesRate      string    `json:"activities_rate,omitempty"`
		QualificationType   string    `json:"qualification_type,omitempty"`
		Qualifications      *[]string `json:"qualifications,omitempty"`
		SettlementId        string    `json:"settlement_id,omitempty"`
	} `json:"settlement_info,omitempty"`
	BankAccountInfo struct {
		AccountBank     string `json:"account_bank,omitempty"`
		AccountName     string `json:"account_name,omitempty"`
		AccountNumber   string `json:"account_number,omitempty"`
		BankAccountType string `json:"bank_account_type,omitempty"`
		BankAddressCode string `json:"bank_address_code,omitempty"`
		BankBranchId    string `json:"bank_branch_id,omitempty"`
	} `json:"bank_account_info,omitempty"`
	AdditionInfo struct {
		LegalPersonCommitment string   `json:"legal_person_commitment,omitempty"`
		LegalPersonVideo      string   `json:"legal_person_video,omitempty"`
		BusinessAdditionPics  []string `json:"business_addition_pics,omitempty"`
		BusinessAdditionMsg   string   `json:"business_addition_msg,omitempty"`
	} `json:"addition_info,omitempty"`
}
