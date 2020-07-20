package contact

import (
	"encoding/json"
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// createTagURL 创建标签接口地址
	createTagURL = "https://qyapi.weixin.qq.com/cgi-bin/tag/create"
	// updateTagURL 更新标签接口地址
	updateTagURL = "https://qyapi.weixin.qq.com/cgi-bin/tag/update"
	// deleteTagURL 删除标签接口地址
	deleteTagURL = "https://qyapi.weixin.qq.com/cgi-bin/tag/delete"
	// getTagUsersURL 获取标签成员接口地址
	getTagUsersURL = "https://qyapi.weixin.qq.com/cgi-bin/tag/get"
	// addTagUsersURL 增加标签成员接口地址
	addTagUsersURL = "https://qyapi.weixin.qq.com/cgi-bin/tag/addtagusers"
	// delTagUsersURL 删除标签成员接口地址
	delTagUsersURL = "https://qyapi.weixin.qq.com/cgi-bin/tag/deltagusers"
	// getTagListURL 获取标签列表接口地址
	getTagListURL = "https://qyapi.weixin.qq.com/cgi-bin/tag/list"
)

// Tag 员工标签
type Tag struct {
	ID   int    `json:"tagid,omitempty"` // 标签id
	Name string `json:"tagname"`         // 标签名字
}

// CreateDepartmentResp 创建部门返回结果
type CreateTagResp struct {
	util.CommonError
	ID int `json:"tagid"` // 创建的标签id
}

// CreateTag 创建标签
// 文档地址：https://work.weixin.qq.com/api/doc/90000/90135/90210
func (contact *Contact) CreateTag(tag *Tag) (ID int, err error) {
	accessToken, err := contact.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", createTagURL, accessToken)
	resp, err := util.PostJSON(url, tag)
	if err != nil {
		return
	}

	var result CreateTagResp
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("create tag error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	ID = result.ID

	return
}

// UpdateTag 更新标签名字
// 文档地址：https://work.weixin.qq.com/api/doc/90000/90135/90211
func (contact *Contact) UpdateTag(tag *Tag) (err error) {
	accessToken, err := contact.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", updateTagURL, accessToken)
	resp, err := util.PostJSON(url, tag)
	if err != nil {
		return
	}

	var result util.CommonError
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("update tag error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	return
}

// DeleteTag 删除标签
// 文档地址：https://work.weixin.qq.com/api/doc/90000/90135/90212
func (contact *Contact) DeleteTag(tagID int) (err error) {
	accessToken, err := contact.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s&tagid=%d", deleteTagURL, accessToken, tagID)
	resp, err := util.HTTPGet(url)
	if err != nil {
		return
	}

	var result util.CommonError
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("delete tag error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	return
}

// TinyUser 简易用户信息
type TinyUser struct {
	UserID string `json:"userid"` // 成员userid
	Name   string `json:"name"`   // 成员名称
}

// GetTagUsersResp 获取标签成员返回
type GetTagUsersResp struct {
	util.CommonError
	TagName   string     `json:"tagname"`   // 标签名
	UserList  []TinyUser `json:"userlist"`  // 标签中包含的成员列表
	PartyList []int      `json:"partylist"` // 标签中包含的部门id列表
}

// GetTagUsers 获取标签成员
// 文档地址：https://work.weixin.qq.com/api/doc/90000/90135/90213
func (contact *Contact) GetTagUsers(tagID int) (result GetTagUsersResp, err error) {
	accessToken, err := contact.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s&tagid=%d", getTagUsersURL, accessToken, tagID)
	resp, err := util.HTTPGet(url)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("get tag users error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	return
}

// TagUserParams 修改标签成员参数
type TagUserParams struct {
	TagID     int      `json:"tagid"`               // 标签id
	UserList  []string `json:"userlist,omitempty"`  // 标签中包含的成员列表
	PartyList []int    `json:"partylist,omitempty"` // 标签中包含的部门id列表
}

type UpdateTagUsersResp struct {
	util.CommonError
	InvalidList  string `json:"invalidlist,omitempty"`  // 非法的成员帐号列表
	InvalidParty []int  `json:"invalidparty,omitempty"` // 非法的部门id列表
}

// AddTagUsers 增加标签成员
// 文档地址：https://work.weixin.qq.com/api/doc/90000/90135/90214
func (contact *Contact) AddTagUsers(params *TagUserParams) (result UpdateTagUsersResp, err error) {
	accessToken, err := contact.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", addTagUsersURL, accessToken)
	resp, err := util.PostJSON(url, params)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("add tag users error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	return
}

// DelTagUsers 删除标签成员
// 文档地址：https://work.weixin.qq.com/api/doc/90000/90135/90215
func (contact *Contact) DelTagUsers(params *TagUserParams) (result UpdateTagUsersResp, err error) {
	accessToken, err := contact.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", delTagUsersURL, accessToken)
	resp, err := util.PostJSON(url, params)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("del tag users error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	return
}

// GetTagListResp 获取标签列表返回
type GetTagListResp struct {
	util.CommonError
	TagList []Tag `json:"taglist"`
}

// GetTagList 获取标签列表
// 文档地址：https://work.weixin.qq.com/api/doc/90000/90135/90216
func (contact *Contact) GetTagList() (ret []Tag, err error) {
	accessToken, err := contact.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", getTagListURL, accessToken)
	resp, err := util.HTTPGet(url)
	if err != nil {
		return
	}

	var result GetTagListResp
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("get tag list error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	ret = result.TagList

	return
}
