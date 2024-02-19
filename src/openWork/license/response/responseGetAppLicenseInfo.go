package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/license/model"
)

type ResponseGetAppLicenseInfo struct {
	response.ResponseWork
	model.LicenseInfo
}
