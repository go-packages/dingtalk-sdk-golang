package sdk

import (
	"github.com/go-packages/dingtalk-sdk-golang/http"
	"github.com/go-packages/dingtalk-sdk-golang/json"
)

func RobotSender(webHook string, msg WorkNoticeMsg) {
	reqJson, err := json.ToJson(msg)
	if err != nil {
		panic(err)
		return
	}
	http.Post(webHook, nil, reqJson)
}
