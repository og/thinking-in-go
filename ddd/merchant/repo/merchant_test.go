package merchantRepo_test

import (
	merchantModel "github.com/og/thinking-in-go/ddd/merchant/model"
	merchantRepo "github.com/og/thinking-in-go/ddd/merchant/repo"
	gtest "github.com/og/x/test"
	"testing"
)

func TestRepo_CreateMerchant(t *testing.T) {
	as := gtest.NewAS(t)
	repo := merchantRepo.NewRepo()
	merchantName := "可口可乐"
	repo.CreateMerchant(merchantRepo.CreateMerchantData{
		Name: merchantName,
	})
	{
		merchant, hasMerchant := repo.MerchantByName(merchantName)
		as.True(hasMerchant)
		as.Equal(len(merchant.ID), 36)
		as.Equal(merchant, merchantModel.Merchant{
			ID:   merchant.ID,
			Name: merchantName,
		})
	}
}
