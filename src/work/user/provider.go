package user

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/user/batchJobs"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/user/exportJobs"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/user/linkedCorp"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/work/user/tag"
)

func RegisterProvider(app kernel.ApplicationInterface) (
	*Client,
	*batchJobs.Client,
	*exportJobs.Client,
	*linkedCorp.Client,
	*tag.Client,
	error,
) {
	//config := app.GetConfig()

	client, err := NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	UserBatchJobs, err := batchJobs.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	UserExportJobs, err := exportJobs.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	UserLinkedCorp, err := linkedCorp.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	UserTag, err := tag.NewClient(app)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}
	return client,
		UserBatchJobs,
		UserExportJobs,
		UserLinkedCorp,
		UserTag,
		nil
}
