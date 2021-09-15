package user

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/work/user/batchJobs"
	"github.com/ArtisanCloud/power-wechat/src/work/user/exportJobs"
	"github.com/ArtisanCloud/power-wechat/src/work/user/linkedCorp"
	"github.com/ArtisanCloud/power-wechat/src/work/user/tag"
)

func RegisterProvider(app kernel.ApplicationInterface) (
	*Client,
	*batchJobs.Client,
	*exportJobs.Client,
	*linkedCorp.Client,
	*tag.Client,
) {
	//config := app.GetConfig()

	client := NewClient(app)

	UserBatchJobs := batchJobs.NewClient(app)
	UserExportJobs := exportJobs.NewClient(app)
	UserLinkedCorp := linkedCorp.NewClient(app)
	UserTag := tag.NewClient(app)

	return client,
		UserBatchJobs,
		UserExportJobs,
		UserLinkedCorp,
		UserTag
}
