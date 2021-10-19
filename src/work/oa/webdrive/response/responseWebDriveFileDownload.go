package response

import (
	"github.com/ArtisanCloud/powerwechat/src/kernel/response"
)

type ResponseWebDriveFileDownload struct {
	*response.ResponseWork

	DownloadURL string `json:"download_url"`
	CookieName  string `json:"cookie_name"`
	CookieValue string `json:"cookie_value"`
}
