package contact

import (
	"fmt"

	"github.com/silenceper/wechat/v2/util"
)

const (
	// getJobResult 获取异步任务结果API地址
	getJobResult = "https://qyapi.weixin.qq.com/cgi-bin/batch/getresult"
)

// GetJobResult 获取异步任务结果
// 文档地址：https://work.weixin.qq.com/api/doc/90000/90135/90983
func (contact *Contact) GetJobResult(jobID string) (resp []byte, err error) {
	accessToken, err := contact.GetAccessToken()
	if err != nil {
		return
	}

	url := fmt.Sprintf("%s?access_token=%s&jobid=%s", getJobResult, accessToken, jobID)
	resp, err = util.HTTPGet(url)
	if err != nil {
		return
	}

	return
}
