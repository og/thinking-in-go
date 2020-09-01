package merchantRepoInterface

import (
	merchantModel "github.com/og/thinking-in-go/ddd/merchant/model"
	merchantRepo "github.com/og/thinking-in-go/ddd/merchant/repo"
)

type Repo interface {
	CreateMerchant(data merchantRepo.CreateMerchantData) (merchant merchantModel.Merchant)
	MerchantByName(name string) (model merchantModel.Merchant, has bool)
}
