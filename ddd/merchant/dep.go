package merchantService

import (
	"github.com/og/thinking-in-go/ddd/checkreq"
	merchantStore "github.com/og/thinking-in-go/ddd/merchant/store"
)

type Service struct {
	store merchantStore.Store
	checkReq checkreq.CheckReq
}
