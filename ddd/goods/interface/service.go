package goodsInterface

import (
	goodsDTS "github.com/og/thinking-in-go/ddd/goods/dts"
)

type Service interface {
	CreateGoods(data goodsDTS.ReqCreateGoods) (reply goodsDTS.ReplyCreateGoods, reject error)
}
