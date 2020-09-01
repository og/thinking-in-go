package goodsRepo

import (
	goodsModel "github.com/og/thinking-in-go/ddd/goods/model"
	repoModel "github.com/og/thinking-in-go/ddd/shop/model"
	"strings"
)


type CreateGoodsData struct {
	ShopID repoModel.IDShop
	Title string
	Price float64
	Banner []string
	DetailPhoto []string
	Sale bool
}
func (Repo) CreateGoods(data CreateGoodsData) (goods goodsModel.Goods) {
	goods = goodsModel.Goods{
		ShopID: data.ShopID,
		Title: data.Title,
		Price: data.Price,
		Banner: data.Banner,
		DetailPhoto: data.DetailPhoto,
		Sale: data.Sale,
	}
	goods.BeforeCreate()
	database.Goods = append(database.Goods, goods)
	return goods
}
// 命名规则{数据}By{值}In{范围}
func (Repo) GoodsByTitleInRepo(searchGoodsTitle string, shopID repoModel.IDShop) (goodsModel.Goods, bool) {
	for _, goods := range database.Goods {
		if strings.Contains(goods.Title, searchGoodsTitle) {
			return goods, true
		}
	}
	return goodsModel.Goods{}, false
}

