package merchantService_test

import (
	merchantService "github.com/og/thinking-in-go/ddd/merchant"
	merchantDTS "github.com/og/thinking-in-go/ddd/merchant/dts"
	merchantRepo "github.com/og/thinking-in-go/ddd/merchant/repo"
	"testing"
)

func TestService_CreateMerchant(t *testing.T) {
	merchantS := merchantService.NewService(merchantRepo.NewRepo())
	merchantS.CreateMerchant(merchantDTS.ReqCreateMerchant{
		Name: "可口可乐",
	})
}
