package merchantDTS

import (
	merchantModel "github.com/og/thinking-in-go/ddd/merchant/repo/model"
	tj "github.com/typejson/go"
)
type ReqCreateMerchant struct {
	Name string `json:"name"`
}
func (v ReqCreateMerchant) TJ(r *tj.Rule) {
	r.String(v.Name, tj.StringSpec{
		Name:              "商家名",
	})
}
type ReplyCreateMerchant struct {
	MerchantID merchantModel.IDMerchant `json:"merchantID"`
}