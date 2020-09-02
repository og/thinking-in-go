package merchantRepoInterface

import (
	merchantRepoDTS "github.com/og/thinking-in-go/ddd/merchant/repo/dts"
	merchantModel "github.com/og/thinking-in-go/ddd/merchant/repo/model"
)

type Repo interface {
	CreateMerchant(data merchantRepoDTS.CreateMerchantData) (merchant merchantModel.Merchant)
	MerchantByName(name string) (model merchantModel.Merchant, has bool)
}
