package goodsService

import (
	goodsInterface "github.com/og/thinking-in-go/ddd/goods/interface"
	shopInterface "github.com/og/thinking-in-go/ddd/shop/interface"
)

type Service struct {
	repo goodsInterface.Repo
	shop shopInterface.Service
}
func NewService(repo goodsInterface.Repo, shop shopInterface.Service) Service {
	return Service{
		repo: repo,
		shop: shop,
	}
}