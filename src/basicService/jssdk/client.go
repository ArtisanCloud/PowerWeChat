package jssdk

import (
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ArtisanCloud/PowerLibs/http/response"
	"github.com/ArtisanCloud/PowerLibs/object"
	"github.com/ArtisanCloud/PowerWeChat/src/kernel"
	response2 "github.com/ArtisanCloud/PowerWeChat/src/kernel/response"
	"net/http"
	"sort"
	"strings"
	"time"
)

type Client struct {
	*kernel.BaseClient

	*kernel.InteractsWithCache

	TicketEndpoint string
	url            string
}

func NewClient(app *kernel.ApplicationInterface) *Client {
	client := &Client{
		BaseClient:         kernel.NewBaseClient(app, nil),
		InteractsWithCache: &kernel.InteractsWithCache{},
	}

	client.TicketEndpoint = "https://api.weixin.qq.com/cgi-bin/ticket/getticket"

	return client
}

func (comp *Client) BuildConfig(request *http.Request, jsApiList []string, debug bool, beta bool, isJson bool, openTagList []string, url string) (interface{}, error) {

	signature, err := comp.ConfigSignature(request, url, "", time.Time{})
	if err != nil {
		return signature, err
	}
	config := object.MergeHashMap(&object.HashMap{
		"debug":       debug,
		"beta":        beta,
		"jsApiList":   jsApiList,
		"openTagList": openTagList,
	}, signature)

	if isJson {
		return json.Marshal(config)
	} else {
		return config, nil
	}

}

func (comp *Client) GetConfigArray(request *http.Request, apis []string, debug bool, beta bool, openTagList []string, url string) (string, error) {

	result, err := comp.BuildConfig(request, apis, debug, beta, false, openTagList, url)

	return result.(string), err
}

func (comp *Client) GetTicket(refresh bool, ticketType string) (*object.HashMap, error) {

	cacheKey := fmt.Sprintf("powerwechat.basic_service.jssdk.ticket.%s.%s", ticketType, comp.GetAppID())

	if !refresh && comp.GetCache().Has(cacheKey) {
		ticket, err := comp.GetCache().Get(cacheKey, nil)
		return ticket.(*object.HashMap), err
	}

	mapRSBody := &object.HashMap{}
	resultBody := ""
	rs, err := comp.RequestRaw(comp.TicketEndpoint, "GET", &object.HashMap{
		"query": &object.StringMap{
			"type": ticketType,
		}}, nil, &resultBody)
	if err != nil {
		return nil, err
	}

	err = object.JsonDecode([]byte(resultBody), mapRSBody)
	if (*mapRSBody)["errcode"].(float64) != 0 {
		return mapRSBody, errors.New((*mapRSBody)["errmsg"].(string))
	}

	result, err := comp.CastResponseToType(rs.(*response.HttpResponse).Response, response2.TYPE_MAP)
	if err != nil {
		return nil, err
	}

	resultData := result.(*object.HashMap)
	err = comp.GetCache().Set(cacheKey, result, time.Duration((*resultData)["expires_in"].(float64)-500)*time.Second)
	if err != nil {
		return nil, err
	}

	if !comp.GetCache().Has(cacheKey) {
		return nil, errors.New("Failed to cache jssdk ticket. ")
	}

	return resultData, nil
}

func (comp *Client) ConfigSignature(request *http.Request, url string, nonce string, timestamp time.Time) (*object.HashMap, error) {

	if url != "" {
		url = comp.GetUrl(request)
	}
	if nonce != "" {
		nonce = object.QuickRandom(10)
	}
	if timestamp.IsZero() {
		timestamp = time.Now()
	}

	result, err := comp.GetTicket(false, "jsapi")
	if err != nil {
		return result, err
	}
	ticket := (*result)["ticket"].(string)

	return &object.HashMap{
		"appId":     comp.getAgentID(),
		"nonceStr":  nonce,
		"timestamp": timestamp,
		"url":       url,
		"signature": comp.GetTicketSignature(ticket, nonce, timestamp, url),
	}, nil

}

func (comp *Client) GetTicketSignature(ticket string, nonce string, timestamp time.Time, url string) string {

	param := fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%s&url=%s", ticket, nonce, timestamp, url)

	hash := sha1.New()
	hash.Write([]byte(param))
	return string(hash.Sum(nil))
}

func (comp *Client) dictionaryOrderSignature(params []string) string {
	sort.Strings(params)
	strJoined := strings.Join(params, "")
	hash := sha1.New()
	hash.Write([]byte(strJoined))
	return string(hash.Sum(nil))
}

func (comp *Client) SetUrl(url string) *Client {

	comp.url = url

	return comp
}

func (comp *Client) GetUrl(externalRequest *http.Request) string {

	if comp.url != "" {
		return comp.url
	}
	return externalRequest.URL.String()
}

func (comp *Client) GetAppID() string {
	config := (*comp.App).GetConfig()
	return config.GetString("app_id", "")
}

func (comp *Client) getAgentID() string {
	config := (*comp.App).GetConfig()
	return config.GetString("agent_id", "")
}
