package request

type RequestCustomerStrategyCreate struct {
	ParentID     int64            `json:"parent_id" `
	StrategyName string           `json:"strategy_name"`
	AdminList    []string         `json:"admin_list"`
	Privilege    RequestCustomerStrategyPrivilege   `json:"privilege"`
	Range        []RequestCustomerStrategyRange `json:"range"`
}

type RequestCustomerStrategyPrivilege struct {
	ViewCustomerList        bool `json:"view_customer_list"`
	ViewCustomerData        bool `json:"view_customer_data"`
	ViewRoomList            bool `json:"view_room_list"`
	ContactMe               bool `json:"contact_me"`
	JoinRoom                bool `json:"join_room"`
	ShareCustomer           bool `json:"share_customer"`
	OperResignCustomer      bool `json:"oper_resign_customer"`
	SendCustomerMsg         bool `json:"send_customer_msg"`
	EditWelcomeMsg          bool `json:"edit_welcome_msg"`
	ViewBehaviorData        bool `json:"view_behavior_data"`
	ViewRoomData            bool `json:"view_room_data"`
	SendGroupMsg            bool `json:"send_group_msg"`
	RoomDeduplication       bool `json:"room_deduplication"`
	RapidReply              bool `json:"rapid_reply"`
	OnjobCustomerTransfer   bool `json:"onjob_customer_transfer"`
	EditAntiSpamRule        bool `json:"edit_anti_spam_rule"`
	ExportCustomerList      bool `json:"export_customer_list"`
	ExportCustomerData      bool `json:"export_customer_data"`
	ExportCustomerGroupList bool `json:"export_customer_group_list"`
	ManageCustomerTag       bool `json:"manage_customer_tag"`
}

type RequestCustomerStrategyRange struct {
	Type    int    `json:"type"`
	PartyID int    `json:"partyid"`
	UserID  string `json:"userid"`
}
