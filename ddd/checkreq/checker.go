package checkreq

import (
	"errors"
	tj "github.com/typejson/go"
)

var checker = tj.NewCN()
type CheckReq struct {}
func (CheckReq) Check(data tj.Data) error {
	report := checker.Scan(data)
	if report.Fail {
		return  errors.New(report.Message)
	}
	return nil
}
