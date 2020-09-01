package goodsInterface

import (
	goodsModel "github.com/og/thinking-in-go/ddd/goods/model"
	goodsRepo "github.com/og/thinking-in-go/ddd/goods/repo"
	repoModel "github.com/og/thinking-in-go/ddd/shop/model"
)

type Repo interface {
	CreateGoods(data goodsRepo.CreateGoodsData) (goods goodsModel.Goods)
	GoodsByTitleInRepo(searchGoodsTitle string, shopID repoModel.IDShop) (goodsModel.Goods, bool)
}
