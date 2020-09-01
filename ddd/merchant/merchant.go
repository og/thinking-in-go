package merchantService

import (
	"errors"
	merchantDTS "github.com/og/thinking-in-go/ddd/merchant/dts"
	merchantRepo "github.com/og/thinking-in-go/ddd/merchant/repo"
)


func (dep Service) CreateMerchant(req merchantDTS.ReqCreateMerchant) (reply merchantDTS.ReplyCreateMerchant, reject error) {
	/* 格式校验 */{
		reject = dep.checkReq.Check(req) ; if reject != nil { return }
	}
	/* 合法性校验 */{/* 暂无 */}
	/* 数据存储 */{
		_, hasMerchant := dep.repo.MerchantByName(req.Name)
		if hasMerchant {
			return reply, errors.New("商家(" + req.Name + ")已存在")
		}
		merchant := dep.repo.CreateMerchant(merchantRepo.CreateMerchantData{
			Name: req.Name,
		})
		reply.MerchantID = merchant.ID
	}
	return
}