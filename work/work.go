package work

import (
	"github.com/silenceper/wechat/v2/credential"
	"github.com/silenceper/wechat/v2/work/config"
	"github.com/silenceper/wechat/v2/work/context"
)

// Work 企业微信相关API
type Work struct {
	ctx *context.Context
}

// NewWork 实例化企业微信API
func NewWork(cfg *config.Config) *Work {
	defaultAkHandle := credential.NewDefaultWorkAccessToken(cfg.CorpID, cfg.CorpSecret, cfg.AgentID, credential.CacheKeyWorkPrefix, cfg.Cache)
	ctx := &context.Context{
		Config:            cfg,
		AccessTokenHandle: defaultAkHandle,
	}
	return &Work{ctx}
}
