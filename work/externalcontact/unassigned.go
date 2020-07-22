package externalcontact

import (
	"encoding/json"
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// getUnassignedListURL 获取离职成员的客户列表接口地址
	getUnassignedListURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_unassigned_list"
)

// UnassignedItem 离职成员
type UnassignedItem struct {
	HandoverUserID string `json:"handover_userid"`
	ExternalUserID string `json:"external_userid"`
	DimissionTime  int64  `json:"dimission_time"`
}

// GetUnassignedListResp 获取离职成员的客户列表返回结果
type GetUnassignedListResp struct {
	util.CommonError
	Info   []UnassignedItem `json:"info"`
	IsLast bool             `json:"is_last"`
}

// GetUnassignedList 获取离职成员的客户列表
// 文档地址：https://work.weixin.qq.com/api/doc/90000/90135/92124
func (ec *ExternalContact) GetUnassignedList(pageID, pageSize int) (result GetUnassignedListResp, err error) {
	accessToken, err := ec.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", getUnassignedListURL, accessToken)
	resp, err := util.PostJSON(url, map[string]int{
		"page_id":   pageID,
		"page_size": pageSize,
	})
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("get unassigned list error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	return
}
