package goodsRepo

import (
	goodsModel "github.com/og/thinking-in-go/ddd/goods/model"
)
// 为展示数据分层，将数据先保存在内存中。
var database = struct {
	Goods []goodsModel.Goods
}{}
