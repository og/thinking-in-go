package merchantRepoMock

import (
	merchantRepo "github.com/og/thinking-in-go/ddd/merchant/repo"
	merchantModel "github.com/og/thinking-in-go/ddd/merchant/repo/model"
)


func (dep MockRepo) CreateMerchant(data merchantRepo.CreateMerchantData) (merchant merchantModel.Merchant) {
	return
}


func (dep MockRepo) MerchantByName(name string) (model merchantModel.Merchant, has bool) {
	return
}
