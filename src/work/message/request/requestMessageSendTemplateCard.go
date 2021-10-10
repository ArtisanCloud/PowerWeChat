package request

type RequestMessageSendTemplateCard struct {
	RequestMessageSend
	TemplateCard *RequestTemplateCard `json:"template_card"`
}

type TemplateCardSource struct {
	IconUrl   string `json:"icon_url"`
	Desc      string `json:"desc"`
	DescColor int    `json:"desc_color"`
}

type TemplateCardActionMenu struct {
	Desc       string                        `json:"desc"`
	ActionList []*TemplateCardActionListItem `json:"action_list"`
}
type TemplateCardActionListItem struct {
	Text string `json:"text"`
	Key  string `json:"key"`
}

type TemplateCardMainTitle struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

type TemplateCardQuoteArea struct {
	Type      int    `json:"type"`
	Url       string `json:"url"`
	Title     string `json:"title"`
	QuoteText string `json:"quote_text"`
}

type TemplateCardEmphasisContent struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

type TemplateCardHorizontalContentListItem struct {
	Keyname string `json:"keyname"`
	Value   string `json:"value"`
	Type    int    `json:"type,omitempty"`
	Url     string `json:"url,omitempty"`
	MediaID string `json:"media_id,omitempty"`
	UserID  string `json:"userid,omitempty"`
}

type TemplateCardJumpListItem struct {
	Type     int    `json:"type"`
	Title    string `json:"title"`
	Url      string `json:"url,omitempty"`
	AppID    string `json:"appid,omitempty"`
	Pagepath string `json:"pagepath,omitempty"`
}

type TemplateCardAction struct {
	Type     int    `json:"type"`
	Url      string `json:"url"`
	AppID    string `json:"appid"`
	Pagepath string `json:"pagepath"`
}

type TemplateCardImageTextArea struct {
	Type     int    `json:"type"`
	Url      string `json:"url"`
	Title    string `json:"title"`
	Desc     string `json:"desc"`
	ImageUrl string `json:"image_url"`
}

type TemplateCardImage struct {
	Url         string  `json:"url"`
	AspectRatio float64 `json:"aspect_ratio"`
}

type TemplateCardVerticalContentListItem struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

type TemplateCardButtonSelection struct {
	QuestionKey string                        `json:"question_key"`
	Title       string                        `json:"title"`
	OptionList  []*TemplateCardOptionListItem `json:"option_list"`
	SelectedID  string                        `json:"selected_id"`
}
type TemplateCardOptionListItem struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}
type TemplateCardButtonListItem struct {
	Text  string `json:"text"`
	Style int    `json:"style"`
	Key   string `json:"key"`
}
type TemplateCardCheckbox struct {
	QuestionKey string                                `json:"question_key"`
	OptionList  []*TemplateCardCheckboxOptionListItem `json:"option_list"`
	Mode        int                                   `json:"mode"`
}
type TemplateCardCheckboxOptionListItem struct {
	ID        string `json:"id"`
	Text      string `json:"text"`
	IsChecked bool   `json:"is_checked"`
}
type TemplateCardSubmitButton struct {
	Text string `json:"text"`
	Key  string `json:"key"`
}
type TemplateCardSelectList struct {
	QuestionKey string                        `json:"question_key"`
	Title       string                        `json:"title"`
	SelectedID  string                        `json:"selected_id"`
	OptionList  []*TemplateCardOptionListItem `json:"option_list"`
}

type RequestTemplateCard struct {
	CardType              string                                   `json:"card_type"`
	Source                *TemplateCardSource                      `json:"source"`
	ActionMenu            *TemplateCardActionMenu                  `json:"action_menu"`
	TaskID                string                                   `json:"task_id"`
	MainTitle             *TemplateCardMainTitle                   `json:"main_title"`
	QuoteArea             *TemplateCardQuoteArea                   `json:"quote_area"`
	EmphasisContent       *TemplateCardEmphasisContent             `json:"emphasis_content"`
	SubTitleText          string                                   `json:"sub_title_text"`
	HorizontalContentList []*TemplateCardHorizontalContentListItem `json:"horizontal_content_list"`
	JumpList              []*TemplateCardJumpListItem              `json:"jump_list"`
	CardAction            *TemplateCardAction                      `json:"card_action"`
	ImageTextArea         *TemplateCardImageTextArea               `json:"image_text_area"`
	CardImage             *TemplateCardImage                       `json:"card_image"`
	VerticalContentList   []*TemplateCardVerticalContentListItem   `json:"vertical_content_list"`
	ButtonSelection       *TemplateCardButtonSelection             `json:"button_selection"`
	ButtonList            []*TemplateCardButtonListItem            `json:"button_list"`
	Checkbox              *TemplateCardCheckbox                    `json:"checkbox"`
	SubmitButton          *TemplateCardSubmitButton                `json:"submit_button"`
	SelectList            []*TemplateCardSelectList                `json:"select_list"`
}
