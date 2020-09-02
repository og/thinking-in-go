package shopModel

import (
	f "github.com/og/gofree"
	merchantModel "github.com/og/thinking-in-go/ddd/merchant/repo/model"
)

type IDShop string
type Shop struct {
	ID         IDShop                   `db:"id"`
	MerchantID merchantModel.IDMerchant `db:"merchant_id"`
	Name       string                   `db:"name"`
}
func (shop *Shop) BeforeCreate() {
	shop.ID = IDShop(f.UUID())
}