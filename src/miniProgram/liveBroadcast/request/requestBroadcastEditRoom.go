package request

type RequestBroadcastEditRoom struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	CoverImg      string `json:"coverImg"`
	StartTime     int    `json:"startTime"`
	EndTime       int    `json:"endTime"`
	AnchorName    string `json:"anchorName"`
	AnchorWechat  string `json:"anchorWechat"`
	ShareImg      string `json:"shareImg"`
	CloseLike     int    `json:"closeLike"`
	CloseGoods    int    `json:"closeGoods"`
	CloseComment  int    `json:"closeComment"`
	IsFeedsPublic int    `json:"isFeedsPublic"`
	CloseReplay   int    `json:"closeReplay"`
	CloseShare    int    `json:"closeShare"`
	CloseKF       int    `json:"closeKf"`
	FeedsImg      string `json:"feedsImg"`
}
