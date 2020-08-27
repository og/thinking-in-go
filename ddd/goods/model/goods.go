package goodsModel

import (
	f "github.com/og/gofree"
	storeModel "github.com/og/thinking-in-go/ddd/shop/model"
)

type IDGoods string
type Goods struct {
	ID IDGoods `db:"id"`
	ShopID storeModel.IDShop `db:"shop_id"`
	Title string `db:"title"`
	Price float64 `db:"price"`
	Banner []string `db:"banner"`
	DetailPhoto []string `db:"detail_photo"`
	Sale bool `db:"sale"`

	// CreatedAt time.Time `db:"created_at"`
	// UpdatedAt time.Time `db:"updated_at"`
	// DeletedAt sql.NullTime `db:"deleted_at"`
}

func (goods *Goods) BeforeCreate() {
	goods.ID = IDGoods(f.UUID())
}