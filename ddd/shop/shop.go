package shopService

import (
	"errors"
	merchantModel "github.com/og/thinking-in-go/ddd/merchant/repo/model"
	shopDTS "github.com/og/thinking-in-go/ddd/shop/dts"
	repoModel "github.com/og/thinking-in-go/ddd/shop/model"
	shopRepo "github.com/og/thinking-in-go/ddd/shop/repo"
)



func (dep Service) CreateShop(req shopDTS.ReqCreateShop) (reply shopDTS.ReplyCreateShop, reject error) {
	/* 所属权验证 */{/*暂无*/}
	/* 重复验证 */{
		_, hasRepo := dep.repo.ShopByNameInMerchant(req.Name, req.MerchantID)
		if hasRepo {
			return reply, errors.New("店铺(" + req.Name + ")已存在")
		}
	}
	/* 数据存储 */{
		shop := dep.repo.CreateShop(shopRepo.CreateShopData{
			MerchantID: req.MerchantID,
			Name:       req.Name,
		})
		reply.ShopID = shop.ID
		return
	}
}


func (dep Service) OwnershipShop(shopID repoModel.IDShop, merchantID merchantModel.IDMerchant) (reject error) {
	shop, hasShop := dep.repo.ShopByRepoID(shopID)
	if !hasShop {
		return errors.New("商店不存在")
	}
	if shop.MerchantID != merchantID {
		return errors.New("商店不属于此商家")
	}
	return nil
}
