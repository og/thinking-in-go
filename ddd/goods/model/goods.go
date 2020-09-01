package goodsModel

import (
	f "github.com/og/gofree"
	repoModel "github.com/og/thinking-in-go/ddd/shop/model"
)

type IDGoods string
type Goods struct {
	ID IDGoods `db:"id"`
	ShopID repoModel.IDShop `db:"shop_id"`
	Title string `db:"title"`
	Price float64 `db:"price"`
	Banner []string `db:"banner"`
	DetailPhoto []string `db:"detail_photo"`
	Sale bool `db:"sale"`
}

func (goods *Goods) BeforeCreate() {
	goods.ID = IDGoods(f.UUID())
}
type IDGoodsAttr string
func(id IDGoodsAttr) String() string {return string(id)}

type GoodsAttrKind string
func (attr GoodsAttrKind) String() string {return string(attr)}
func (GoodsAttrKind) Enum() (enum struct{
	Text GoodsAttrKind
	SingleSelect GoodsAttrKind
	MultipleSelect GoodsAttrKind
}) {
	enum.Text = "text"
	enum.SingleSelect = "singleSelect"
	enum.MultipleSelect = "multipleSelect"
	return
}
type GoodsAttr struct {
	ID IDGoodsAttr `db:"id"`
	Kind  GoodsAttrKind `db:"kind"`
	Name string `db:"name"`

}
type IDGoodsAttrValue string
type GoodsAttrValue struct {
	ID IDGoodsAttrValue `db:"id"`
	GoodsAttrID IDGoodsAttr `db:"goods_attr_id"`
	TextValue string `db:"text_value"`
}
type IDGoodsAttrSelect string
func (id IDGoodsAttrSelect) String() string { return string(id)}
type GoodsAttrSelect struct {
	ID IDGoodsAttrSelect `db:"id"`
	GoodsAttrID IDGoodsAttr `db:"goods_attr_id"`
	Value string `db:"value"`
}

type GoodsGather struct {
	Goods Goods
	GoodsAttr []GoodsAttr
	GoodsAttrValue []GoodsAttrValue
}