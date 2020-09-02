package merchantService

import (
	"github.com/og/thinking-in-go/ddd/req"
	"github.com/og/thinking-in-go/ddd/merchant/repo/interface"
	"github.com/og/thinking-in-go/ddd/util/request"
)

type Service struct {
	repo     merchantRepoInterface.Repo
	requestUtil request.RequestUtil
}
func NewService(repo merchantRepoInterface.Repo) Service {
	return Service{
		repo:     repo,
		requestUtil: request.RequestUtil{},
	}
}