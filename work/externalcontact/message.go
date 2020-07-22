package externalcontact

import (
	"encoding/json"
	"fmt"

	"github.com/silenceper/wechat/v2/util"
	"github.com/silenceper/wechat/v2/work/message"
)

const (
	// addMsgTemplateURL 添加企业群发消息任务API地址
	addMsgTemplateURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_msg_template"
	// getGroupMsgResultURL 获取企业群发消息发送结果API地址
	getGroupMsgResultURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_group_msg_result"
	// sendWelcomeMsgURL 发送新客户欢迎语API地址
	sendWelcomeMsgURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/send_welcome_msg"
	// addGroupWelcomeTemplateURL 添加群欢迎语素材API地址
	addGroupWelcomeTemplateURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/group_welcome_template/add"
	// editGroupWelcomeTemplateURL 编辑群欢迎语素材
	editGroupWelcomeTemplateURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/group_welcome_template/edit"
	// getGroupWelcomeTemplateURL 获取群欢迎语素材API地址
	getGroupWelcomeTemplateURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/group_welcome_template/get"
	// delGroupWelcomeTemplateURL 删除群欢迎语素材API地址
	delGroupWelcomeTemplateURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/group_welcome_template/del"
)

// AddMsgTemplateParams 添加企业群发消息任务参数
type AddMsgTemplateParams struct {
	ChatType       string               `json:"chat_type,omitempty"` // 群发任务的类型
	ExternalUserID []string             `json:"external_userid,omitempty"`
	Sender         string               `json:"sender,omitempty"`
	Text           *message.Text        `json:"text,omitempty"`
	Image          *message.Image       `json:"image,omitempty"`
	Link           *message.Link        `json:"link,omitempty"`
	Miniprogram    *message.Miniprogram `json:"miniprogram,omitempty"`
}

// AddMsgTemplateResp 添加企业群发消息任务返回结果
type AddMsgTemplateResp struct {
	util.CommonError
	FailList []string `json:"fail_list"` // 无效或无法发送的external_userid列表
	MsgID    string   `json:"msgid"`     // 企业群发消息的id
}

// AddMsgTemplate 添加企业群发消息任务
// 文档地址： https://work.weixin.qq.com/api/doc/90000/90135/92135
func (ec *ExternalContact) AddMsgTemplate(params *AddMsgTemplateParams) (result AddMsgTemplateResp, err error) {
	accessToken, err := ec.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", addMsgTemplateURL, accessToken)
	resp, err := util.PostJSON(url, params)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("add msg template error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	return
}

// GroupMsgResultItem 企业群发消息发送结果
type GroupMsgResultItem struct {
	ExternalUserID string `json:"external_userid"`
	ChatID         string `json:"chat_id"`
	UserID         string `json:"userid"`
	Status         int    `json:"status"`
	SendTime       int64  `json:"send_time"`
}

// GetGroupMsgResultResp 获取企业群发消息发送结果返回结果
type GetGroupMsgResultResp struct {
	util.CommonError
	DetailList []GroupMsgResultItem `json:"detail_list"`
}

// GetGroupMsgResult 获取企业群发消息发送结果
// 文档地址：https://work.weixin.qq.com/api/doc/90000/90135/92136
func (ec *ExternalContact) GetGroupMsgResult(msgID string) (ret []GroupMsgResultItem, err error) {
	accessToken, err := ec.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", getGroupMsgResultURL, accessToken)
	resp, err := util.PostJSON(url, map[string]string{
		"msgid": msgID,
	})
	if err != nil {
		return
	}

	var result GetGroupMsgResultResp
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("get group msg result error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}
	ret = result.DetailList

	return
}

// SendWelcomeMsgParams 发送新客户欢迎语参数
type SendWelcomeMsgParams struct {
	WelcomeCode string               `json:"welcome_code"`
	Text        *message.Text        `json:"text,omitempty"`
	Image       *message.Image       `json:"image,omitempty"`
	Link        *message.Link        `json:"link,omitempty"`
	Miniprogram *message.Miniprogram `json:"miniprogram,omitempty"`
}

// SendWelcomeMsg 发送新客户欢迎语
// 文档地址：https://work.weixin.qq.com/api/doc/90000/90135/92137
func (ec *ExternalContact) SendWelcomeMsg(params *SendWelcomeMsgParams) (err error) {
	accessToken, err := ec.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", sendWelcomeMsgURL, accessToken)
	resp, err := util.PostJSON(url, params)
	if err != nil {
		return
	}

	var result util.CommonError
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("send welcome msg error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	return
}

// GroupWelcomeTemplate 群欢迎语素材
type GroupWelcomeTemplate struct {
	TemplateID  string               `json:"template_id,omitempty"`
	Text        *message.Text        `json:"text,omitempty"`
	Image       *message.Image       `json:"image,omitempty"`
	Link        *message.Link        `json:"link,omitempty"`
	Miniprogram *message.Miniprogram `json:"miniprogram,omitempty"`
}

// AddGroupWelcomeTemplateResp 添加群欢迎语素材返回结果
type AddGroupWelcomeTemplateResp struct {
	util.CommonError
	TemplateID string `json:"template_id"`
}

// AddGroupWelcomeTemplate 添加群欢迎语素材
// 文档地址：https://work.weixin.qq.com/api/doc/90000/90135/92366
func (ec *ExternalContact) AddGroupWelcomeTemplate(params *GroupWelcomeTemplate) (templateID string, err error) {
	accessToken, err := ec.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", addGroupWelcomeTemplateURL, accessToken)
	resp, err := util.PostJSON(url, params)
	if err != nil {
		return
	}

	var result AddGroupWelcomeTemplateResp
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("add group welcome template error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	templateID = result.TemplateID

	return
}

// EditGroupWelcomeTemplate 编辑群欢迎语素材
// 文档地址：https://work.weixin.qq.com/api/doc/90000/90135/92366
func (ec *ExternalContact) EditGroupWelcomeTemplate(params *GroupWelcomeTemplate) (err error) {
	accessToken, err := ec.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", editGroupWelcomeTemplateURL, accessToken)
	resp, err := util.PostJSON(url, params)
	if err != nil {
		return
	}

	var result util.CommonError
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("edit group welcome template error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	return
}

// GetGroupWelcomeTemplateResp 获取群欢迎语素材返回结果
type GetGroupWelcomeTemplateResp struct {
	util.CommonError
	GroupWelcomeTemplate
}

// GetGroupWelcomeTemplate 获取群欢迎语素材
// 文档地址：https://work.weixin.qq.com/api/doc/90000/90135/92366
func (ec *ExternalContact) GetGroupWelcomeTemplate(templateID string) (template GroupWelcomeTemplate, err error) {
	accessToken, err := ec.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", getGroupWelcomeTemplateURL, accessToken)
	resp, err := util.PostJSON(url, map[string]string{
		"template_id": templateID,
	})
	if err != nil {
		return
	}

	var result GetGroupWelcomeTemplateResp
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("get group welcome template error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	template = result.GroupWelcomeTemplate

	return
}

// DelGroupWelcomeTemplate 删除群欢迎语素材
// 文档地址：https://work.weixin.qq.com/api/doc/90000/90135/92366
func (ec *ExternalContact) DelGroupWelcomeTemplate(templateID string) (err error) {
	accessToken, err := ec.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", delGroupWelcomeTemplateURL, accessToken)
	resp, err := util.PostJSON(url, map[string]string{
		"template_id": templateID,
	})
	if err != nil {
		return
	}

	var result util.CommonError
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("get group welcome template error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	return
}
