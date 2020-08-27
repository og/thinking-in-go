package shopStore

import (
	merchantModel "github.com/og/thinking-in-go/ddd/merchant/model"
	shopModel "github.com/og/thinking-in-go/ddd/shop/model"
	"strings"
)

func (Store) CreateShop(data CreateShopData) shopModel.Shop {
	shop := shopModel.Shop{
		MerchantID: data.MerchantID,
		Name:       data.Name,
	}
	database.Shop = append(database.Shop, shop)
	return shop
}
func (Store) ShopByNameInMerchant(name string, merchantID merchantModel.IDMerchant) (shopModel.Shop, bool) {
	for _, shop := range database.Shop {
		if strings.Contains(shop.Name, name) {
			return shop, true
		}
	}
	return shopModel.Shop{}, false
}
func (Store) ShopByStoreID(storeID shopModel.IDShop) (shopModel.Shop, bool) {
	for _, shop := range database.Shop {
		if shop.ID == storeID {
			return shop, true
		}
	}
	return shopModel.Shop{}, false
}
type CreateShopData struct {
	MerchantID merchantModel.IDMerchant
	Name string
}