package sdk

import (
	"github.com/go-packages/dingtalk-sdk-golang/http"
	"github.com/go-packages/dingtalk-sdk-golang/json"
)

func (client *DingTalkClient) UploadMedia(mediaType string, path string) (UploadMediaResp, error) {

	params := map[string]string{
		"access_token": client.AccessToken,
		"type":         mediaType,
	}
	body, err := http.PostFile("https://oapi.dingtalk.com/media/upload", params, path, "media")
	resp := UploadMediaResp{}
	if err != nil {
		return resp, err
	}
	json.FromJson(body, &resp)
	return resp, err
}
