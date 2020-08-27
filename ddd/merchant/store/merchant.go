package merchantStore

import (
	merchantModel "github.com/og/thinking-in-go/ddd/merchant/model"
)

type CreateMerchantData struct {
	Name string
}
func (Store) CreateMerchant(data CreateMerchantData) (merchant merchantModel.Merchant) {
	merchant = merchantModel.Merchant{
		Name: data.Name,
	}
	merchant.BeforeCreate()
	database.Merchant = append(database.Merchant, merchant)
	return
}
