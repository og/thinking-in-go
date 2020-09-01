package goodsService

import (
	goodsDTS "github.com/og/thinking-in-go/ddd/goods/dts"
	goodsRepo "github.com/og/thinking-in-go/ddd/goods/repo"
	merchantService "github.com/og/thinking-in-go/ddd/merchant"
	merchantDTS "github.com/og/thinking-in-go/ddd/merchant/dts"
	merchantRepo "github.com/og/thinking-in-go/ddd/merchant/repo"
	shopService "github.com/og/thinking-in-go/ddd/shop"
	shopDTS "github.com/og/thinking-in-go/ddd/shop/dts"
	shopRepo "github.com/og/thinking-in-go/ddd/shop/repo"
	gtest "github.com/og/x/test"
	"testing"
)

func TestCreate(t *testing.T) {
	as := gtest.NewAS(t)
	merchantS := merchantService.NewService(merchantRepo.NewRepo())
	merchantReply, reject := merchantS.CreateMerchant(merchantDTS.ReqCreateMerchant{
		Name: "测试1",
	})
	shopS := shopService.NewService(shopRepo.Repo{})
	as.NoError(reject)
	shopReply, reject := shopS.CreateShop(shopDTS.ReqCreateShop{
		MerchantID: merchantReply.MerchantID,
		Name:       "一号店",
	})
	as.NoError(reject)
	goodsS := NewService(goodsRepo.NewRepo(), shopService.NewService(shopRepo.NewRepo()))
	goodsReply, reject := goodsS.CreateGoods(goodsDTS.ReqCreateGoods{
		ShopID:      shopReply.ShopID,
		MerchantID:  merchantReply.MerchantID,
		Title:       "标题",
		Price:       10,
		Banner:      nil,
		DetailPhoto: nil,
	})
	as.NoError(reject)
	as.Equal(len(goodsReply.GoodsID), 36)
	as.Equal(goodsReply, goodsDTS.ReplyCreateGoods{
		GoodsID: goodsReply.GoodsID,
	})
}
