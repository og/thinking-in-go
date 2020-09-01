package shopService

import (
	shopInterface "github.com/og/thinking-in-go/ddd/shop/interface"
)

type Service struct {
	repo shopInterface.Repo
}
func NewService(repo shopInterface.Repo) shopInterface.Service {
	return Service{
		repo: repo,
	}
}