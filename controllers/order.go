package controllers

import (
	"encoding/json"
	"fmt"
	"shopadmin/models"
	"shopadmin/services"
	"shopadmin/utils"
	"time"
)

type OrderController struct {
	BaseController
}

//获取订单列表数据
func (o *OrderController) GetOrderList() {
	userId := o.getUserId()
	var orderService = new(services.OrderService)
	orderList := orderService.GetUserOrderList(userId, o.pageData)
	if orderList == nil {
		utils.HttpSuccess(nil, o.Ctx)
	}
	orderIds := orderService.GetOrderIds(orderList)
	orderGoodsLists := orderService.GetOrderGoods(orderIds)
	var returnData []services.OrderListRtnJson
	for _, orderItem := range orderList {
		var orderGoods []*models.NideshopOrderGoods
		var goosCount int
		for _, goodsItem := range orderGoodsLists {
			if orderItem.Id == goodsItem.OrderId {
				orderGoods = append(orderGoods, goodsItem)
				goosCount += goodsItem.Number
			}
		}
		returnData = append(returnData, services.OrderListRtnJson{
			orderItem,
			orderGoods,
			goosCount,
			orderItem.GetOrderStatusText(),
			orderItem.GetOrderHandleOptions(),
		})
	}
	utils.HttpSuccess(returnData, o.Ctx)
}

//结算订单
func (o *OrderController) SubmitOrder() {
	userId := o.getUserId()
	var submitBody services.OrderSUbmitBody
	if err := json.Unmarshal(o.Ctx.Input.RequestBody, &submitBody); err != nil {
		utils.HttpFail(400, "请求参数有误", o.Ctx)
	}
	//收获地址
	var addressService = new(services.AddressService)
	addressInfo := addressService.GetAddressById(submitBody.AddressId)
	if addressInfo == nil {
		utils.HttpFail(400, "请选择收获地址", o.Ctx)
	}
	//购买的商品
	var cartService services.CartService
	cartList := cartService.GetCheckedCart(userId, 1)
	if cartList == nil {
		utils.HttpFail(400, "请选择商品", o.Ctx)
	}
	//商品总价格
	var goodstotalprice float64
	for _, v := range cartList {
		goodstotalprice += v.RetailPrice * float64(v.Number)
	}
	//邮费
	var freightPrice float64 = 0.0
	//优惠价格
	var couponPrice float64 = 0.0
	//实际支付价格
	var actualPrice = goodstotalprice + freightPrice - couponPrice
	fmt.Println(actualPrice)
	//添加订单信息
	var orderService = new(services.OrderService)
	orderInfo := models.NideshopOrder{
		OrderSn:      orderService.GenerateOrderNumber(),
		UserId:       userId,
		Consignee:    addressInfo.Name,
		Province:     addressInfo.ProvinceId,
		City:         addressInfo.CityId,
		District:     addressInfo.DistrictId,
		Address:      addressInfo.Address,
		Mobile:       addressInfo.Mobile,
		FreightPrice: freightPrice,
		CouponId:     0,
		CouponPrice:  couponPrice,
		GoodsPrice:   goodstotalprice,
		OrderPrice:   goodstotalprice + freightPrice,
		ActualPrice:  actualPrice,
		AddTime:      time.Now(),
		ConfirmTime:  time.Now(),
		Postscript:   "",
	}
	order, err := orderService.OrderSave(orderInfo, cartList)
	if err == nil {
		utils.HttpSuccess(struct {
			OrderInfo *models.NideshopOrder `json:"orderInfo"`
		}{
			order,
		}, o.Ctx)
	}
	utils.HttpFail(400, err.Error(), o.Ctx)
}

//获取订单详情
func (o *OrderController) GetOrderDetail() {
	orderId, _ := o.GetInt("orderId")
	var orderService = new(services.OrderService)
	orderInfo := orderService.GetOrderById(orderId)
	if orderInfo == nil {
		utils.HttpFail(400, "该订单不存在", o.Ctx)
	}
	orderGoodsList := orderService.GetOrderGoods([]int{orderId})
	option := orderInfo.GetOrderHandleOptions()

	result := services.OrderDetailRtnJson{
		orderInfo,
		orderGoodsList,
		option,
	}
	utils.HttpSuccess(result, o.Ctx)
}
