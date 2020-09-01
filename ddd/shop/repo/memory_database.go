package shopRepo

import (
	 "github.com/og/thinking-in-go/ddd/shop/model"
)

var database = struct {
	Shop []shopModel.Shop
}{}