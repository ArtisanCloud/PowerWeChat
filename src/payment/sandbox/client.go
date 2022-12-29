package sandbox

import (
	"crypto/md5"
	"encoding/json"
	"errors"
	kernel2 "github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/kernel"
	"time"
)

type Client struct {
	*kernel.BaseClient

	*kernel2.InteractsWithCache
}

func NewClient(app *kernel.ApplicationPaymentInterface) (*Client, error) {
	baseClient, err := kernel.NewBaseClient(app)
	if err != nil {
		return nil, err
	}
	return &Client{
		baseClient,
		&kernel2.InteractsWithCache{},
	}, nil
}

func (comp *Client) GetKey() (string, error) {

	cache := comp.GetCache()
	key := comp.GetCacheKey()

	cacheKey, err := cache.Get(key, "")
	if cacheKey == nil || err != nil {
		return "", err
	}
	strCacheKey := cacheKey.(string)
	if strCacheKey != "" {
		return strCacheKey, nil
	}

	response, err := comp.RequestArray(nil, "sandboxnew/pay/getsignkey", "POST", nil, nil, nil)
	if err != nil {
		return "", err
	}
	var (
		returnCode     string
		sandboxSignKey string
	)
	if (*response)["return_code"] != nil {
		returnCode = (*response)["return_code"].(string)
	}
	if returnCode == "SUCCESS" {
		sandboxSignKey = (*response)["sandbox_signkey"].(string)
		err = cache.Set(key, sandboxSignKey, time.Hour)
		if err != nil {
			return "", err
		} else {
			return sandboxSignKey, nil
		}
	} else {
		returnMessage := "sandbox client get key error"
		if (*response)["retmsg"] != nil {
			returnMessage = (*response)["retmsg"].(string)
		} else if (*response)["return_msg"] != nil {
			returnMessage = (*response)["return_msg"].(string)
		}

		return "", errors.New(returnMessage)
	}

}

func (comp *Client) GetCacheKey() string {

	config := (*comp.App).GetConfig()
	strData := config.GetString("app_id", "") + config.GetString("mch_id", "")
	data, _ := json.Marshal(strData)
	buffer := md5.Sum(data)

	cacheKey := "powerwechat.payment.sandbox." + string(buffer[:])

	// tbf
	//fmt2.Dump(cacheKey)

	return cacheKey
}
