package goodsService

import (
	"errors"
	goodsModel "github.com/og/thinking-in-go/ddd/goods/model"
	goodsStore "github.com/og/thinking-in-go/ddd/goods/store"
	merchantModel "github.com/og/thinking-in-go/ddd/merchant/model"
	storeModel "github.com/og/thinking-in-go/ddd/shop/model"
)

type ReqCreateGoods struct {
	ShopID storeModel.IDShop `json:"shopID"`
	MerchantID merchantModel.IDMerchant `json:"merchantID"`
	Title string `json:"title"`
	Price float64 `json:"price"`
	Banner []string `json:"banner"`
	DetailPhoto []string `json:"detailPhoto"`
}
type ReplyCreateGoods struct {
	GoodsID goodsModel.IDGoods `json:"goodsID"`
}
func (dep Service) CreateGoods(data ReqCreateGoods) (reply ReplyCreateGoods, reject error) {
	/* 所属权验证 */{
		reject = dep.shop.OwnershipShop(data.ShopID, data.MerchantID) ; if reject !=nil { return }
	}
	/* 重复验证 */{
		_, hasGoods :=  dep.store.GoodsByTitleInStore(data.Title, data.ShopID)
		if hasGoods {
			return reply, errors.New("商品" + data.Title +"已存在")
		}
	}
	/* 存储 */{
		initialSale := false
		goods := dep.store.CreateGoods(goodsStore.CreateGoodsData{
			ShopID:     data.ShopID,
			Title:       data.Title,
			Price:       data.Price,
			Banner:      data.Banner,
			DetailPhoto: data.DetailPhoto,
			Sale:        initialSale,
		})
		reply.GoodsID = goods.ID
	}
	return
}
