package request

type RequestGroupRobotMsg struct {
	MsgType      string                    `json:"msgtype"`
	Text         GroupRobotMsgText         `json:"text,omitempty"`
	Markdown     GroupRobotMsgMarkdown     `json:"markdown,omitempty"`
	Image        GroupRobotMsgImage        `json:"image,omitempty"`
	News         GroupRobotMsgNews         `json:"news,omitempty"`
	File         GroupRobotMsgFile         `json:"file"`
	TemplateCard GroupRobotMsgTemplateCard `json:"template_card"`
}

type GroupRobotMsgText struct {
	Content             string   `json:"content"`
	MentionedList       []string `json:"mentioned_list,omitempty"`
	MentionedMobileList []string `json:"mentioned_mobile_list,omitempty"`
}

type GroupRobotMsgMarkdown struct {
	Content string `json:"content"`
}

type GroupRobotMsgImage struct {
	Base64 string `json:"base64"`
	Md5    string `json:"md5"`
}

type GroupRobotMsgNews struct {
	Articles []GroupRobotMsgNewsArticles `json:"articles"`
}

type GroupRobotMsgNewsArticles struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	Url         string `json:"url"`
	PicUrl      string `json:"picurl,omitempty"`
}

type GroupRobotMsgFile struct {
	MediaID string `json:"media_id"`
}

type GroupRobotMsgTemplateCard struct {
	CardType              string                          `json:"card_type"`
	Source                TemplateCardSource              `json:"source,omitempty"`
	MainTitle             TemplateCardMainTitle           `json:"main_title,omitempty"`
	CardImage             TemplateCardImage               `json:"card_image,omitempty"`
	VerticalContentList   []TemplateCardVerticalContentListItem       `json:"vertical_content_list"`
	EmphasisContent       TemplateCardEmphasisContent     `json:"emphasis_content,omitempty"`
	SubTitleText          string                          `json:"sub_title_text"`
	HorizontalContentList []TemplateCardHorizontalContentListItem     `json:"horizontal_content_list,omitempty"`
	JumpList              []TemplateCardJumpListItem `json:"jump_list,omitempty"`
	CardAction            TemplateCardCardAction                      `json:"card_action,omitempty"`
}
type TemplateCardSource struct {
	IconUrl string `json:"icon_url"`
	Desc    string `json:"desc"`
}
type TemplateCardMainTitle struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}
type TemplateCardImage struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}
type TemplateCardVerticalContentListItem struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}
type TemplateCardEmphasisContent struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}
type TemplateCardHorizontalContentListItem struct {
	KeyName string `json:"keyname"`
	Value   string `json:"value"`
	Type    int    `json:"type,omitempty"`
	Url     string `json:"url,omitempty"`
	MediaID string `json:"media_id,omitempty"`
}
type TemplateCardJumpListItem struct {
	Type     int    `json:"type"`
	Url      string `json:"url,omitempty"`
	Title    string `json:"title"`
	AppID    string `json:"appid,omitempty"`
	PagePath string `json:"pagepath,omitempty"`
}
type TemplateCardCardAction struct {
	Type     int    `json:"type"`
	Url      string `json:"url"`
	AppID    string `json:"appid"`
	PagePath string `json:"pagepath"`
}
