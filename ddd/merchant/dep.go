package merchantService

import (
	merchantStore "github.com/og/thinking-in-go/ddd/merchant/store"
)

type Service struct {
	store merchantStore.Store
}
