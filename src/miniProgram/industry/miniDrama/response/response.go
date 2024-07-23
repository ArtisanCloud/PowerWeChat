package response

type BaseResponse struct {
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

// 拉取上传Response
type VideoMediaUploadByURLResponse struct {
	BaseResponse
	TaskId int64 `json:"task_id"` // 任务id 用于查询
}

// 查询任务 使用上面返回的task id 进行查询
type VideoMediaSearchTaskResponse struct {
	BaseResponse
	TaskInfo struct {
		Id         int64  `json:"id"`          // 任务id。 不明白返回个0是啥意思
		TaskType   int64  `json:"task_type"`   // 任务类型枚举值：1. 拉取上传任务。
		Status     int64  `json:"status"`      // 任务状态枚举值：1. 等待中；2. 正在处理；3. 已完成；4. 失败。
		Errcode    int64  `json:"errcode"`     // 任务错误码，0表示成功，其它表示失败。
		Errmsg     string `json:"errmsg"`      // 任务错误原因
		CreateTime int64  `json:"create_time"` // 创建时间，时间戳。
		FinishTime int64  `json:"finish_time"` // 完成时间，时间戳。
		MediaId    int64  `json:"media_id"`    // 媒体文件id。
	} `json:"task_info"`
}

// 拉取上传Response
type VideoMediaUploadByFileResponse struct {
	BaseResponse
	MediaId int64 `json:"media_id"` // 任务id 用于查询
}

// 分片上传Id申请
type VideoApplyChunkUploadByIdResponse struct {
	BaseResponse
	UploadId string `json:"upload_id"`
}

// 分片上传结果
type VideoApplyChunkUploadResponse struct {
	BaseResponse
	Etag string `json:"etag"`
}

// 分片上传最终结果
type VideoChunkUploadCompleteResponse struct {
	BaseResponse
	MediaId string `json:"media_id"` // 媒资ID
}

// 审核信息。 文档写了 但是接口没返回 自主加了 omitempty
type AuditDetail struct {
	Status                 int64    `json:"status,omitempty"`                    // 0为无效值；1为审核中；2为审核驳回；3为审核通过；
	CreateTime             int64    `json:"create_time,omitempty"`               // 提审时间戳。
	AuditTime              int64    `json:"audit_time,omitempty"`                // 审核时间戳。
	Reason                 string   `json:"reason,omitempty"`                    // 审核备注。该值可能为空。
	EvidenceMaterialIdList []string `json:"evidence_material_id_list,omitempty"` // 审核证据截图id列表，截图id可以用作get_material接口的参数来获得截图内容。
}

// 演员列表
type ActorList struct {
	Actor []*Actor `json:"actor,omitempty"`
}

// 演员详细信息
type Actor struct {
	Name            string `json:"name"`              // 演员姓名，格式要求：可填写30个字符，支持输入中文、英文和·
	PhotoMaterialId string `json:"photo_material_id"` // 演员照片临时media_id
	Role            string `json:"role"`              // 饰演角色，格式要求：可填写30个字符，支持输入中文、英文及，。；”“？、—《》！（）-·
	Profile         string `json:"profile"`           // 演员简介，填写内容可参考：生日、星座、籍贯、身高、毕业院校、历史/代表作品、获奖信息、成就等(以上要素需至少满足3项或涵盖任一要素且不少于20字）。格式要求：可填写100个字符，支持输入中文、英文、阿拉伯数字及，。：；“”？、—《》！（）-·
}

// 媒资信息列表。
type MediaInfo struct {
	MediaId     int64        `json:"media_id"`     // 媒资文件id。
	CreateTime  int64        `json:"create_time"`  // 上传时间，时间戳。
	ExpireTime  int64        `json:"expire_time"`  // 过期时间，时间戳。
	DramaId     int64        `json:"drama_id"`     // 所属剧目id
	FileSize    int64        `json:"file_size"`    // 媒资文件大小，单位：字节。
	Duration    int64        `json:"duration"`     // 播放时长，单位：秒。
	Name        string       `json:"name"`         // 媒资文件名。
	Description string       `json:"description"`  // 描述。
	CoverUrl    string       `json:"cover_url"`    // 封面图临时链接。
	OriginalUrl string       `json:"original_url"` // 原始视频临时链接。
	Mp4Url      string       `json:"mp4_url"`      // mp4格式临时链接 。
	HlsUrl      string       `json:"hls_url"`      // hls格式临时链接。
	AuditDetail *AuditDetail `json:"audit_detail,omitempty"`
}

type MediaListResponse struct {
	BaseResponse
	MediaInfoList []*MediaInfo `json:"media_info_list"`
}

// 查询媒资信息
type MediaInfoResponse struct {
	BaseResponse
	MediaInfo MediaInfo `json:"media_info"`
}

// 获取媒资播放链接
type MediaLinkInfo struct {
	CoverUrl    string `json:"cover_url"`   // 封面图临时链接。
	Description string `json:"description"` // 描述。
	Duration    int    `json:"duration"`    // 播放时长，单位：秒。
	HlsUrl      string `json:"hls_url"`     // hls格式临时链接。
	Mp4Url      string `json:"mp4_url"`     // mp4格式临时链接 。
	Name        string `json:"name"`        // 媒资文件名。
	OriginalUrl string `json:"original_url"`
}

// 媒资播放链接
type MediaLinkResponse struct {
	BaseResponse
	MediaInfo MediaLinkInfo `json:"media_info"`
}

// 剧目提审
type SubmitAuditResponse struct {
	BaseResponse
	DramaId int64 `json:"drama_id"` // // 剧目id。
}

// 剧目信息
type DramaInfo struct {
	DramaId           int64      `json:"drama_id"`           // 剧目id。
	CreateTime        int64      `json:"create_time"`        // 创建时间，时间戳。
	Name              string     `json:"name"`               // 剧名。
	Playwright        string     `json:"playwright"`         // 编剧
	Producer          string     `json:"producer"`           // 制作方
	ProductionLicense string     `json:"production_license"` // 备案号
	CoverUrl          string     `json:"cover_url"`          // 封面图临时链接，根据提审时提交的cover_material_id转存得到。
	MediaCount        int64      `json:"media_count"`        // 剧集数目
	Recommendations   string     `json:"recommendations"`    // 推荐语
	Description       string     `json:"description"`        // 描述
	PromotionPoster   string     `json:"promotion_poster"`
	Status            int64      `json:"status"`
	MediaList         []struct { // 审核状态。
		MediaId int64 `json:"media_id"` // 媒资文件id。
	} `json:"media_list"`
	AuditDetail *AuditDetail `json:"audit_detail,omitempty"`
	ActorList   *ActorList   `json:"actor_list,omitempty"`
}

// 获取剧目列表
type GetDramaListResponse struct {
	BaseResponse
	DramaInfoList []*DramaInfo `json:"drama_info_list"`
}

// 获取剧目信息
type GetDramaInfoResponse struct {
	BaseResponse
	DramaInfo *DramaInfo `json:"drama_info"`
}

// 查询剧目审核
type SearchDramaAuditInfoResponse struct {
	BaseResponse
	AuditDetail *AuditDetail `json:"audit_detail"`
}

// 查询CDN
type SearchCdnInfoResponse struct {
	BaseResponse
	DataInterval int64          `json:"data_interval"`
	ItemList     []*CdnItemList `json:"item_list"`
}

type CdnItemList struct {
	Time  int64 `json:"time"`
	Value int64 `json:"value"`
}

// 查询CDN日志
type SearchCdnLogDownLoadUrlResponse struct {
	BaseResponse
	DomesticCdnLogs []*DomesticCdnLogs `json:"domestic_cdn_logs"`
	TotalCount      int64              `json:"total_count"`
}

type DomesticCdnLogs struct {
	Date      string `json:"date"`
	EndTime   int64  `json:"end_time"`
	Name      string `json:"name"`
	StartTime int64  `json:"start_time"`
	Url       string `json:"url"`
}

// 流量包详情
type ListPackageDetailResponse struct {
	BaseResponse
	PackageList []*PackageList `json:"package_list"`
	TotalCount  int            `json:"total_count"`
}

type PackageList struct {
	All       int64  `json:"all"`
	EndTime   int64  `json:"end_time"`
	IsDeleted int64  `json:"is_deleted"`
	OrderId   int64  `json:"order_id"`
	PackageId string `json:"package_id"`
	StartTime int64  `json:"start_time"`
	Status    int64  `json:"status"`
	Used      int64  `json:"used"`
}

// 剧目授权
type DramaAuthorizedResponse struct {
	BaseResponse
	Result []*DramaAuthorizedResult `json:"result"`
}

// 授权结果
type DramaAuthorizedResult struct {
	DramaId int64  `json:"drama_id"`
	Errcode int64  `json:"errcode"`
	Errmsg  string `json:"errmsg,omitempty"`
}

// 查询剧目授权
type SearchDramaAuthorizedResponse struct {
	BaseResponse
	Objects    []*AuthorizedObjects `json:"objects"`
	TotalCount int64                `json:"total_count"`
}

// 授权信息
type AuthorizedObjects struct {
	DramaId         int64  `json:"drama_id,omitempty"`
	AuthorizedAppid string `json:"authorized_appid"`
	AuthorizedTime  int64  `json:"authorized_time"`
	AuthzExpireTime int64  `json:"authz_expire_time"`
}

// 查询账号授权
type SearchAccountAuthorizedResponse struct {
	BaseResponse
	Objects []*AuthorizedObjects `json:"objects"`
}
