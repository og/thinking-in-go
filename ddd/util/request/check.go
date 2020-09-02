package requestUtil

import (
	resUtil "github.com/og/thinking-in-go/ddd/util/response"
	ge "github.com/og/x/error"
	tj "github.com/typejson/go"
)

var checker = tj.NewCN()
func Check(data tj.Data) ge.Reject {
	report := checker.Scan(data)
	if report.Fail {
		return  ge.Reject{
			Response:     resUtil.Res{
				Type: resUtil.Res{}.Type.Enum().Fail,
				Msg:  report.Message,
			},
			ShouldRecord: false,
		}
	}
	return ge.NotReject()
}
