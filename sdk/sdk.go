package sdk

import (
	"github.com/go-packages/dingtalk-sdk-golang/encrypt"
	"github.com/go-packages/dingtalk-sdk-golang/http"
	"os"
	"strconv"
	"time"
)

type DingTalkSDK struct {
	SuiteKey    string
	SuiteSecret string
	Token       string
	AesKey      string
	AppId       int64
}

type Corp struct {
	CorpId      string
	SuiteTicket string
	SuiteKey    string
	SuiteSecret string
}

type DingTalkClient struct {
	AccessToken string
	AgentId     int64
}

type DingTalkOauthClient struct {
	OauthAppId     string
	OauthAppSecret string
}

func NewDingTalkOauthClient() *DingTalkOauthClient {
	return &DingTalkOauthClient{
		OauthAppId:     os.Getenv("OAUTH_APP_ID"),
		OauthAppSecret: os.Getenv("OAUTH_APP_SECRET"),
	}
}

func NewSDK() *DingTalkSDK {
	appId, err := strconv.ParseInt(os.Getenv("APP_ID"), 10, 64)
	if err != nil {
		panic(err)
	}
	return &DingTalkSDK{
		SuiteKey:    os.Getenv("SUITE_KEY"),
		SuiteSecret: os.Getenv("SUITE_SECRET"),
		Token:       os.Getenv("SUITE_TOKEN"),
		AesKey:      os.Getenv("SUITE_AES_KEY"),
		AppId:       appId,
	}
}

func NewCorp(suiteTicket string, corpId string) *Corp {
	return &Corp{
		CorpId:      corpId,
		SuiteTicket: suiteTicket,
		SuiteKey:    os.Getenv("SUITE_KEY"),
		SuiteSecret: os.Getenv("SUITE_SECRET"),
	}
}

func NewDingTalkClient(accessToken string, agentId int64) *DingTalkClient {
	return &DingTalkClient{
		AccessToken: accessToken,
		AgentId:     agentId,
	}
}

func (s *DingTalkSDK) CreateCrypto() *Crypto {
	if s.SuiteKey == "" {
		panic("SUITE_KEY is not config in env!")
	}
	if s.Token == "" {
		panic("SUITE_TOKEN is not config in env!")
	}
	if s.AesKey == "" {
		panic("SUITE_AES_KEY is not config in env!")
	}
	return NewCrypto(s.Token, s.AesKey, s.SuiteKey)
}

func (s *DingTalkSDK) CreateCorp(corpId string, suiteTicket string) *Corp {
	if s.SuiteKey == "" {
		panic("SUITE_KEY is not config in env!")
	}
	if s.SuiteSecret == "" {
		panic("SUITE_SECRET is not config in env!")
	}
	return &Corp{
		CorpId:      corpId,
		SuiteTicket: suiteTicket,
		SuiteKey:    s.SuiteKey,
		SuiteSecret: s.SuiteSecret,
	}
}

func (corp *Corp) CreateDingTalkClient() (*DingTalkClient, error) {
	tokenInfo, err := corp.GetCorpToken()
	if err != nil {
		return nil, err
	}
	authInfo, err := corp.GetAuthInfo()
	if err != nil {
		return nil, err
	}
	return NewDingTalkClient(tokenInfo.AccessToken, authInfo.AuthInfo.Agent[0].AgentId), nil
}

func ExcuteOapi(url string, oauthAppId string, oauthAppSecret string, body string) (string, error) {
	timestamp := time.Now().UnixNano() / 1e6
	nativeSignature := strconv.FormatInt(timestamp, 10)

	afterHmacSHA256 := encrypt.SHA256(nativeSignature, oauthAppSecret)
	afterBase64 := encrypt.BASE64(afterHmacSHA256)
	afterUrlEncode := encrypt.URLEncode(afterBase64)

	params := map[string]string{
		"timestamp": strconv.FormatInt(timestamp, 10),
		"accessKey": oauthAppId,
		"signature": afterUrlEncode,
	}

	return http.Post(url, params, body)
}
