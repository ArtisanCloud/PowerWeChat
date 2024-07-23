package request

// 该接口用于将一个网络上的视频拉取上传到平台。
// 不填写 cover_url 字段时默认截取视频首帧作为封面。
// Content-Type 需要指定为 application/json
// 该接口为异步接口，上传完成会推送上传完成事件到开发者服务器，开发者也可以调用"查询任务"接口来轮询上传结果。

type VideoMediaUploadByURLRequest struct {
	MediaUrl      string `json:"media_url"`                // 视频 URL，示例值："https://developers.weixin.qq.com/test.mp4"。
	CoverUrl      string `json:"cover_url"`                // 封面图片 URL，示例值："https://developers.weixin.qq.com/test.jpg"
	MediaName     string `json:"media_name"`               // 文件名，需按照“剧目名 - 对应剧集数”格式命名文件，示例值："我的中国梦 - 第1集"。 经测试 可以不用严格按照此格式
	SourceContext string `json:"source_context,omitempty"` // 来源上下文，会在上传完成事件中透传给开发者。
}

// 该接口用于将服务端本地文件上传到平台
type VideoMediaUploadByFileRequest struct {
	MediaName     string `json:"media_name"`               // 文件名，需按照“剧目名 - 对应剧集数”格式命名文件，示例值："我的中国梦 - 第1集"。 经测试 可以不用严格按照此格式
	MediaType     string `json:"media_type"`               // 视频格式，支持：MP4，TS，MOV，MXF，MPG，FLV，WMV，AVI，M4V，F4V，MPEG，3GP，ASF，MKV，示例值："MP4"。
	MediaData     string `json:"media_data"`               // 视频服务端路径
	CoverType     string `json:"cover_type,omitempty"`     // 视频封面图格式，支持：JPG、JPEG、PNG、BMP、TIFF、AI、CDR、EPS、TIF，示例值："JPG"。
	CoverData     string `json:"cover_data,omitempty"`     // 视频封面图文件内容，二进制。
	SourceContext string `json:"source_context,omitempty"` // 来源上下文，会在上传完成事件中透传给开发者。
}

type VideoApplyChunkUploadByIdRequest struct {
	MediaName     string `json:"media_name"`               // 文件名，需按照“剧目名 - 对应剧集数”格式命名文件，示例值："我的中国梦 - 第1集"。 经测试 可以不用严格按照此格式
	MediaType     string `json:"media_type"`               // 视频格式，支持：MP4，TS，MOV，MXF，MPG，FLV，WMV，AVI，M4V，F4V，MPEG，3GP，ASF，MKV，示例值："MP4"。
	CoverType     string `json:"cover_type,omitempty"`     // 视频封面图格式，支持：JPG、JPEG、PNG、BMP、TIFF、AI、CDR、EPS、TIF，示例值："JPG"。
	SourceContext string `json:"source_context,omitempty"` // 来源上下文，会在上传完成事件中透传给开发者。
}

type VideoApplyChunkUploadRequest struct {
	UploadId     string `json:"upload_id"`     // 分片上传ID 必填
	ChunkId      int64  `json:"chunk_id"`      // 分片唯一ID
	ChunkPath    string `json:"chunk_path"`    // 文件路径
	ResourceType int64  `json:"resource_type"` // 指定该分片属于视频还是图片的枚举值：1. 视频，2. 图片。
}

type VideoChunkUploadCompleteRequest struct {
	UploadId       string      `json:"upload_id"`                  // 分片上传ID 必填
	MediaPartInfos []*PartInfo `json:"media_part_infos"`           // 本次分片上传中媒体文件每个分片的信息。
	CoverPartInfos []*PartInfo `json:"cover_part_infos,omitempty"` // 本次分片上传中封面图片文件每个分片的信息。
}

type PartInfo struct {
	PartNumber int64  `json:"part_number"` // 分片编号。
	Etag       string `json:"etag"`        // 使用上传分片接口上传成功后返回的 etag 的值
}

// 获取媒资列表
// media_name 参数支持模糊匹配，当需要模糊匹配时可以在前面或后面加上 %，否则为精确匹配。例如 "test%" 可以匹配到 "test123", "testxxx", "test"。

type MediaListRequest struct {
	DramaId   int64  `json:"drama_id,omitempty"`   // 根据剧目id获取剧集信息。
	MediaName string `json:"media_name,omitempty"` // 媒资文件名，模糊匹配。
	StartTime int64  `json:"start_time,omitempty"` // 媒资上传时间>=start_time。
	EndTime   int64  `json:"end_time,omitempty"`   // 媒资上传时间<end_time。
	Limit     int64  `json:"limit,omitempty"`      // 分页拉取的最大返回结果数。默认值：100；最大值：100。
	Offset    int64  `json:"offset,omitempty"`     // 分页拉取的起始偏移量。默认值：0。
}

