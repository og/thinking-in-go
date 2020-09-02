package resUtil

import (
	ge "github.com/og/x/error"
)

type ResType string
func (ResType) Enum() (enum struct{
	Pass ResType
	Fail ResType
}) {
	enum.Pass = ResType("pass")
	enum.Fail = ResType("fail")
	return
}
type ResCode string
type Res struct {
	Type ResType `json:"type"`
	Data interface{} `json:"data"`
	Code ResCode `json:"code"`
	Msg string `json:"msg"`
}


func RejectFail(msg string) ge.Reject {
	return ge.Reject{
		Response: Res{
			Type: Res{}.Type.Enum().Fail,
			Msg:  msg,
		},
	}
}
func RejectFailAndRecord(msg string) ge.Reject {
	return ge.Reject{
		Response: Res{
			Type: Res{}.Type.Enum().Fail,
			Msg:  msg,
		},
		ShouldRecord: true,
	}
}
