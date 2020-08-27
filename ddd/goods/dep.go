package goodsService

import (
	goodsStore "github.com/og/thinking-in-go/ddd/goods/store"
	shopService "github.com/og/thinking-in-go/ddd/shop"
)

type Service struct {
	store goodsStore.Store
	shop shopService.Service
}