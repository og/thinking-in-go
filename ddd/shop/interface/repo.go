package shopInterface

import (
	merchantModel "github.com/og/thinking-in-go/ddd/merchant/repo/model"
	shopModel "github.com/og/thinking-in-go/ddd/shop/model"
	shopRepo "github.com/og/thinking-in-go/ddd/shop/repo"
)

type Repo interface {
	CreateShop(data shopRepo.CreateShopData) shopModel.Shop
	ShopByNameInMerchant(name string, merchantID merchantModel.IDMerchant) (shopModel.Shop, bool)
	ShopByRepoID(repoID shopModel.IDShop) (shopModel.Shop, bool)
}