package request

type RequestBroadcastCreateRoom struct {
	Name            string `json:"name"`
	CoverImg        string `json:"coverImg"`
	StartTime       int64  `json:"startTime"`
	EndTime         int64  `json:"endTime"`
	AnchorName      string `json:"anchorName"`
	AnchorWechat    string `json:"anchorWechat"`
	SubAnchorWechat string `json:"subAnchorWechat"`
	CreaterWechat   string `json:"createrWechat"`
	ShareImg        string `json:"shareImg"`
	FeedsImg        string `json:"feedsImg"`
	IsFeedsPublic   int    `json:"isFeedsPublic"`
	Type            int    `json:"type"`
	CloseLike       int    `json:"closeLike"`
	CloseGoods      int    `json:"closeGoods"`
	CloseComment    int    `json:"closeComment"`
	CloseReplay     int    `json:"closeReplay"`
	CloseShare      int    `json:"closeShare"`
	CloseKf         int    `json:"closeKf"`
}
