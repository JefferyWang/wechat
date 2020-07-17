package config

import "github.com/silenceper/wechat/v2/cache"

// Config config for work wechat
type Config struct {
	CorpID     string      `json:"corp_id"`     // 企业id
	CorpSecret string      `json:"corp_secret"` // 应用的凭证密钥
	AgentID    int         `json:"agent_id"`    // 应用ID
	Cache      cache.Cache // 缓存
}
