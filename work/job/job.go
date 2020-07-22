package job

import (
	"github.com/silenceper/wechat/v2/util"
)

// Params 异步任务参数
type Params struct {
	MediaID  string   `json:"media_id"`
	ToInvite bool     `json:"to_invite"`
	Callback Callback `json:"callback"`
}

// Resp 异步任务返回
type Resp struct {
	util.CommonError
	JobID string `json:"jobid"`
}

// Callback 回调信息
type Callback struct {
	URL            string `json:"url"`
	Token          string `json:"token"`
	EncodingAESKey string `json:"encodingaeskey"`
}
