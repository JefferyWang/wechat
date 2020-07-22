package externalcontact

import (
	"encoding/json"
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// getGroupChatListURL 获取客户群列表API地址
	getGroupChatListURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/list"
	// getGroupChatURL 获取客户群详情API地址
	getGroupChatURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/get"
)

// OwnerFilter 群主过滤
type OwnerFilter struct {
	UserIDList  []string `json:"userid_list,omitempty"`
	PartyIDList []int    `json:"partyid_list,omitempty"`
}

// GetGroupChatListParams 获取客户群列表参数
type GetGroupChatListParams struct {
	StatusFilter int         `json:"status_filter,omitempty"`
	Offset       int         `json:"offset,omitempty"`
	OwnerFilter  OwnerFilter `json:"owner_filter,omitempty"`
	Limit        int         `json:"limit"`
}

// GroupChatItem 客户群信息
type GroupChatItem struct {
	ChatID string `json:"chat_id"`
	Status int    `json:"status"`
}

// GetGroupChatListResp 获取客户群列表返回结果
type GetGroupChatListResp struct {
	util.CommonError
	GroupChatList []GroupChatItem `json:"group_chat_list"`
}

// GetGroupChatList 获取客户群列表
// 文档地址： https://work.weixin.qq.com/api/doc/90000/90135/92120
func (ec *ExternalContact) GetGroupChatList(params *GetGroupChatListParams) (ret []GroupChatItem, err error) {
	accessToken, err := ec.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", getGroupChatListURL, accessToken)
	resp, err := util.PostJSON(url, params)
	if err != nil {
		return
	}

	var result GetGroupChatListResp
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("get group chat list error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	ret = result.GroupChatList

	return
}

type GroupChatMember struct {
	UserID    string `json:"userid"`            // 群成员id
	Type      int    `json:"type"`              // 成员类型，1-企业成员，2-外部联系人
	UnionID   string `json:"unionid,omitempty"` // 外部联系人在微信开放平台的唯一身份标识
	JoinTime  int64  `json:"join_time"`         // 入群时间
	JoinScene int    `json:"join_scene"`        // 入群方式，1-由成员邀请入群（直接邀请入群），2-由成员邀请入群（通过邀请链接入群），3-通过扫描群二维码入群
}

// GroupChat 客户群详情
type GroupChat struct {
	ChatID     string            `json:"chat_id"`
	Name       string            `json:"name"`
	Owner      string            `json:"owner"`
	CreateTime int64             `json:"create_time"`
	Notice     string            `json:"notice"`
	MemberList []GroupChatMember `json:"member_list"`
}

// GetGroupChatResp 获取客户群详情返回结果
type GetGroupChatResp struct {
	util.CommonError
	GroupChat GroupChat `json:"group_chat"`
}

// GetGroupChat 获取客户群详情
// 文档地址： https://work.weixin.qq.com/api/doc/90000/90135/92122
func (ec *ExternalContact) GetGroupChat(chatID string) (ret GroupChat, err error) {
	accessToken, err := ec.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", getGroupChatURL, accessToken)
	resp, err := util.PostJSON(url, map[string]string{
		"chat_id": chatID,
	})
	if err != nil {
		return
	}

	var result GetGroupChatResp
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("get group chat error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	ret = result.GroupChat

	return
}
