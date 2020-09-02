package merchantService

import (
	merchantDTS "github.com/og/thinking-in-go/ddd/merchant/dts"
	merchantRepoDTS "github.com/og/thinking-in-go/ddd/merchant/repo/dts"
	requestUtil "github.com/og/thinking-in-go/ddd/util/request"
	resUtil "github.com/og/thinking-in-go/ddd/util/response"
	ge "github.com/og/x/error"
)


func (dep Service) CreateMerchant(req merchantDTS.ReqCreateMerchant) (reply merchantDTS.ReplyCreateMerchant, reject ge.Reject) {
	/* 格式校验 */{
		reject = requestUtil.Check(req) ; if reject.Fail() {return}
	}
	/* 合法性校验 */{/* 暂无 */}
	/* 数据存储 */{
		_, hasMerchant := dep.repo.MerchantByName(req.Name)
		if hasMerchant {
			return reply, resUtil.RejectFail("商家（" + req.Name + ")名已存在")
		}
		merchant := dep.repo.CreateMerchant(merchantRepoDTS.CreateMerchantData{
			Name: req.Name,
		})
		reply.MerchantID = merchant.ID
	}
	return
}