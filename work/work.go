package work

import (
	"github.com/silenceper/wechat/v2/credential"
	"github.com/silenceper/wechat/v2/work/auth"
	"github.com/silenceper/wechat/v2/work/basic"
	"github.com/silenceper/wechat/v2/work/callback"
	"github.com/silenceper/wechat/v2/work/config"
	"github.com/silenceper/wechat/v2/work/contact"
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

// GetBasic url 相关配置
func (w *Work) GetBasic() *basic.Basic {
	return basic.NewBasic(w.ctx)
}

// GetContact 企业成员/部门/标签相关
func (w *Work) GetContact() *contact.Contact {
	return contact.NewContact(w.ctx)
}

// GetCallback 企业微信回调相关
func (w *Work) GetCallback(token, aesKey string) *callback.Callback {
	return callback.NewCallback(w.ctx, token, aesKey)
}

// GetAuth 企业微信认证鉴权相关
func (w *Work) GetAuth() *auth.Auth {
	return auth.NewAuth(w.ctx)
}
