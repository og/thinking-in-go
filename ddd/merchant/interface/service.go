package merchantInterface

import (
	merchantDTS "github.com/og/thinking-in-go/ddd/merchant/dts"
)

type Service interface {
	CreateMerchant(req merchantDTS.ReqCreateMerchant) (reply merchantDTS.ReplyCreateMerchant, reject error)
}