// 获取媒资播放连接
type GetMediaLinkRequest struct {
	MediaId int64  `json:"media_id"`         // 媒资id。
	T       int64  `json:"t"`                // 播放地址的过期时间戳。有效的时间最长不能超过2小时后。
	Us      string `json:"us,omitempty"`     // 链接标识。平台默认会生成一个仅包含小写字母和数字的字符串用于增强链接的唯一性(如us=647488c4792c15185b8fd2a6)。如开发者需要增加自己的标识，比如区分播放的渠道，可使用该参数，该参数最终的值是"开发者标识-平台标识"（如开发者传入abcd，则最终的临时链接中us=abcd-647488c4792c15185b8fd2a6）
	Expr    int64  `json:"expr,omitempty"`   // 试看时长，单位：秒，最大值不能超过视频长度
	Rlimit  int64  `json:"rlimit,omitempty"` // 最多允许多少个不同 IP 的终端播放，以十进制表示，最大值为9，不填表示不做限制。当限制 URL 只能被1个人播放时，建议 rlimit 不要严格限制成1（例如可设置为3），因为移动端断网后重连 IP 可能改变。
	Hhref   string `json:"whref,omitempty"`  // 允许访问的域名列表，支持1条 - 10条，用半角逗号分隔。域名前不要带协议名（http://和https://），域名为前缀匹配（如填写 abc.com，则 abc.com/123 和 abc.com.cn也会匹配），且支持通配符（如 *.abc.com）。
	Bkref   string `json:"bkref,omitempty"`  // 禁止访问的域名列表，支持1条 - 10条，用半角逗号分隔。域名前不要带协议名（http://和https://），域名为前缀匹配（如填写 abc.com，则 abc.com/123 和 abc.com.cn也会匹配），且支持通配符（如 *.abc.com）。
}

// 剧目提审
type SubmitAuditRequest struct {
	DramaId                            int64          `json:"drama_id"`                                        // 剧目id，首次提审不需要填该参数，重新提审时必填
	Name                               string         `json:"name"`                                            // 剧名，首次提审时必填，重新提审时根据是否需要修改选填。第一次必须填写！
	MediaCount                         int64          `json:"media_count"`                                     // 剧集数目。首次提审时必填， 重新提审时可不填，如要填写也要和第一次提审时一样。
	MediaIdList                        []int64        `json:"media_id_list"`                                   // 剧集媒资media_id列表。首次提审时必填，而且元素个数必须与media_count一致。重新提审时，如果剧集有内容有变化，则需要填写，而且元素个数也必须与 media_count一致。
	Description                        string         `json:"description"`                                     // 剧目简介，可填写200个字符。
	Recommendations                    string         `json:"recommendations"`                                 // 剧目推荐语，可填写15个字符。
	CoverMaterialId                    string         `json:"cover_material_id"`                               // 剧目海报临时media_id。首次提审时必填，重新提审时根据是否需要修改选填。
	PromotionPosterMaterialId          string         `json:"promotion_poster_material_id,omitempty"`          // 推广海报临时media_id。
	Producer                           string         `json:"producer"`                                        // 制作方 。首次提审时必填，重新提审时根据是否需要修改选填。
	AuthorizedMaterialId               string         `json:"authorized_material_id"`                          // 权利声明/播放授权材料material_id。 需上传《权利声明及不侵权承诺函》
	QualificationType                  int64          `json:"qualification_type"`                              // 剧目资质：1-取得《网络剧片发行许可证》或重点节目备案号；2-未取得《网络剧片发行许可证》或重点节目备案，且制作成本小于30万元
	RegistrationNumber                 string         `json:"registration_number,omitempty"`                   // 剧目备案号，当qualification_type=1时必填。根据提供的剧目资质证明文件填写对应的网络剧片发行许可证编号或剧目备案号。如：(沪)网剧审字(2023)第001号 或 V123456788888888
	QualificationCertificateMaterialId string         `json:"qualification_certificate_material_id,omitempty"` // 剧目资质证明文件，当qualification_type=1时必填。请提供网络剧片发行许可证或广电备案系统截图
	CostCommitmentLetterMaterialId     string         `json:"cost_commitment_letter_material_id,omitempty"`    // 《成本配置比例情况报告》material_id，当qualification_type=2时必填。
	CostOfProduction                   int64          `json:"cost_of_production,omitempty"`                    // 剧目制作成本（单位：万元），当qualification_type=2时必填。请填写“1-29” 的整数（如非整数将截断取整），数值需与《成本配置比例情况报告》中对应剧目制作成本一致。
	Expedited                          int64          `json:"expedited,omitempty"`                             // 填1表示审核加急，0或不填为不加急。每天有5次加急机会。该字段在首次提审时才有效，重新提审时会沿用首次提审时的属性，重新提审不会扣次数。最终是否为加急单，可以根据DramaInfo.expedited属性判断
	ActorList                          *ActorList     `json:"actor_list"`                                      // 演员信息，需填写2-5位演员。
	ReplaceMediaList                   []*ReplaceInfo `json:"replace_media_list,omitempty"`                    // 用于重新提审时替换审核不通过的剧集。
}

