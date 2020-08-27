package shopService

import (
	"errors"
	merchantModel "github.com/og/thinking-in-go/ddd/merchant/model"
	storeModel "github.com/og/thinking-in-go/ddd/shop/model"
	shopStore "github.com/og/thinking-in-go/ddd/shop/store"
)


type ReqCreateShop struct {
	MerchantID merchantModel.IDMerchant
	Name string
}
type ReplyCreateShop struct {
	ShopID storeModel.IDShop `json:"shopID"`
}
func (dep Service) CreateShop(req ReqCreateShop) (reply ReplyCreateShop, reject error) {
	/* 所属权验证 */{/*暂无*/}
	/* 重复验证 */{
		_, hasStore := dep.store.ShopByNameInMerchant(req.Name, req.MerchantID)
		if hasStore {
			return reply, errors.New("店铺(" + req.Name + ")已存在")
		}
	}
	/* 数据存储 */{
		shop := dep.store.CreateShop(shopStore.CreateShopData{
			MerchantID: req.MerchantID,
			Name:       req.Name,
		})
		reply.ShopID = shop.ID
		return
	}
}


func (dep Service) OwnershipShop(shopID storeModel.IDShop, merchantID merchantModel.IDMerchant) (reject error) {
	shop, hasShop := dep.store.ShopByStoreID(shopID)
	if !hasShop {
		return errors.New("商店不存在")
	}
	if shop.MerchantID != merchantID {
		return errors.New("商店不属于此商家")
	}
	return nil
}
