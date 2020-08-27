package merchantModel

import (
	f "github.com/og/gofree"
)

type IDMerchant string
type Merchant struct {
	ID IDMerchant `db:"id"`
	Name string `db:"name"`
}
func (merchant *Merchant) BeforeCreate() {
	merchant.ID = IDMerchant(f.UUID())
}