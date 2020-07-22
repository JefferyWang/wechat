package externalcontact

import (
	"encoding/json"
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// getCorpTagListURL 获取企业标签库API地址
	getCorpTagListURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get_corp_tag_list"
	// addCorpTagURL 添加企业客户标签API地址
	addCorpTagURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/add_corp_tag"
	// editCorpTagURL 编辑企业客户标签API地址
	editCorpTagURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/edit_corp_tag"
	// delCorpTagURL 删除企业客户标签API地址
	delCorpTagURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/del_corp_tag"
	// markTagURL 编辑客户企业标签API地址
	markTagURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/mark_tag"
)

// TagGroup 标签组
type TagGroup struct {
	GroupID    string `json:"group_id"`
	GroupName  string `json:"group_name"`
	CreateTime int64  `json:"create_time,omitempty"`
	Order      int    `json:"order"`
	Deleted    bool   `json:"deleted,omitempty"`
	Tags       []Tag  `json:"tag"`
}

// Tag 标签
type Tag struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name"`
	CreateTime int64  `json:"create_time,omitempty"`
	Order      int    `json:"order,omitempty"`
	Deleted    bool   `json:"deleted,omitempty"`
}

// GetCorpTagListResp 获取企业标签库返回结果
type GetCorpTagListResp struct {
	util.CommonError
	TagGroups []TagGroup `json:"tag_group"`
}

// GetCorpTagList 获取企业标签库
// 文档地址：https://work.weixin.qq.com/api/doc/90000/90135/92117
func (ec *ExternalContact) GetCorpTagList(tagIDs ...string) (ret []TagGroup, err error) {
	accessToken, err := ec.GetAccessToken()
	if err != nil {
		return
	}

	params := map[string]interface{}{}
	if len(tagIDs) > 0 {
		params["tag_id"] = tagIDs
	}

	url := fmt.Sprintf("%s?access_token=%s", getCorpTagListURL, accessToken)
	resp, err := util.PostJSON(url, params)
	if err != nil {
		return
	}

	var result GetCorpTagListResp
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("get corp tag list error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	ret = result.TagGroups

	return
}

// AddCorpTagResp 添加企业客户标签返回结果
type AddCorpTagResp struct {
	util.CommonError
	TagGroup TagGroup `json:"tag_group"`
}

// AddCorpTag 添加企业客户标签
// 文档地址： https://work.weixin.qq.com/api/doc/90000/90135/92117
func (ec *ExternalContact) AddCorpTag(params *TagGroup) (ret TagGroup, err error) {
	accessToken, err := ec.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", addCorpTagURL, accessToken)
	resp, err := util.PostJSON(url, params)
	if err != nil {
		return
	}

	var result AddCorpTagResp
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("add corp tag error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	ret = result.TagGroup

	return
}

// EditCorpTag 编辑企业客户标签
// 文档地址： https://work.weixin.qq.com/api/doc/90000/90135/92117
func (ec *ExternalContact) EditCorpTag(params *Tag) (err error) {
	accessToken, err := ec.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", editCorpTagURL, accessToken)
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
		err = fmt.Errorf("edit corp tag error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	return
}

// DelCorpTag 删除企业客户标签
// 文档地址： https://work.weixin.qq.com/api/doc/90000/90135/92117
func (ec *ExternalContact) DelCorpTag(tagIDs []string, groupIDs []string) (err error) {
	accessToken, err := ec.GetAccessToken()
	if err != nil {
		return
	}

	params := map[string]interface{}{
		"tag_id":   tagIDs,
		"group_id": groupIDs,
	}

	url := fmt.Sprintf("%s?access_token=%s", delCorpTagURL, accessToken)
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
		err = fmt.Errorf("del corp tag error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	return
}

// MarkTagParams 编辑客户企业标签参数
type MarkTagParams struct {
	UserID         string   `json:"userid"`
	ExternalUserID string   `json:"external_userid"`
	AddTag         []string `json:"add_tag,omitempty"`
	RemoveTag      []string `json:"remove_tag,omitempty"`
}

// MarkTag 编辑客户企业标签
// 文档地址： https://work.weixin.qq.com/api/doc/90000/90135/92118
func (ec *ExternalContact) MarkTag(params *MarkTagParams) (err error) {
	accessToken, err := ec.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", markTagURL, accessToken)
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
		err = fmt.Errorf("mark corp tag error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	return
}
