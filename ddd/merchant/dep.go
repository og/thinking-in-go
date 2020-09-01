package merchantService

import (
	"github.com/og/thinking-in-go/ddd/checkreq"
	"github.com/og/thinking-in-go/ddd/merchant/repo/interface"
)

type Service struct {
	repo     merchantRepoInterface.Repo
	checkReq checkreq.CheckReq
}
func NewService(repo merchantRepoInterface.Repo) Service {
	return Service{
		repo:    repo,
		checkReq: checkreq.CheckReq{},
	}
}