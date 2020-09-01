package merchantRepoMock

import (
	merchantModel "github.com/og/thinking-in-go/ddd/merchant/model"
	merchantRepo "github.com/og/thinking-in-go/ddd/merchant/repo"
)


func (dep MockRepo) CreateMerchant(data merchantRepo.CreateMerchantData) (merchant merchantModel.Merchant) {
	return
}


func (dep MockRepo) MerchantByName(name string) (model merchantModel.Merchant, has bool) {
	return
}
