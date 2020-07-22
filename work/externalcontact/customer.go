package externalcontact

import (
	"encoding/json"
	"fmt"

	"github.com/silenceper/wechat/v2/work/contact"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// getExtUserListURL 获取客户列表API地址
	getExtUserListURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/list"
	// getExtUserDetailURL 获取客户详情API地址
	getExtUserDetailURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/get"
	// remarkExtUserURL 修改客户备注信息
	remarkExtUserURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/remark"
)

// GetExtUserListResp 获取客户列表返回结果
type GetExtUserListResp struct {
	util.CommonError
	ExternalUserID []string `json:"external_userid"`
}

// GetExtUserList 获取客户列表
// 文档地址：https://work.weixin.qq.com/api/doc/90000/90135/92113
func (ec *ExternalContact) GetExtUserList(userID string) (ret []string, err error) {
	accessToken, err := ec.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s&userid=%s", getExtUserListURL, accessToken, userID)
	resp, err := util.HTTPGet(url)
	if err != nil {
		return
	}

	var result GetExtUserListResp
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("get ext user list error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	ret = result.ExternalUserID

	return
}

// ExtContactTag 客户标签
type ExtContactTag struct {
	GroupName string `json:"group_name"` // 标签分组名称
	TagName   string `json:"tag_name"`   // 标签名称
	Type      int    `json:"type"`       // 标签类型，1-企业设置，2-用户自定义
}

// FollowUser 添加了此外部联系人的企业成员
type FollowUser struct {
	UserID         string          `json:"userid"`           // 客户id
	Remark         string          `json:"remark"`           // 客户备注
	Description    string          `json:"description"`      // 客户描述
	CreateTime     int64           `json:"createtime"`       // 添加客户时间
	Tags           []ExtContactTag `json:"tags"`             // 客户标签
	RemarkCorpName string          `json:"remark_corp_name"` // 备注的企业名称
	RemarkMobiles  []string        `json:"remark_mobiles"`   // 备注的手机号码
	State          string          `json:"state"`            // 添加此客户的渠道
	AddWay         string          `json:"add_way"`          // 该成员添加此客户的来源
	OperUserID     string          `json:"oper_userid"`      // 发起添加的userid，如果成员主动添加，为成员的userid；如果是客户主动添加，则为客户的外部联系人userid；如果是内部成员共享/管理员分配，则为对应的成员/管理员userid
}

// ExtContact 外部联系人信息
type ExtContact struct {
	ExternalUserID  string                  `json:"external_userid"`  // 外部联系人的userid
	Name            string                  `json:"name"`             // 外部联系人的名称
	Avatar          string                  `json:"avatar"`           // 外部联系人的头像
	Type            int                     `json:"type"`             // 外部联系人的类型，1-微信用户，2-企业微信用户
	Gender          int                     `json:"gender"`           // 外部联系人的性别，0-未知，1-男性，2-女性
	UnionID         string                  `json:"unionid"`          // 外部联系人在微信开放平台的唯一身份标识
	Position        string                  `json:"position"`         // 外部联系人的职位
	CorpName        string                  `json:"corp_name"`        // 外部联系人所在企业的简称
	CorpFullName    string                  `json:"corp_full_name"`   // 外部联系人所在企业的主体名称
	ExternalProfile contact.ExternalProfile `json:"external_profile"` // 外部联系人的自定义展示信息
}

// GetExtUserDetailResp 获取客户详情返回结果
type GetExtUserDetailResp struct {
	util.CommonError
	ExternalContact ExtContact `json:"external_contact"`
}

// GetExtUserDetail 获取客户详情
// 文档地址：https://work.weixin.qq.com/api/doc/90000/90135/92114
func (ec *ExternalContact) GetExtUserDetail(extUserID string) (ret ExtContact, err error) {
	accessToken, err := ec.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s&external_userid=%s", getExtUserDetailURL, accessToken, extUserID)
	resp, err := util.HTTPGet(url)
	if err != nil {
		return
	}

	var result GetExtUserDetailResp
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("get ext user detail error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	ret = result.ExternalContact

	return
}

type Remark struct {
	UserID           string   `json:"userid"`                       // 企业成员的userid
	ExternalUserID   string   `json:"external_userid"`              // 外部联系人userid
	Remark           string   `json:"remark,omitempty"`             // 此用户对外部联系人的备注
	Description      string   `json:"description,omitempty"`        // 此用户对外部联系人的描述
	RemarkCompany    string   `json:"remark_company,omitempty"`     // 此用户对外部联系人备注的所属公司名称
	RemarkMobiles    []string `json:"remark_mobiles,omitempty"`     // 此用户对外部联系人备注的手机号
	RemarkPicMediaID string   `json:"remark_pic_mediaid,omitempty"` // 备注图片的mediaid
}

// RemarkExtUser 修改客户备注信息
// 文档地址： https://work.weixin.qq.com/api/doc/90000/90135/92115
func (ec *ExternalContact) RemarkExtUser(params *Remark) (err error) {
	accessToken, err := ec.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", remarkExtUserURL, accessToken)
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
		err = fmt.Errorf("remark external user error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	return
}
