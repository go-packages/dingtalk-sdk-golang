package sdk

import (
	"github.com/go-packages/dingtalk-sdk-golang/json"
	"testing"
)

func TestGetAuthScopes(t *testing.T) {
	resp, err := CreateClient().GetAuthScopes()
	t.Log(json.ToJson(resp))
	t.Log(err)
}
