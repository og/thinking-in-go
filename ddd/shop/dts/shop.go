package shopDTS

import (
	merchantModel "github.com/og/thinking-in-go/ddd/merchant/repo/model"
	repoModel "github.com/og/thinking-in-go/ddd/shop/model"
)

type ReqCreateShop struct {
	MerchantID merchantModel.IDMerchant
	Name       string
}
type ReplyCreateShop struct {
	ShopID repoModel.IDShop `json:"shopID"`
}