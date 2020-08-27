package merchantStore

import (
	merchantModel "github.com/og/thinking-in-go/ddd/merchant/model"
)

var database = struct {
	Merchant []merchantModel.Merchant
}{}