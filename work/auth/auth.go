package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/silenceper/wechat/v2/util"

	"github.com/silenceper/wechat/v2/work/context"
)

const (
	// webAppAuthRedirectURL 网页授权登录地址
	webAppAuthRedirectURL = "https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&scope=snsapi_base&state=%s#wechat_redirect"
	// qrAppAuthRedirectURL 扫码授权登录地址
	qrAppAuthRedirectURL = "https://open.work.weixin.qq.com/wwopen/sso/qrConnect?appid=%s&agentid=%d&redirect_uri=%s&state=%s"
	// getUserInfoURL 获取访问用户身份地址
	getUserInfoURL = "https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo?access_token=%s&code=%s"
)

// Auth struct
type Auth struct {
	*context.Context
}

// NewAuth 实例
func NewAuth(context *context.Context) *Auth {
	basic := new(Auth)
	basic.Context = context
	return basic
}

// GetWebAppRedirectURL 获取网页应用跳转的url地址
func (oauth *Auth) GetWebAppRedirectURL(redirectURI, state string) (string, error) {
	// url encode
	urlStr := url.QueryEscape(redirectURI)
	return fmt.Sprintf(webAppAuthRedirectURL, oauth.CorpID, urlStr, state), nil
}

// WebAppRedirect 跳转到网页授权
func (oauth *Auth) WebAppRedirect(writer http.ResponseWriter, req *http.Request, redirectURI, state string) error {
	location, err := oauth.GetWebAppRedirectURL(redirectURI, state)
	if err != nil {
		return err
	}
	http.Redirect(writer, req, location, http.StatusFound)
	return nil
}

// GetQRAppRedirectURL 获取扫码授权登录的url地址
func (oauth *Auth) GetQRAppRedirectURL(redirectURI, state string) (string, error) {
	// url encode
	urlStr := url.QueryEscape(redirectURI)
	return fmt.Sprintf(qrAppAuthRedirectURL, oauth.CorpID, oauth.AgentID, urlStr, state), nil
}

// QRAppRedirect 跳转到扫码授权登录
func (oauth *Auth) QRAppRedirect(writer http.ResponseWriter, req *http.Request, redirectURI, state string) error {
	location, err := oauth.GetQRAppRedirectURL(redirectURI, state)
	if err != nil {
		return err
	}
	http.Redirect(writer, req, location, http.StatusFound)
	return nil
}

type GetUserInfoResp struct {
	util.CommonError
	UserID   string `json:"UserId"`   // 成员UserID
	OpenID   string `json:"OpenId"`   // 非企业成员的标识，对当前企业唯一。不超过64字节
	DeviceID string `json:"DeviceId"` // 网页应用会返回，扫码授权登录不会返回。手机设备号(由企业微信在安装时随机生成，删除重装会改变，升级不受影响)
}

// GetUserInfo 获取访问用户身份
// 文档地址：https://work.weixin.qq.com/api/doc/90000/90135/91023
func (oauth *Auth) GetUserInfo(code string) (result GetUserInfoResp, err error) {
	ak, err := oauth.GetAccessToken()
	if err != nil {
		return
	}
	urlStr := fmt.Sprintf(getUserInfoURL, ak, code)
	var response []byte
	response, err = util.HTTPGet(urlStr)
	if err != nil {
		return
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.ErrCode != 0 {
		err = fmt.Errorf("GetUserInfo error : errcode=%v , errmsg=%v", result.ErrCode, result.ErrMsg)
		return
	}
	return
}
