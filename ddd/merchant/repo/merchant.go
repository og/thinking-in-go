package merchantRepo

import (
	merchantModel "github.com/og/thinking-in-go/ddd/merchant/model"
	"strings"
)

type CreateMerchantData struct {
	Name string
}
func (Repo) MerchantByName(name string) (model merchantModel.Merchant, has bool) {
	for _, merchant := range database.Merchant {
		if strings.Contains(merchant.Name, name) {
			return merchant, true
		}
	}
	return merchantModel.Merchant{}, false
}
func (Repo) CreateMerchant(data CreateMerchantData) (merchant merchantModel.Merchant) {
	merchant = merchantModel.Merchant{
		Name: data.Name,
	}
	merchant.BeforeCreate()
	database.Merchant = append(database.Merchant, merchant)
	return
}
