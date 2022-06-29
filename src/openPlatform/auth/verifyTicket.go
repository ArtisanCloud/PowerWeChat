package auth

import (
	"errors"
	"github.com/ArtisanCloud/PowerLibs/v2/cache"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
)

type VerifyTicket struct {
	App *kernel.ApplicationInterface

	*kernel.InteractsWithCache
}

func NewVerifyTicket(app *kernel.ApplicationInterface) (*VerifyTicket, error) {
	config := (*app).GetContainer().GetConfig()

	var cacheClient cache.CacheInterface = nil
	if (*config)["cache"] != nil {
		cacheClient = (*config)["cache"].(cache.CacheInterface)
	}

	ticket := &VerifyTicket{
		App: app,

		InteractsWithCache: kernel.NewInteractsWithCache(cacheClient),
	}

	return ticket, nil
}

func (verifyTicket *VerifyTicket) SetTicket(ticket string) error {
	verifyTicket.GetCache().Set(verifyTicket.getCacheKey(), ticket, 3600)

	if !verifyTicket.GetCache().Has(verifyTicket.getCacheKey()) {
		return errors.New("failed to cache verify ticket")
	}

	return nil
}

func (verifyTicket *VerifyTicket) GetTicket() (ticket string, err error) {

	cached, err := verifyTicket.GetCache().Get(verifyTicket.getCacheKey(), "")
	if err != nil {
		return "", err
	}

	ticket = cached.(string)

	return ticket, nil
}

func (verifyTicket *VerifyTicket) getCacheKey() string {
	config := (*verifyTicket.App).GetConfig()
	return "powerwechat.open_platform.verify_ticket." + config.GetString("app_id", "")

}
