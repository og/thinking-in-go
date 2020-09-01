package shopRepo

import (
	merchantModel "github.com/og/thinking-in-go/ddd/merchant/model"
	shopModel "github.com/og/thinking-in-go/ddd/shop/model"
	"strings"
)
type CreateShopData struct {
	MerchantID merchantModel.IDMerchant
	Name string
}
func (Repo) CreateShop(data CreateShopData) shopModel.Shop {
	shop := shopModel.Shop{
		MerchantID: data.MerchantID,
		Name:       data.Name,
	}
	database.Shop = append(database.Shop, shop)
	return shop
}
func (Repo) ShopByNameInMerchant(name string, merchantID merchantModel.IDMerchant) (shopModel.Shop, bool) {
	for _, shop := range database.Shop {
		if strings.Contains(shop.Name, name) {
			return shop, true
		}
	}
	return shopModel.Shop{}, false
}
func (Repo) ShopByRepoID(repoID shopModel.IDShop) (shopModel.Shop, bool) {
	for _, shop := range database.Shop {
		if shop.ID == repoID {
			return shop, true
		}
	}
	return shopModel.Shop{}, false
}
