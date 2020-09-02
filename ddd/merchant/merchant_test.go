package merchantService_test

import (
	merchantService "github.com/og/thinking-in-go/ddd/merchant"
	merchantDTS "github.com/og/thinking-in-go/ddd/merchant/dts"
	merchantRepo "github.com/og/thinking-in-go/ddd/merchant/repo"
	gtest "github.com/og/x/test"
	"testing"
)

func TestService_CreateMerchant(t *testing.T) {
	as := gtest.NewAS(t)
	merchantS := merchantService.NewService(merchantRepo.NewRepo())
	merchantReply, reject := merchantS.CreateMerchant(merchantDTS.ReqCreateMerchant{
		Name: "可口可乐",
	})
	as.Equal(len(merchantReply.MerchantID), 36)
	as.True(reject.Fail())
}
