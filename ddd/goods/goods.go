package goodsService

import (
	"errors"
	goodsDTS "github.com/og/thinking-in-go/ddd/goods/dts"
	goodsRepo "github.com/og/thinking-in-go/ddd/goods/repo"
)


func (dep Service) CreateGoods(data goodsDTS.ReqCreateGoods) (reply goodsDTS.ReplyCreateGoods, reject error) {
	/* 所属权验证 */{
		reject = dep.shop.OwnershipShop(data.ShopID, data.MerchantID) ; if reject !=nil { return }
	}
	/* 重复验证 */{
		_, hasGoods :=  dep.repo.GoodsByTitleInRepo(data.Title, data.ShopID)
		if hasGoods {
			return reply, errors.New("商品" + data.Title +"已存在")
		}
	}
	/* 存储 */{
		initialSale := false
		goods := dep.repo.CreateGoods(goodsRepo.CreateGoodsData{
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
