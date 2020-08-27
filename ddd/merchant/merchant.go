package merchantService

import (
	merchantModel "github.com/og/thinking-in-go/ddd/merchant/model"
	merchantStore "github.com/og/thinking-in-go/ddd/merchant/store"
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
func (dep Service) CreateMerchant(req ReqCreateMerchant) (reply ReplyCreateMerchant, reject error) {
	/* 格式校验 */{
		reject = dep.checkReq.Check(req) ; if reject != nil { return }
	}
	/* 合法性校验 */{/* 暂无 */}
	/* 数据存储 */{
		merchant := dep.store.CreateMerchant(merchantStore.CreateMerchantData{
			Name: req.Name,
		})
		reply.MerchantID = merchant.ID
	}
	return
}