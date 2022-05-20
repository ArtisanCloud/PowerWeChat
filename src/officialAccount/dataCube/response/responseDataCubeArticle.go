package response

import "github.com/ArtisanCloud/PowerWeChat/src/kernel/response"

type ArticleSummary struct {
	RefDate          string `json:"ref_date"`
	MsgID            string `json:"msgid"`
	Title            string `json:"title"`
	IntPageReadUser  int    `json:"int_page_read_user"`
	IntPageReadCount int    `json:"int_page_read_count"`
	OriPageReadUser  int    `json:"ori_page_read_user"`
	OriPageReadCount int    `json:"ori_page_read_count"`
	ShareUser        int    `json:"share_user"`
	ShareCount       int    `json:"share_count"`
	AddToFavUser     int    `json:"add_to_fav_user"`
	AddToFavCount    int    `json:"add_to_fav_count"`
}

type ResponseDataCubeArticleSummary struct {
	*response.ResponseOfficialAccount

	List []*ArticleSummary `json:"list"`
}

// ----------------------------------------------------------------------

type ArticleDetail struct {
	StatDate                    string `json:"stat_date"`
	TargetUser                  int    `json:"target_user"`
	IntPageReadUser             int    `json:"int_page_read_user"`
	IntPageReadCount            int    `json:"int_page_read_count"`
	OriPageReadUser             int    `json:"ori_page_read_user"`
	OriPageReadCount            int    `json:"ori_page_read_count"`
	ShareUser                   int    `json:"share_user"`
	ShareCount                  int    `json:"share_count"`
	AddToFavUser                int    `json:"add_to_fav_user"`
	AddToFavCount               int    `json:"add_to_fav_count"`
	IntPageFromSessionReadUser  int    `json:"int_page_from_session_read_user"`
	IntPageFromSessionReadCount int    `json:"int_page_from_session_read_count"`
	IntPageFromHistMsgReadUser  int    `json:"int_page_from_hist_msg_read_user"`
	IntPageFromHistMsgReadCount int    `json:"int_page_from_hist_msg_read_count"`
	IntPageFromFeedReadUser     int    `json:"int_page_from_feed_read_user"`
	IntPageFromFeedReadCount    int    `json:"int_page_from_feed_read_count"`
	IntPageFromFriendsReadUser  int    `json:"int_page_from_friends_read_user"`
	IntPageFromFriendsReadCount int    `json:"int_page_from_friends_read_count"`
	IntPageFromOtherReadUser    int    `json:"int_page_from_other_read_user"`
	IntPageFromOtherReadCount   int    `json:"int_page_from_other_read_count"`
	FeedShareFromSessionUser    int    `json:"feed_share_from_session_user"`
	FeedShareFromSessionCnt     int    `json:"feed_share_from_session_cnt"`
	FeedShareFromFeedUser       int    `json:"feed_share_from_feed_user"`
	FeedShareFromFeedCnt        int    `json:"feed_share_from_feed_cnt"`
	FeedShareFromOtherUser      int    `json:"feed_share_from_other_user"`
	FeedShareFromOtherCnt       int    `json:"feed_share_from_other_cnt"`
}

type ArticleTotal struct {
	RefDate string           `json:"ref_date"`
	Msgid   string           `json:"msgid"`
	Title   string           `json:"title"`
	Details []*ArticleDetail `json:"details"`
}

type ResponseDataCubeArticleTotal struct {
	*response.ResponseOfficialAccount

	List []*ArticleTotal `json:"list"`
}

// ----------------------------------------------------------------------

type UserReadSummary struct {
	RefDate          string `json:"ref_date"`
	UserSource       int    `json:"user_source"`
	IntPageReadUser  int    `json:"int_page_read_user"`
	IntPageReadCount int    `json:"int_page_read_count"`
	OriPageReadUser  int    `json:"ori_page_read_user"`
	OriPageReadCount int    `json:"ori_page_read_count"`
	ShareUser        int    `json:"share_user"`
	ShareCount       int    `json:"share_count"`
	AddToFavUser     int    `json:"add_to_fav_user"`
	AddToFavCount    int    `json:"add_to_fav_count"`
}

type ResponseDataCubeUserReadSummary struct {
	*response.ResponseOfficialAccount

	List []*UserReadSummary `json:"list"`
}

// ----------------------------------------------------------------------

type UserShareSummary struct {
	RefDate    string `json:"ref_date"`
	ShareScene int    `json:"share_scene"`
	ShareCount int    `json:"share_count"`
	ShareUser  int    `json:"share_user"`
}

type ResponseDataCubeUserShareSummary struct {
	*response.ResponseOfficialAccount

	List []*UserShareSummary `json:"list"`
}

// ----------------------------------------------------------------------

type UserShareHourly struct {
	RefDate    string `json:"ref_date"`
	RefHour    int    `json:"ref_hour"`
	ShareScene int    `json:"share_scene"`
	ShareCount int    `json:"share_count"`
	ShareUser  int    `json:"share_user"`
}

type ResponseDataCubeUserShareHourly struct {
	*response.ResponseOfficialAccount

	List []*UserShareHourly`json:"list"`
}

// ----------------------------------------------------------------------
