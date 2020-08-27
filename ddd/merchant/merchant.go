package merchantService

import (
	merchantModel "github.com/og/thinking-in-go/ddd/merchant/model"
	merchantStore "github.com/og/thinking-in-go/ddd/merchant/store"
)

type ReqCreateMerchant struct {
	Name string `json:"name"`
}
type ReplyCreateMerchant struct {
	MerchantID merchantModel.IDMerchant `json:"merchantID"`
}
func (dep Service) CreateMerchant(req ReqCreateMerchant) (reply ReplyCreateMerchant) {
	merchant := dep.store.CreateMerchant(merchantStore.CreateMerchantData{
		Name: req.Name,
	})
	reply.MerchantID = merchant.ID
	return
}