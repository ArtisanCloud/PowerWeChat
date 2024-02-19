package handlers

import (
	"net/http"

	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/contract"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/server/models"
	suit "github.com/ArtisanCloud/PowerWeChat/v3/src/openWork/suitAuth"
)

type SuiteTicket struct {
	contract.EventHandlerInterface

	App *kernel.ApplicationInterface
}

func NewSuiteTicket(app *kernel.ApplicationInterface) *SuiteTicket {
	handler := &SuiteTicket{
		App: app,
	}
	return handler
}

func (handler *SuiteTicket) Handle(request *http.Request, header contract.EventInterface, content interface{}) interface{} {

	ev := content.(*models.EventSuiteTicket)
	ticket := ev.SuiteTicket

	if ticket != "" {
		(*handler.App).GetComponent("SuiteTicket").(*suit.SuiteTicket).SetTicket(ticket)
	}

	return nil
}