type ReplaceInfo struct {
	Old int64 `json:"old"`
	New int64 `json:"new"`
}

type ListRequest struct {
	Limit  int64 `json:"limit,omitempty"`  // 分页拉取的最大返回结果数。默认值：100；最大值：100。
	Offset int64 `json:"offset,omitempty"` // 分页拉取的起始偏移量。默认值：0。
}

// 剧目审核
type SubmitReplaceDramaAuditRequest struct {
	DramaId          int64          `json:"drama_id"`
	ReplaceMediaList []*ReplaceInfo `json:"replace_media_list"`
}

// 替换审核过的剧集
type ReplaceAuditDramaRequest struct {
	DramaId    int64 `json:"drama_id"`
	OldMediaId int64 `json:"old_media_id"`
	NewMediaId int64 `json:"new_media_id"`
}

// 剧目基本信息修改
type UpdateDramaInfoRequest struct {
	DramaId                            int64      `json:"drama_id"`
	Description                        string     `json:"description,omitempty"`
	CoverMaterialId                    string     `json:"cover_material_id,omitempty"`
	Recommendations                    string     `json:"recommendations,omitempty"`
	PromotionPosterMaterialId          string     `json:"promotion_poster_material_id,omitempty"`
	AlternateName                      string     `json:"alternate_name,omitempty"`
	ActorList                          *ActorList `json:"actor_list,omitempty"`
	QualificationType                  int64      `json:"qualification_type,omitempty"`
	QualificationCertificateMaterialId string     `json:"qualification_certificate_material_id,omitempty"`
	RegistrationNumber                 string     `json:"registration_number,omitempty"`
	CostOfProduction                   int64      `json:"cost_of_production,omitempty"`
	CostCommitmentLetterMaterialId     string     `json:"cost_commitment_letter_material_id,omitempty"`
}

type ActorList struct {
	Actor []*Actor `json:"actor,omitempty"`
}

type Actor struct {
	Name            string `json:"name"`              // 演员姓名，格式要求：可填写30个字符，支持输入中文、英文和·
	PhotoMaterialId string `json:"photo_material_id"` // 演员照片临时media_id
	Role            string `json:"role"`              // 饰演角色，格式要求：可填写30个字符，支持输入中文、英文及，。；”“？、—《》！（）-·
	Profile         string `json:"profile"`           // 演员简介，填写内容可参考：生日、星座、籍贯、身高、毕业院校、历史/代表作品、获奖信息、成就等(以上要素需至少满足3项或涵盖任一要素且不少于20字）。格式要求：可填写100个字符，支持输入中文、英文、阿拉伯数字及，。：；“”？、—《》！（）-·
}

// 查询剧目审核信息
type SearchDramaAuditInfoRequest struct {
	DramaId   int64 `json:"drama_id"`   // 剧目id
	AuditType int64 `json:"audit_type"` // 审核类型。0为首次提审；1为再次提审；2为替换剧集提审；3为修改剧目基本信息审核。
}

// 查询cdn
type SearchCdnInfoRequest struct {
	StartTime    int64 `json:"start_time"`
	EndTime      int64 `json:"end_time"`
	DataInterval int64 `json:"data_interval,omitempty"`
}

// 剧目授权
type DramaAuthorizedRequest struct {
	AuthorizedAppid string  `json:"authorized_appid"`
	DramaId         []int64 `json:"drama_id"`
	AuthzExpireTime int64   `json:"authz_expire_time,omitempty"`
}

// 查询剧目授权
type SearchDramaAuthorizedRequest struct {
	DramaId int64 `json:"drama_id"`
	ListRequest
}

// 查询剧目被授权
type SearchDoDramaAuthorizedRequest struct {
	AuthorizerAppid string `json:"authorizer_appid"`
	ListRequest
}

// 查询账号授权
type AccountAuthorizedRequest struct {
	AuthorizedAppid string `json:"authorized_appid"`
	AuthzExpireTime int64  `json:"authz_expire_time,omitempty"`
}
