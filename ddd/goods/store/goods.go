package goodsStore

import (
	goodsModel "github.com/og/thinking-in-go/ddd/goods/model"
	storeModel "github.com/og/thinking-in-go/ddd/shop/model"
	"strings"
)

type Store struct {}

type CreateGoodsData struct {
	ShopID storeModel.IDShop
	Title string
	Price float64
	Banner []string
	DetailPhoto []string
	Sale bool
}
func (Store) CreateGoods(data CreateGoodsData) (goods goodsModel.Goods) {
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
func (Store) GoodsByTitleInStore(searchGoodsTitle string, shopID storeModel.IDShop) (goodsModel.Goods, bool) {
	for _, goods := range database.Goods {
		if strings.Contains(goods.Title, searchGoodsTitle) {
			return goods, true
		}
	}
	return goodsModel.Goods{}, false
}

