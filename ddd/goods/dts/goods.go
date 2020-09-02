package goodsDTS

import (
	goodsModel "github.com/og/thinking-in-go/ddd/goods/model"
	merchantModel "github.com/og/thinking-in-go/ddd/merchant/repo/model"
	repoModel "github.com/og/thinking-in-go/ddd/shop/model"
)

type ReqCreateGoods struct {
	ShopID      repoModel.IDShop         `json:"shopID"`
	MerchantID  merchantModel.IDMerchant `json:"merchantID"`
	Title       string                   `json:"title"`
	Price       float64                  `json:"price"`
	Banner      []string                 `json:"banner"`
	DetailPhoto []string                 `json:"detailPhoto"`
}
type ReplyCreateGoods struct {
	GoodsID goodsModel.IDGoods `json:"goodsID"`
}