package externalcontact

import (
	"encoding/json"
	"fmt"

	"github.com/silenceper/wechat/v2/work/message"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// getFollowUserListURL 获取配置了客户联系功能的成员列表API地址
	getFollowUserListURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_follow_user_list"
	// addContactWayURL 配置客户联系「联系我」方式API地址
	addContactWayURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_contact_way"
	// getContactWayURL 获取企业已配置的「联系我」方式API地址
	getContactWayURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_contact_way"
	// updateContactWayURL 更新企业已配置的「联系我」方式API地址
	updateContactWayURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/update_contact_way"
	// delContactWayURL 删除企业已配置的「联系我」方式API地址
	delContactWayURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/del_contact_way"
	// closeTempChatURL 结束临时会话API地址
	closeTempChatURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/close_temp_chat"
)

// GetFollowUserListResp 获取配置了客户联系功能的成员列表返回
type GetFollowUserListResp struct {
	util.CommonError
	List []string `json:"follow_user"`
}

// GetFollowUserList 获取配置了客户联系功能的成员列表
func (ec *ExternalContact) GetFollowUserList() (ret []string, err error) {
	accessToken, err := ec.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", getFollowUserListURL, accessToken)
	resp, err := util.HTTPGet(url)
	if err != nil {
		return
	}

	var result GetFollowUserListResp
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("get follow user list error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	ret = result.List

	return
}

// Conclusion 结束语
type Conclusion struct {
	Text        *message.Text        `json:"text,omitempty"`
	Image       *message.Image       `json:"image,omitempty"`
	Link        *message.Link        `json:"link,omitempty"`
	Miniprogram *message.Miniprogram `json:"miniprogram,omitempty"`
}

// ContactWay 联系我方式参数
type ContactWay struct {
	ConfigID      string     `json:"config_id,omitempty"`
	Type          int        `json:"type"`
	Scene         int        `json:"scene"`
	Style         int        `json:"style,omitempty"`
	Remark        string     `json:"remark,omitempty"`
	SkipVerify    bool       `json:"skip_verify,omitempty"`
	State         string     `json:"state,omitempty"`
	QRCode        string     `json:"qr_code,omitempty"`
	User          []string   `json:"user,omitempty"`
	Party         []int      `json:"party,omitempty"`
	IsTemp        bool       `json:"is_temp,omitempty"`
	ExpiresIn     int        `json:"expires_in,omitempty"`
	ChatExpiresIn int        `json:"chat_expires_in,omitempty"`
	UnionID       string     `json:"unionid,omitempty"`
	Conclusions   Conclusion `json:"conclusions,omitempty"`
}

// AddContactWayResp 配置客户联系「联系我」方式返回结果
type AddContactWayResp struct {
	util.CommonError
	ConfigID string `json:"config_id"`
}

// AddContactWay 配置客户联系「联系我」方式
// 文档地址： https://work.weixin.qq.com/api/doc/90000/90135/92572
func (ec *ExternalContact) AddContactWay(params *ContactWay) (configID string, err error) {
	accessToken, err := ec.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", addContactWayURL, accessToken)
	resp, err := util.PostJSON(url, params)
	if err != nil {
		return
	}

	var result AddContactWayResp
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("add contact way error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	configID = result.ConfigID

	return
}

type GetContactWayResp struct {
	util.CommonError
	ContactWay ContactWay `json:"contact_way"`
}

// GetContactWay 获取企业已配置的「联系我」方式
// 文档地址：https://work.weixin.qq.com/api/doc/90000/90135/92572#%E8%8E%B7%E5%8F%96%E4%BC%81%E4%B8%9A%E5%B7%B2%E9%85%8D%E7%BD%AE%E7%9A%84%E3%80%8C%E8%81%94%E7%B3%BB%E6%88%91%E3%80%8D%E6%96%B9%E5%BC%8F
func (ec *ExternalContact) GetContactWay(configID string) (way ContactWay, err error) {
	accessToken, err := ec.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", getContactWayURL, accessToken)
	resp, err := util.PostJSON(url, map[string]string{
		"config_id": configID,
	})
	if err != nil {
		return
	}

	var result GetContactWayResp
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("get contact way error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	way = result.ContactWay

	return
}

// UpdateContactWay 更新企业已配置的「联系我」方式
// 文档地址：https://work.weixin.qq.com/api/doc/90000/90135/92572#%E6%9B%B4%E6%96%B0%E4%BC%81%E4%B8%9A%E5%B7%B2%E9%85%8D%E7%BD%AE%E7%9A%84%E3%80%8C%E8%81%94%E7%B3%BB%E6%88%91%E3%80%8D%E6%96%B9%E5%BC%8F
func (ec *ExternalContact) UpdateContactWay(params *ContactWay) (err error) {
	accessToken, err := ec.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", updateContactWayURL, accessToken)
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
		err = fmt.Errorf("update contact way error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	return
}

// DelContactWay 删除企业已配置的「联系我」方式
// 文档地址：https://work.weixin.qq.com/api/doc/90000/90135/92572#%E5%88%A0%E9%99%A4%E4%BC%81%E4%B8%9A%E5%B7%B2%E9%85%8D%E7%BD%AE%E7%9A%84%E3%80%8C%E8%81%94%E7%B3%BB%E6%88%91%E3%80%8D%E6%96%B9%E5%BC%8F
func (ec *ExternalContact) DelContactWay(configID string) (err error) {
	accessToken, err := ec.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", delContactWayURL, accessToken)
	resp, err := util.PostJSON(url, map[string]string{
		"config_id": configID,
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
		err = fmt.Errorf("del contact way error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	return
}

// CloseTempChat 结束临时会话
// 文档地址：https://work.weixin.qq.com/api/doc/90000/90135/92572#%E7%BB%93%E6%9D%9F%E4%B8%B4%E6%97%B6%E4%BC%9A%E8%AF%9D
func (ec *ExternalContact) CloseTempChat(userID, extUserID string) (err error) {
	accessToken, err := ec.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", closeTempChatURL, accessToken)
	resp, err := util.PostJSON(url, map[string]string{
		"userid":          userID,
		"external_userid": extUserID,
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
		err = fmt.Errorf("close temp chat error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	return
}
