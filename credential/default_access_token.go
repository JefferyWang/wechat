package credential

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/util"
)

const (
	//AccessTokenURL 获取access_token的接口
	accessTokenURL = "https://api.weixin.qq.com/cgi-bin/token"
	// workAccessTokenURL 获取企业微信access_token的接口
	workAccessTokenURL = "https://qyapi.weixin.qq.com/cgi-bin/gettoken"
	//CacheKeyOfficialAccountPrefix 微信公众号cache key前缀
	CacheKeyOfficialAccountPrefix = "gowechat_officialaccount_"
	//CacheKeyMiniProgramPrefix 小程序cache key前缀
	CacheKeyMiniProgramPrefix = "gowechat_miniprogram_"
	// CacheKeyWorkPrefix 企业微信cache key前缀
	CacheKeyWorkPrefix = "gowechat_work_"
)

//DefaultAccessToken 默认AccessToken 获取
type DefaultAccessToken struct {
	appID           string
	appSecret       string
	cacheKeyPrefix  string
	cache           cache.Cache
	accessTokenLock *sync.Mutex
}

//NewDefaultAccessToken new DefaultAccessToken
func NewDefaultAccessToken(appID, appSecret, cacheKeyPrefix string, cache cache.Cache) AccessTokenHandle {
	if cache == nil {
		panic("cache is ineed")
	}
	return &DefaultAccessToken{
		appID:           appID,
		appSecret:       appSecret,
		cache:           cache,
		cacheKeyPrefix:  cacheKeyPrefix,
		accessTokenLock: new(sync.Mutex),
	}
}

// DefaultWorkAccessToken 默认企业微信AccessToken 获取
type DefaultWorkAccessToken struct {
	corpID          string
	corpSecret      string
	agentID         int
	cacheKeyPrefix  string
	cache           cache.Cache
	accessTokenLock *sync.Mutex
}

// NewDefaultWorkAccessToken new DefaultWorkAccessToken
func NewDefaultWorkAccessToken(corpID, corpSecret string, agentID int, cacheKeyPrefix string, cache cache.Cache) AccessTokenHandle {
	if cache == nil {
		panic("cache is ineed")
	}
	return &DefaultWorkAccessToken{
		corpID:          corpID,
		corpSecret:      corpSecret,
		agentID:         agentID,
		cache:           cache,
		cacheKeyPrefix:  cacheKeyPrefix,
		accessTokenLock: new(sync.Mutex),
	}
}

//ResAccessToken struct
type ResAccessToken struct {
	util.CommonError

	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

//GetAccessToken 获取access_token,先从cache中获取，没有则从服务端获取
func (ak *DefaultAccessToken) GetAccessToken() (accessToken string, err error) {
	//加上lock，是为了防止在并发获取token时，cache刚好失效，导致从微信服务器上获取到不同token
	ak.accessTokenLock.Lock()
	defer ak.accessTokenLock.Unlock()

	accessTokenCacheKey := fmt.Sprintf("%s_access_token_%s", ak.cacheKeyPrefix, ak.appID)
	val := ak.cache.Get(accessTokenCacheKey)
	if val != nil {
		accessToken = val.(string)
		return
	}

	//cache失效，从微信服务器获取
	var resAccessToken ResAccessToken
	resAccessToken, err = GetTokenFromServer(ak.appID, ak.appSecret)
	if err != nil {
		return
	}

	expires := resAccessToken.ExpiresIn - 1500
	err = ak.cache.Set(accessTokenCacheKey, resAccessToken.AccessToken, time.Duration(expires)*time.Second)
	if err != nil {
		return
	}
	accessToken = resAccessToken.AccessToken
	return
}

func (ak *DefaultWorkAccessToken) GetAccessToken() (accessToken string, err error) {
	// 加上lock，是为了防止在并发获取token时，cache刚好失效，导致从微信服务器上获取到不同token
	ak.accessTokenLock.Lock()
	defer ak.accessTokenLock.Unlock()

	accessTokenCacheKey := fmt.Sprintf("%s_access_token_%s_%d", ak.cacheKeyPrefix, ak.corpID, ak.agentID)
	val := ak.cache.Get(accessTokenCacheKey)
	if val != nil {
		accessToken = val.(string)
		return
	}

	// cache失效，从企业微信服务器获取
	var resAccessToken ResAccessToken
	resAccessToken, err = GetWorkTokenFromServer(ak.corpID, ak.corpSecret)
	if err != nil {
		return
	}

	expires := resAccessToken.ExpiresIn - 1500
	err = ak.cache.Set(accessTokenCacheKey, resAccessToken.AccessToken, time.Duration(expires)*time.Second)
	if err != nil {
		return
	}
	accessToken = resAccessToken.AccessToken
	return
}

//GetTokenFromServer 强制从微信服务器获取token
func GetTokenFromServer(appID, appSecret string) (resAccessToken ResAccessToken, err error) {
	url := fmt.Sprintf("%s?grant_type=client_credential&appid=%s&secret=%s", accessTokenURL, appID, appSecret)
	var body []byte
	body, err = util.HTTPGet(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &resAccessToken)
	if err != nil {
		return
	}
	if resAccessToken.ErrMsg != "" {
		err = fmt.Errorf("get access_token error : errcode=%v , errormsg=%v", resAccessToken.ErrCode, resAccessToken.ErrMsg)
		return
	}
	return
}

func GetWorkTokenFromServer(corpID, corpSecret string) (resAccessToken ResAccessToken, err error) {
	url := fmt.Sprintf("%s?corpid=%s&corpsecret=%s", workAccessTokenURL, corpID, corpSecret)
	var body []byte
	body, err = util.HTTPGet(url)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &resAccessToken)
	if err != nil {
		return
	}
	if resAccessToken.ErrMsg != "" {
		err = fmt.Errorf("get work wechat access_token error : errcode=%v , errormsg=%v", resAccessToken.ErrCode, resAccessToken.ErrMsg)
		return
	}
	return
}
