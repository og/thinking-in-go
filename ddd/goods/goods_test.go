package goodsService

import (
	merchantService "github.com/og/thinking-in-go/ddd/merchant"
	shopService "github.com/og/thinking-in-go/ddd/shop"
	"log"
	"testing"
)

func TestCreate(t *testing.T) {
	merchantReply := merchantService.Service{}.CreateMerchant(merchantService.ReqCreateMerchant{
		Name: "测试1",
	})
	shopReply, reject := shopService.Service{}.CreateShop(shopService.ReqCreateShop{
		MerchantID: merchantReply.MerchantID,
		Name:       "一号店",
	})
	if reject != nil {panic(reject)}
	goodsReply, reject := Service{}.CreateGoods(ReqCreateGoods{
		ShopID:      shopReply.ShopID,
		MerchantID:  merchantReply.MerchantID,
		Title:       "标题",
		Price:       10,
		Banner:      nil,
		DetailPhoto: nil,
	})
	if reject != nil {panic(reject)}
	log.Print(goodsReply)
}
