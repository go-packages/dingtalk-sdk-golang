package dingding_sdk_golang

type BaseResp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type GetDepMemberIdsResp struct {
	BaseResp

	UserIds []string `json:"userIds"`
}

type GetUserDetailResp struct {
	BaseResp

	UserList
}

type GetDepMemberListResp struct {
	BaseResp

	HasMore  bool       `json:"hasMore"`
	UserList []UserList `json:"userlist"`
}

type UserList struct {
	Name            string           `json:"name"`
	UnionId         string           `json:"unionid"`
	UserId          string           `json:"userid"`
	IsLeaderInDepts string           `json:"isLeaderInDepts"`
	IsBoos          bool             `json:"isBoss"`
	HiredDate       int64            `json:"hiredDate"`
	IsSenior        bool             `json:"isSenior"`
	Department      []int            `json:"department"`
	OrderInDepts    string           `json:"orderInDepts"`
	Active          bool             `json:"active"`
	Avatar          string           `json:"avatar"`
	IsAdmin         bool             `json:"isAdmin"`
	IsHide          bool             `json:"isHide"`
	JobNumber       string           `json:"jobnumber"`
	Position        string           `json:"position"`
	DingId          string           `json:"dingId"`
	Roles           []UserDetailRole `json:"roles"`
}

type UserDetailRole struct {
	Id        int64  `json:"id"`
	Name      string `json:"name"`
	GroupName string `json:"groupName"`
}
