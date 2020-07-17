package basic

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
	"github.com/silenceper/wechat/v2/work/context"
)

const (
	// getAPIDomainIPURL 获取企业微信API域名IP段地址
	getAPIDomainIPURL = "https://qyapi.weixin.qq.com/cgi-bin/get_api_domain_ip"
	// getCallbackIPURL 获取企业微信回调服务器的ip段地址
	getCallbackIPURL = "https://qyapi.weixin.qq.com/cgi-bin/getcallbackip"
)

// Basic struct
type Basic struct {
	*context.Context
}

// NewBasic 实例
func NewBasic(context *context.Context) *Basic {
	basic := new(Basic)
	basic.Context = context
	return basic
}

// IPListRes 获取企业微信服务器IP地址 返回结果
type IPListRes struct {
	util.CommonError
	IPList []string `json:"ip_list"`
}

// GetAPIDomainIP 获取企业微信API域名IP段
// 文档地址：https://work.weixin.qq.com/api/doc/90000/90135/92520
func (basic *Basic) GetAPIDomainIP() ([]string, error) {
	ak, err := basic.GetAccessToken()
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s?access_token=%s", getAPIDomainIPURL, ak)
	data, err := util.HTTPGet(url)
	if err != nil {
		return nil, err
	}
	ipListRes := &IPListRes{}
	err = util.DecodeWithError(data, ipListRes, "GetWorkAPIDomainIP")
	return ipListRes.IPList, err
}

// GetCallbackIP 获取企业微信回调服务器的ip段地址
// 文档地址： https://work.weixin.qq.com/api/doc/90000/90135/90930
func (basic *Basic) GetCallbackIP() ([]string, error) {
	ak, err := basic.GetAccessToken()
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s?access_token=%s", getCallbackIPURL, ak)
	data, err := util.HTTPGet(url)
	if err != nil {
		return nil, err
	}
	ipListRes := &IPListRes{}
	err = util.DecodeWithError(data, ipListRes, "GetWorkCallbackIP")
	return ipListRes.IPList, err
}
