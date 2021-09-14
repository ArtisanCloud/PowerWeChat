package user

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) (
	*Client,
	*BatchJobs,
	*ExportJobs,
	*LinkedCorp,
	*TagClient,
) {
	//config := app.GetConfig()

	client := NewClient(app)

	UserBatchJobs := NewBatchJobs(app)
	UserExportJobs := NewExportJobs(app)
	UserLinkedCorp := NewLinkedCorp(app)
	UserTag := NewTagClient(app)

	return client,
		UserBatchJobs,
		UserExportJobs,
		UserLinkedCorp,
		UserTag
}
