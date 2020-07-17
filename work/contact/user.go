package contact

import (
	"encoding/json"
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// createUserURL 创建成员API地址
	createUserURL = "https://qyapi.weixin.qq.com/cgi-bin/user/create"
	// getUserURL 读取成员信息API地址
	getUserURL = "https://qyapi.weixin.qq.com/cgi-bin/user/get"
	// updateUserURL 更新成员信息API地址
	updateUserURL = "https://qyapi.weixin.qq.com/cgi-bin/user/update"
	// deleteUserURL 删除成员API地址
	deleteUserURL = "https://qyapi.weixin.qq.com/cgi-bin/user/delete"
	// batchDeleteUserURL 批量删除成员API地址
	batchDeleteUserURL = "https://qyapi.weixin.qq.com/cgi-bin/user/batchdelete"
	// getDeptUsersURL 获取部门成员API地址
	getDeptSimpleUsersURL = "https://qyapi.weixin.qq.com/cgi-bin/user/simplelist"
	// getDeptUsersURL 获取部门成员详情API地址
	getDeptUsersURL = "https://qyapi.weixin.qq.com/cgi-bin/user/list"
	// convertToOpenIDURL 将userid转换为openid的API地址
	convertToOpenIDURL = "https://qyapi.weixin.qq.com/cgi-bin/user/convert_to_openid"
	// convertToUserIDURL 将openid转换为userid的API地址
	convertToUserIDURL = "https://qyapi.weixin.qq.com/cgi-bin/user/convert_to_userid"
	// secondAuthURL 二次验证API地址
	secondAuthURL = "https://qyapi.weixin.qq.com/cgi-bin/user/authsucc"
	// batchInviteURL 邀请成员API地址
	batchInviteURL = "https://qyapi.weixin.qq.com/cgi-bin/batch/invite"
	// getJoinQRCodeURL 获取加入企业二维码API地址
	getJoinQRCodeURL = "https://qyapi.weixin.qq.com/cgi-bin/corp/get_join_qrcode"
	// getActiveStatURL 获取企业活跃成员数
	getActiveStatURL = "https://qyapi.weixin.qq.com/cgi-bin/user/get_active_stat"
)

// Attr 扩展属性
type Attr struct {
	Type        int              `json:"type"`
	Name        string           `json:"name"`
	Text        *TextAttr        `json:"text,omitempty"`
	Web         *WebAttr         `json:"web,omitempty"`
	Miniprogram *MiniprogramAttr `json:"miniprogram,omitempty"`
}

// TextAttr 文本类型属性
type TextAttr struct {
	Value string `json:"value"` // 文本属性的内容
}

// WebAttr 网页类型属性
type WebAttr struct {
	URL   string `json:"url"`   // 网页的url
	Title string `json:"title"` // 网页的展示标题
}

// MiniprogramAttr 小程序类型属性
type MiniprogramAttr struct {
	AppID    string `json:"appid"`    // 小程序appid
	PagePath string `json:"pagepath"` // 小程序页面路径
	Title    string `json:"title"`    // 小程序展示的标题
}

// ExternalProfile 成员对外信息
type ExternalProfile struct {
	ExternalCorpName string `json:"external_corp_name,omitempty"` // 企业对外简称
	ExternalAttr     []Attr `json:"external_attr,omitempty"`      // 属性列表
}

// ExtAttr 自定义字段
type ExtAttr struct {
	Attrs []Attr `json:"attrs,omitempty"`
}

// CreateUserParams 创建企业微信成员参数
type CreateUserParams struct {
	UserID           string          `json:"userid"`                      // 成员userid
	Name             string          `json:"name"`                        // 成员名称
	Alias            string          `json:"alias,omitempty"`             // 成员别名
	Mobile           string          `json:"mobile,omitempty"`            // 手机号码
	Department       []int           `json:"department,omitempty"`        // 成员所属部门id列表
	Order            []int           `json:"order,omitempty"`             // 部门内的排序值
	Position         string          `json:"position,omitempty"`          // 职务信息
	Gender           string          `json:"gender,omitempty"`            // 性别
	Email            string          `json:"email,omitempty"`             // 邮箱
	TelePhone        string          `json:"telephone,omitempty"`         // 座机
	IsLeaderInDept   []int           `json:"is_leader_in_dept,omitempty"` // 在所在的部门内是否为上级
	AvatarMediaID    string          `json:"avatar_mediaid,omitempty"`    // 成员头像的mediaid
	Enable           int             `json:"enable,omitempty"`            // 启用/禁用成员，1-启用，0-禁用
	ExtAttr          ExtAttr         `json:"extattr,omitempty"`           // 自定义字段
	ToInvite         bool            `json:"to_invite,omitempty"`         // 是否邀请该成员使用企业微信
	ExternalProfile  ExternalProfile `json:"external_profile,omitempty"`  // 成员对外属性
	ExternalPosition string          `json:"external_position,omitempty"` // 对外职务
	Address          string          `json:"address,omitempty"`           // 地址
	MainDepartment   int             `json:"main_department,omitempty"`   // 主部门
}

