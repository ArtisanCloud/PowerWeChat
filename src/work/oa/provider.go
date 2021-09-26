package oa

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel"
	"github.com/ArtisanCloud/power-wechat/src/work/oa/calendar"
	"github.com/ArtisanCloud/power-wechat/src/work/oa/dial"
	"github.com/ArtisanCloud/power-wechat/src/work/oa/journal"
	"github.com/ArtisanCloud/power-wechat/src/work/oa/living"
	"github.com/ArtisanCloud/power-wechat/src/work/oa/meeting"
	"github.com/ArtisanCloud/power-wechat/src/work/oa/meetingroom"
	"github.com/ArtisanCloud/power-wechat/src/work/oa/pstncc"
	"github.com/ArtisanCloud/power-wechat/src/work/oa/schedule"
	"github.com/ArtisanCloud/power-wechat/src/work/oa/webdrive"
)

func RegisterProvider(app kernel.ApplicationInterface) (*Client,
	*calendar.Client,
	*dial.Client,
	*journal.Client,
	*living.Client,
	*meeting.Client,
	*meetingroom.Client,
	*pstncc.Client,
	*schedule.Client,
	*webdrive.Client,

) {
	//config := app.GetConfig()

	Client := NewClient(app)
	Calendar := calendar.NewClient(app)
	Dial := dial.NewClient(app)
	Journal := journal.NewClient(app)
	Living := living.NewClient(app)
	Meeting := meeting.NewClient(app)
	MeetingRoom := meetingroom.NewClient(app)
	PSTNCC := pstncc.NewClient(app)
	Schedule := schedule.NewClient(app)
	WebDrive := webdrive.NewClient(app)

	return Client,
		Calendar,
		Dial,
		Journal,
		Living,
		Meeting,
		MeetingRoom,
		PSTNCC,
		Schedule,
		WebDrive

}
