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
	DramaId                  int64   `json:"drama_id,omitempty"`                    // 剧目id，首次提审不需要填该参数，重新提审时必填
	Name                     string  `json:"name,omitempty"`                        // 剧名，首次提审时必填，重新提审时根据是否需要修改选填。第一次必须填写！
	MediaCount               int64   `json:"media_count,omitempty"`                 // 剧集数目。首次提审时必填， 重新提审时可不填，如要填写也要和第一次提审时一样。
	MediaIdList              []int64 `json:"media_id_list,omitempty"`               // 剧集媒资media_id列表。首次提审时必填，而且元素个数必须与media_count一致。重新提审时，如果剧集有内容有变化，则需要填写，而且元素个数也必须与 media_count一致。
	Producer                 string  `json:"producer,omitempty"`                    // 制作方 。首次提审时必填，重新提审时根据是否需要修改选填。
	CoverMaterialId          string  `json:"cover_material_id,omitempty"`           // 封面图片临时media_id。首次提审时必填，重新提审时根据是否需要修改选填。
	AuthorizedMaterialId     string  `json:"authorized_material_id,omitempty"`      // 《网络剧片发行许可证编号》。剧目播放授权材料material_id。如果小程序主体名称和制作方完全一致，则不需要填，否则必填
	RegistrationNumber       string  `json:"registration_number,omitempty"`         // 《剧目备案号》。首次提审时《剧目备案号》与《网络剧片发行许可证编号》二选一。重新提审时根据是否需要修改选填
	PublishLicense           string  `json:"publish_license,omitempty"`             // 网络剧片发行许可证《编号》。首次提审时《剧目备案号》与《网络剧片发行许可证编号》二选一。重新提审时根据是否需要修改选填
	PublishLicenseMaterialId string  `json:"publish_license_material_id,omitempty"` // 网络剧片发行许可证《图片》，首次提审时如果网《络剧片发行许可证编号》非空，则该改字段也非空。重新提审时根据是否变化选填
}

type GetDramaListRequest struct {
	Limit  int64 `json:"limit,omitempty"`  // 分页拉取的最大返回结果数。默认值：100；最大值：100。
	Offset int64 `json:"offset,omitempty"` // 分页拉取的起始偏移量。默认值：0。
}
