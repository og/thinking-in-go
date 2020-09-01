package shopInterface

import (
	merchantModel "github.com/og/thinking-in-go/ddd/merchant/model"
	shopDTS "github.com/og/thinking-in-go/ddd/shop/dts"
	repoModel "github.com/og/thinking-in-go/ddd/shop/model"
)

type Service interface {
	CreateShop(req shopDTS.ReqCreateShop) (reply shopDTS.ReplyCreateShop, reject error)
	OwnershipShop(shopID repoModel.IDShop, merchantID merchantModel.IDMerchant) (reject error)
}

