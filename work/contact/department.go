package contact

import (
	"encoding/json"
	"fmt"

	"github.com/silenceper/wechat/v2/work/job"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// createDepartmentURL 创建部门接口地址
	createDepartmentURL = "https://qyapi.weixin.qq.com/cgi-bin/department/create"
	// updateDepartmentURL 更新部门接口地址
	updateDepartmentURL = "https://qyapi.weixin.qq.com/cgi-bin/department/update"
	// deleteDepartmentURL 删除部门接口地址
	deleteDepartmentURL = "https://qyapi.weixin.qq.com/cgi-bin/department/delete"
	// getDepartmentListURL 获取部门列表接口地址
	getDepartmentListURL = "https://qyapi.weixin.qq.com/cgi-bin/department/list"
	// batchReplacePartyURL 全量覆盖部门接口地址
	batchReplacePartyURL = "https://qyapi.weixin.qq.com/cgi-bin/batch/replaceparty"
)

// Department 部门
type Department struct {
	ID       int    `json:"id"`                // 部门id
	Name     string `json:"name"`              // 部门名称
	NameEn   string `json:"name_en,omitempty"` // 英文名称
	ParentID int    `json:"parentid"`          // 父部门id
	Order    int    `json:"order,omitempty"`   // 在父部门中的次序值
}

// CreateDepartmentResp 创建部门返回结果
type CreateDepartmentResp struct {
	util.CommonError
	ID int `json:"id"` // 创建的部门id
}

// CreateDepartment 创建部门
// 文档地址：https://work.weixin.qq.com/api/doc/90000/90135/90205
func (contact *Contact) CreateDepartment(dept *Department) (ID int, err error) {
	accessToken, err := contact.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", createDepartmentURL, accessToken)
	resp, err := util.PostJSON(url, dept)
	if err != nil {
		return
	}

	var result CreateDepartmentResp
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("create department error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	ID = result.ID

	return
}

// UpdateDepartment 更新部门
// 文档地址：https://work.weixin.qq.com/api/doc/90000/90135/90206
func (contact *Contact) UpdateDepartment(dept *Department) (err error) {
	accessToken, err := contact.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", updateDepartmentURL, accessToken)
	resp, err := util.PostJSON(url, dept)
	if err != nil {
		return
	}

	var result util.CommonError
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("update department error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	return
}

// DeleteDepartment 删除部门
// 文档地址：https://work.weixin.qq.com/api/doc/90000/90135/90207
func (contact *Contact) DeleteDepartment(deptID int) (err error) {
	accessToken, err := contact.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s&id=%d", deleteDepartmentURL, accessToken, deptID)
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
		err = fmt.Errorf("delete department error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	return
}

// GetDepartmentListResp 获取部门列表返回
type GetDepartmentListResp struct {
	util.CommonError
	Department []Department `json:"department"`
}

// GetDepartmentList 获取部门列表
// 文档地址：https://work.weixin.qq.com/api/doc/90000/90135/90208
func (contact *Contact) GetDepartmentList(deptID ...int) (result GetDepartmentListResp, err error) {
	accessToken, err := contact.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", getDepartmentListURL, accessToken)
	if len(deptID) > 0 {
		url = fmt.Sprintf("%s?access_token=%s&id=%d", getDepartmentListURL, accessToken, deptID[0])
	}
	resp, err := util.HTTPGet(url)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("get department list error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	return
}

// BatchReplaceParty 全量覆盖部门
// 文档地址：https://work.weixin.qq.com/api/doc/90000/90135/90982
func (contact *Contact) BatchReplaceParty(params *job.Params) (jobID string, err error) {
	accessToken, err := contact.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s", batchReplacePartyURL, accessToken)
	resp, err := util.PostJSON(url, params)
	if err != nil {
		return
	}

	var result job.Resp
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("batch replace party error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	jobID = result.JobID

	return
}

// PartyJobResultItem 部门任务结果信息
type PartyJobResultItem struct {
	Action  int    `json:"action"`
	PartyID int    `json:"partyid"`
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// GetPartyJobResultResp 企业部门异步任务结果返回
type GetPartyJobResultResp struct {
	util.CommonError
	Status     int                  `json:"status"`
	Type       string               `json:"type"`
	Total      int                  `json:"total"`
	Percentage int                  `json:"percentage"`
	Result     []PartyJobResultItem `json:"result"`
}

// GetPartyJobResult 获取企业部门异步任务结果
// 文档地址：https://work.weixin.qq.com/api/doc/90000/90135/90983
func (contact *Contact) GetPartyJobResult(jobID string) (result GetPartyJobResultResp, err error) {
	resp, err := contact.GetJobResult(jobID)
	if err != nil {
		return
	}

	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("get party job result error, errcode=%d,errmsg=%s", result.ErrCode, result.ErrMsg)
	}

	return
}
