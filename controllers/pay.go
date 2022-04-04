package controllers

import (
	"fmt"
	"shopadmin/services"
	"shopadmin/utils"
)

type PayController struct {
	BaseController
}

func (p *PayController) PrePay() {
	orderId, _ := p.GetInt("orderId")
	fmt.Println(orderId)
	var orderService = new(services.OrderService)
	order := orderService.GetOrderById(orderId)
	if order == nil {
		utils.HttpFail(400, "订单被取消了", p.Ctx)
	}
	if order.PayStatus != 0 {
		utils.HttpFail(400, "该订单已经支付，请勿重复支付", p.Ctx)
	}
	wxOpenId := p.getUserOpenId()
	if wxOpenId == "" {
		utils.HttpFail(400, "支付失败", p.Ctx)
	}
	//手动支付
	err := order.PaySuccess()
	if err == nil {
		utils.HttpSuccess(nil, p.Ctx)
	}
	utils.HttpFail(400, "支付失败", p.Ctx)
}
