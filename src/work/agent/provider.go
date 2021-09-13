package agent

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) (*Client, *WorkbenchClient) {

	client := NewClient(app)

	workbenchClient := NewWorkbenchClient(app)

	return client, workbenchClient

}
