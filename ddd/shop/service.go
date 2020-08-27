package shopService

import (
	storeStore "github.com/og/thinking-in-go/ddd/shop/store"
)

type Service struct {
	store storeStore.Store
}