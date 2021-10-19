package journal

import (
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/src/work/oa/journal/response"
)

type Client struct {
	*kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) *Client {
	return &Client{
		kernel.NewBaseClient(&app, nil),
	}
}

// 批量获取汇报记录单号
// https://work.weixin.qq.com/api/doc/90000/90135/93393
func (comp *Client) GetRecordList(startTime int, endTime int, nextCursor int, size int, filters []*power.StringMap) (*response.ResponseJournalGetRecordList, error) {

	result := &response.ResponseJournalGetRecordList{}

	if size > 100 {
		size = 100
	}

	options := &object.HashMap{
		"starttime": fmt.Sprintf("%d", startTime),
		"endtime":   fmt.Sprintf("%d", endTime),
		"cursor":    fmt.Sprintf("%d", nextCursor),
		"limit":     size,
		"filters":   filters,
	}

	_, err := comp.HttpPostJson("cgi-bin/oa/journal/get_record_list", options, nil, nil, result)

	return result, err
}

// 获取汇报记录详情
// https://work.weixin.qq.com/api/doc/90000/90135/93394
func (comp *Client) GetRecordDetail(JournalUUID string) (*response.ResponseJournalGetRecordDetail, error) {

	result := &response.ResponseJournalGetRecordDetail{}

	options := &object.HashMap{
		"journaluuid": JournalUUID,
	}

	_, err := comp.HttpPostJson("cgi-bin/oa/journal/get_record_detail", options, nil, nil, result)

	return result, err
}

// 获取汇报记录详情
// https://work.weixin.qq.com/api/doc/90000/90135/93395
func (comp *Client) GetStatList(templateID string, startTime int, endTime int) (*response.ResponseJournalGetStatList, error) {

	result := &response.ResponseJournalGetStatList{}

	options := &object.HashMap{
		"template_id": templateID,
		"starttime":   fmt.Sprintf("%d", startTime),
		"endtime":     fmt.Sprintf("%d", endTime),
	}

	_, err := comp.HttpPostJson("cgi-bin/oa/journal/get_stat_list", options, nil, nil, result)

	return result, err
}

