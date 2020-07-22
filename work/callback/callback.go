package callback

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/silenceper/wechat/v2/util"

	"github.com/silenceper/wechat/v2/work/context"
)

// Callback struct
type Callback struct {
	Token  string `json:"token"`
	AESKey string `json:"aseKey"`
	*context.Context
}

// NewCallback 实例
func NewCallback(context *context.Context, token string, aesKey string) *Callback {
	cb := new(Callback)
	cb.Context = context
	cb.Token = token
	cb.AESKey = aesKey
	return cb
}

// CheckRequest 检查请求是否合法，并解密消息体
func (cb *Callback) CheckRequest(req *http.Request) (echoData []byte, reqXMLBytes []byte, err error) {
	// 验证参数是否齐全
	msgSignature := req.URL.Query().Get("msg_signature")
	timeStamp := req.URL.Query().Get("timestamp")
	nonce := req.URL.Query().Get("nonce")
	if msgSignature == "" || timeStamp == "" || nonce == "" {
		err = errors.New("missing required parameters")
		return
	}
	echoStr := req.URL.Query().Get("echostr")
	echoStr, _ = url.QueryUnescape(echoStr)

	wxBizMsgCrypt := util.NewWXBizMsgCrypt(cb.Token, cb.AESKey, "", util.XmlType)
	if echoStr != "" {
		// 校验请求是否合法
		var wxErr *util.CryptError
		echoData, wxErr = wxBizMsgCrypt.VerifyURL(msgSignature, timeStamp, nonce, echoStr)
		if wxErr != nil {
			err = errors.New(wxErr.ErrMsg)
			return
		}
		return
	}

	// 读取请求体
	reqBody, _ := ioutil.ReadAll(req.Body)
	req.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))

	if len(reqBody) > 0 {
		// 解密body体
		var wxErr1 *util.CryptError
		reqXMLBytes, wxErr1 = wxBizMsgCrypt.DecryptMsg(msgSignature, timeStamp, nonce, reqBody)
		if wxErr1 != nil {
			err = errors.New(wxErr1.ErrMsg)
			return
		}
	}

	return
}
