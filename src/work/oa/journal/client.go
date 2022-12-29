package journal

import (
	"context"
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/power"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/oa/journal/response"
)

type Client struct {
	BaseClient *kernel.BaseClient
}

func NewClient(app kernel.ApplicationInterface) (*Client, error) {
	baseClient, err := kernel.NewBaseClient(&app, nil)
	if err != nil {
		return nil, err
	}
	return &Client{
		baseClient,
	}, nil
}

// 批量获取汇报记录单号
// https://developer.work.weixin.qq.com/document/path/93393
func (comp *Client) GetRecordList(ctx context.Context, startTime int, endTime int, nextCursor int, size int, filters []*power.StringMap) (*response.ResponseJournalGetRecordList, error) {

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

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/oa/journal/get_record_list", options, nil, nil, result)

	return result, err
}

// 获取汇报记录详情
// https://developer.work.weixin.qq.com/document/path/93394
func (comp *Client) GetRecordDetail(ctx context.Context, JournalUUID string) (*response.ResponseJournalGetRecordDetail, error) {

	result := &response.ResponseJournalGetRecordDetail{}

	options := &object.HashMap{
		"journaluuid": JournalUUID,
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/oa/journal/get_record_detail", options, nil, nil, result)

	return result, err
}

// 获取汇报记录详情
// https://developer.work.weixin.qq.com/document/path/93395
func (comp *Client) GetStatList(ctx context.Context, templateID string, startTime int, endTime int) (*response.ResponseJournalGetStatList, error) {

	result := &response.ResponseJournalGetStatList{}

	options := &object.HashMap{
		"template_id": templateID,
		"starttime":   fmt.Sprintf("%d", startTime),
		"endtime":     fmt.Sprintf("%d", endTime),
	}

	_, err := comp.BaseClient.HttpPostJson(ctx, "cgi-bin/oa/journal/get_stat_list", options, nil, nil, result)

	return result, err
}