// CreateUser 创建成员
// 文档地址：https://work.weixin.qq.com/api/doc/90000/90135/90194
func (contact *Contact) CreateUser(user CreateUserParams) (err error) {
	accessToken, err := contact.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", createUserURL, accessToken)
	resp, err := util.PostJSON(url, user)
	if err != nil {
		return err
	}

	var result util.CommonError
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("create user error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	return
}

// User 成员信息
type User struct {
	UserID           string          `json:"userid"`            // 成员userid
	Name             string          `json:"name"`              // 成员名称
	Mobile           string          `json:"mobile"`            // 手机号码
	Department       []int           `json:"department"`        // 成员所属部门id列表
	Order            []int           `json:"order"`             // 部门内的排序值
	Position         string          `json:"position"`          // 职务信息
	Gender           string          `json:"gender"`            // 性别
	Email            string          `json:"email"`             // 邮箱
	IsLeaderInDept   []int           `json:"is_leader_in_dept"` // 在所在的部门内是否为上级
	Avatar           string          `json:"avatar"`            // 成员头像
	ThumbAvatar      string          `json:"thumb_avatar"`      // 成员头像缩略图
	TelePhone        string          `json:"telephone"`         // 座机
	Alias            string          `json:"alias"`             // 成员别名
	ExtAttr          ExtAttr         `json:"extattr"`           // 自定义字段
	Status           int             `json:"status"`            // 状态, 1=已激活，2=已禁用，4=未激活，5=退出企业
	QRCode           string          `json:"qr_code"`           // 员工个人二维码
	ExternalProfile  ExternalProfile `json:"external_profile"`  // 成员对外属性
	ExternalPosition string          `json:"external_position"` // 对外职务
	Address          string          `json:"address"`           // 地址
	OpenUserID       string          `json:"open_userid"`       // 全局唯一。对于同一个服务商，不同应用获取到企业内同一个成员的open_userid是相同的，最多64个字节。仅第三方应用可获取
	MainDepartment   int             `json:"main_department"`   // 主部门
	HideMobile       int             `json:"hide_mobile"`       // 是否隐藏手机号
	EnglishName      string          `json:"english_name"`      // 英文名
}

type GetUserResp struct {
	util.CommonError
	User
}

// GetUser 读取成员信息
// 文档地址：https://work.weixin.qq.com/api/doc/90000/90135/90196
func (contact *Contact) GetUser(userID string) (ret User, err error) {
	accessToken, err := contact.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s&userid=%s", createUserURL, accessToken, userID)
	resp, err := util.HTTPGet(url)
	if err != nil {
		return
	}

	var result GetUserResp
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("get user error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	ret = result.User
	return
}

// UpdateUser 更新成员信息
// 文档地址：https://work.weixin.qq.com/api/doc/90000/90135/90197
func (contact *Contact) UpdateUser(user CreateUserParams) (err error) {
	accessToken, err := contact.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", updateUserURL, accessToken)
	resp, err := util.PostJSON(url, user)
	if err != nil {
		return err
	}

	var result util.CommonError
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("update user error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	return
}

// DeleteUser 删除成员
// 文档地址：https://work.weixin.qq.com/api/doc/90000/90135/90198
func (contact *Contact) DeleteUser(userID string) (err error) {
	accessToken, err := contact.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s&userid=%s", deleteUserURL, accessToken, userID)
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
		err = fmt.Errorf("get user error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	return
}

// BatchDeleteUser 批量删除成员
// 文档地址：https://work.weixin.qq.com/api/doc/90000/90135/90199
func (contact *Contact) BatchDeleteUser(userIDs []string) (err error) {
	accessToken, err := contact.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", batchDeleteUserURL, accessToken)
	body := map[string]interface{}{
		"useridlist": userIDs,
	}
	resp, err := util.PostJSON(url, body)
	if err != nil {
		return err
	}

	var result util.CommonError
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("batch delete user error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	return
}

// SimpleUser 简易用户信息
type SimpleUser struct {
	UserID     string `json:"userid"`      // 成员userid
	Name       string `json:"name"`        // 成员名称
	Department []int  `json:"department"`  // 成员所属部门id列表
	OpenUserID string `json:"open_userid"` // 全局唯一。对于同一个服务商，不同应用获取到企业内同一个成员的open_userid是相同的，最多64个字节。仅第三方应用可获取
}

type GetDeptSimpleUsersResp struct {
	util.CommonError
	UserList []SimpleUser `json:"userlist"`
}

// GetDeptSimpleUsers 获取部门成员
// 文档地址：https://work.weixin.qq.com/api/doc/90000/90135/90200
func (contact *Contact) GetDeptSimpleUsers(deptID int, fetchChild int) (ret []SimpleUser, err error) {
	accessToken, err := contact.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s&department_id=%d&fetch_child=%d", getDeptSimpleUsersURL, accessToken, deptID, fetchChild)
	resp, err := util.HTTPGet(url)
	if err != nil {
		return
	}

	var result GetDeptSimpleUsersResp
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("get dept simple users error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	ret = result.UserList

	return
}

type GetDeptUsersResp struct {
	util.CommonError
	UserList []User `json:"userlist"`
}

// GetDeptUsers 获取部门成员详情
// 文档地址：https://work.weixin.qq.com/api/doc/90000/90135/90201
func (contact *Contact) GetDeptUsers(deptID int, fetchChild int) (ret []User, err error) {
	accessToken, err := contact.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s&department_id=%d&fetch_child=%d", getDeptUsersURL, accessToken, deptID, fetchChild)
	resp, err := util.HTTPGet(url)
	if err != nil {
		return
	}

	var result GetDeptUsersResp
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("get dept users error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	ret = result.UserList

	return
}

// ConvertToOpenIDResp userid转openid返回
type ConvertToOpenIDResp struct {
	util.CommonError
	OpenID string `json:"openid"`
}

// ConvertToOpenID userid转openid
// 文档地址：https://work.weixin.qq.com/api/doc/90000/90135/90202
func (contact *Contact) ConvertToOpenID(userID string) (openID string, err error) {
	accessToken, err := contact.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", convertToOpenIDURL, accessToken)
	body := map[string]string{
		"userid": userID,
	}
	resp, err := util.PostJSON(url, body)
	if err != nil {
		return
	}

	var result ConvertToOpenIDResp
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("convert to openid error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	openID = result.OpenID

	return
}

// ConvertToUserIDResp openid转userid返回
type ConvertToUserIDResp struct {
	util.CommonError
	UserID string `json:"userid"`
}

// ConvertToOpenID userid转openid
// 文档地址：https://work.weixin.qq.com/api/doc/90000/90135/90202
func (contact *Contact) ConvertToUserID(openID string) (userID string, err error) {
	accessToken, err := contact.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", convertToUserIDURL, accessToken)
	body := map[string]string{
		"openid": openID,
	}
	resp, err := util.PostJSON(url, body)
	if err != nil {
		return
	}

	var result ConvertToUserIDResp
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("convert to userid error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	userID = result.UserID

	return
}

// SecondAuth 二次验证
// 文档地址：https://work.weixin.qq.com/api/doc/90000/90135/90203
func (contact *Contact) SecondAuth(userID string) (err error) {
	accessToken, err := contact.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s&userid=%s", secondAuthURL, accessToken, userID)
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
		err = fmt.Errorf("second auth error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	return
}

type InviteParams struct {
	User  []string `json:"user"`
	Party []int    `json:"party"`
	Tag   []int    `json:"tag"`
}

type InviteResp struct {
	util.CommonError
	InvalidUser  []string `json:"invaliduser"`
	InvalidParty []int    `json:"invalidparty"`
	InvalidTag   []int    `json:"invalidtag"`
}

// Invite 邀请成员
// 文档地址：https://work.weixin.qq.com/api/doc/90000/90135/90975
func (contact *Contact) Invite(params InviteParams) (result InviteResp, err error) {
	accessToken, err := contact.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", batchInviteURL, accessToken)
	resp, err := util.PostJSON(url, params)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("batch invite error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	return
}

// GetJoinQRCodeResp 获取加入企业二维码返回
type GetJoinQRCodeResp struct {
	util.CommonError
	JoinQRCode string `json:"join_qrcode"`
}

// GetJoinQRCode 获取加入企业二维码
// 文档地址：https://work.weixin.qq.com/api/doc/90000/90135/91714
func (contact *Contact) GetJoinQRCode(sizeType int) (ret string, err error) {
	accessToken, err := contact.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s&size_type=%d", getJoinQRCodeURL, accessToken, sizeType)
	resp, err := util.HTTPGet(url)
	if err != nil {
		return
	}

	var result GetJoinQRCodeResp
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("get join qrcode error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	ret = result.JoinQRCode

	return
}

// GetActiveStatResp 获取企业活跃成员数返回
type GetActiveStatResp struct {
	util.CommonError
	ActiveCnt int `json:"active_cnt"`
}

// GetActiveStat 获取企业活跃成员数
// 文档地址：https://work.weixin.qq.com/api/doc/90000/90135/92714
func (contact *Contact) GetActiveStat(date string) (activeCnt int, err error) {
	accessToken, err := contact.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", getActiveStatURL, accessToken)
	body := map[string]string{
		"date": date,
	}
	resp, err := util.PostJSON(url, body)
	if err != nil {
		return
	}

	var result GetActiveStatResp
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("get active stat error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	activeCnt = result.ActiveCnt

	return
}
