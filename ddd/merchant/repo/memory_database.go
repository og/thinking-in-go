package merchantRepo

import (
	merchantModel "github.com/og/thinking-in-go/ddd/merchant/repo/model"
)

var database = struct {
	Merchant []merchantModel.Merchant
}{}