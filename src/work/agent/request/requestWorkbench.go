package request

type RequestSetWorkbenchTemplate struct {
	AgentID         int               `json:"agentid"`
	Type            string            `json:"type"`
	KeyData         WorkBenchKeyData  `json:"keydata"`
	Image           WorkBenchImage    `json:"image,omitempty"`
	List            WorkBenchListItem `json:"list,omitempty"`
	WebView         WorkBenchWebView  `json:"webview,omitempty"`
	ReplaceUserData bool              `json:"replace_user_data"`
}

type RequestSetWorkBenchData struct {
	AgentID int               `json:"agentid"`
	UserID  string            `json:"userid"`
	Type    string            `json:"type"`
	KeyData WorkBenchKeyData  `json:"keydata,omitempty"`
	Image   WorkBenchImage    `json:"image,omitempty"`
	List    WorkBenchListData `json:"list,omitempty"`
	WebView WorkBenchWebView  `json:"webview,omitempty"`
}

type WorkBenchKeyData struct {
	Items []WorkBenchKeyDataItem `json:"items"`
}
type WorkBenchKeyDataItem struct {
	Key      string `json:"key"`
	Data     string `json:"data"`
	JumpUrl  string `json:"jump_url"`
	PagePath string `json:"pagepath"`
}

type WorkBenchImage struct {
	Url      string `json:"url"`
	JumpUrl  string `json:"jump_url"`
	PagePath string `json:"pagepath"`
}

type WorkBenchListData struct {
	Items []WorkBenchListItem `json:"items"`
}
type WorkBenchListItem struct {
	Title    string `json:"title"`
	JumpUrl  string `json:"jump_url"`
	PagePath string `json:"pagepath"`
}

type WorkBenchWebView struct {
	Url      string `json:"url"`
	JumpUrl  string `json:"jump_url"`
	PagePath string `json:"pagepath"`
}
