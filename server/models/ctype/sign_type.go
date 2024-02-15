package ctype

import "encoding/json"

type SignStatus int

const (
	SignQQ    SignStatus = 1
	SignGitee SignStatus = 2
	SignEmail SignStatus = 3
)

func (s SignStatus) MarshalJson() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s SignStatus) String() string {
	var str string
	switch s {
	case SignQQ:
		str = "QQ"
	case SignGitee:
		str = "Gitee"
	case SignEmail:
		str = "Email"
	default:
		str = "游客"
	}
	return str
}
