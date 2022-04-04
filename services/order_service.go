package services

import (
	"github.com/beego/beego/v2/client/orm"
	"math/rand"
	"shopadmin/models"
	"shopadmin/utils"
	"time"
)

type OrderService struct {
	BaseService
}

//提交订单的请求数据

type OrderSUbmitBody struct {
	AddressId int
	CouponId  int
}

//orderlist 返回的数据
type OrderListRtnJson struct {
	*models.NideshopOrder
	GoodsList       []*models.NideshopOrderGoods `json:"goodsList"`
	GoodsCount      int                          `json:"goodsCount"`
	OrderStatusText string                       `json:"order_status_text"`
	HandOption      models.OrderHandleOption     `json:"handleOption"`
}

//order detail 返回数据
type OrderDetailRtnJson struct {
	OrderInfo    *models.NideshopOrder        `json:"orderInfo"`
	OrderGoods   []*models.NideshopOrderGoods `json:"orderGoods"`
	HandleOption models.OrderHandleOption     `json:"handleOption"`
}

//获取当前用户的订单数据
func (*OrderService) GetUserOrderList(userId int, pageData map[string]int) []*models.NideshopOrder {
	var orderList []*models.NideshopOrder
	o := orm.NewOrm()
	num, err := o.QueryTable(new(models.NideshopOrder)).Filter("user_id", userId).All(&orderList)
	if err != nil || num == 0 {
		return nil
	}
	return orderList
}

//根据id获取订单
func (*OrderService) GetOrderById(id int) *models.NideshopOrder {
	var order models.NideshopOrder
	o := orm.NewOrm()
	err := o.QueryTable(new(models.NideshopOrder)).Filter("id", id).One(&order)
	if err != nil {
		return nil
	}
	return &order

}

//生成订单号
func (*OrderService) GenerateOrderNumber() string {
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(10000) + 100000
	randNumString := utils.TransIntToString(randNum)
	return utils.GetNowYear() + utils.GetNowMonth() + utils.GetNowDay() + utils.GetNowHour() + utils.GetNowMinute() + utils.GetNowSecond() + randNumString
}

//保存订单信息
func (*OrderService) OrderSave(orderInfo models.NideshopOrder, cartList []*models.NideshopCart) (*models.NideshopOrder, error) {
	o := orm.NewOrm()
	//事务处理
	to, err := o.Begin()
	if err != nil {
		return nil, err
	}

	id, err := to.Insert(&orderInfo)
	if err != nil || id == 0 {
		to.Rollback()
		return nil, err
	}
	//商品订单数据
	var orderGoods []models.NideshopOrderGoods
	for _, v := range cartList {
		orderGoods = append(orderGoods, models.NideshopOrderGoods{
			OrderId:                   int(id),
			GoodsId:                   v.GoodsId,
			GoodsSn:                   v.GoodsSn,
			ProductId:                 v.ProductId,
			GoodsName:                 v.GoodsName,
			Number:                    v.Number,
			MarketPrice:               v.MarketPrice,
			RetailPrice:               v.RetailPrice,
			GoodsSpecifitionNameValue: v.GoodsSpecifitionNameValue,
			GoodsSpecifitionIds:       v.GoodsSpecifitionIds,
			ListPicUrl:                v.ListPicUrl,
		})
	}
	_, err = to.InsertMulti(100, &orderGoods)
	if err != nil {
		to.Rollback()
		return nil, err
	}
	//提交失误
	to.Commit()
	return &orderInfo, nil
}

//获取order  对应 order_goods 所有的id
func (*OrderService) GetOrderIds(orderList []*models.NideshopOrder) []int {
	var orderIds []int
	for _, v := range orderList {
		orderIds = append(orderIds, v.Id)
	}
	return orderIds
}

func (*OrderService) GetOrderGoods(orderIds []int) []*models.NideshopOrderGoods {
	var orderGoodsList []*models.NideshopOrderGoods
	o := orm.NewOrm()
	num, err := o.QueryTable(new(models.NideshopOrderGoods)).Filter("order_id__in", orderIds).All(&orderGoodsList)
	if err != nil || num == 0 {
		return nil
	}
	return orderGoodsList
}
