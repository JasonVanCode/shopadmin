package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

//商品对应规格表值表
type NideshopOrder struct {
	Id             int       `orm:"column(id)" json:"id"`
	OrderSn        string    `orm:"column(order_sn)" json:"order_sn"`
	UserId         int       `orm:"column(user_id)" json:"user_id"`
	OrderStatus    int       `orm:"column(order_status)" json:"order_status"`
	ShippingStatus int       `orm:"column(shipping_status)" json:"shipping_status"`
	PayStatus      int       `orm:"column(pay_status)" json:"pay_status"`
	Consignee      string    `orm:"column(consignee);description(收货人)" json:"consignee"`
	Country        int       `orm:"column(country)" json:"country"`
	Province       int       `orm:"column(province)" json:"province"`
	City           int       `orm:"column(city)" json:"city"`
	District       int       `orm:"column(district);description(地区)" json:"district"`
	Address        string    `orm:"column(address)" json:"address"`
	Mobile         string    `orm:"column(mobile)" json:"mobile"`
	Postscript     string    `orm:"column(postscript)" json:"postscript"`
	ShippingFee    float64   `orm:"column(shipping_fee);digits(10);decimals(2);description(运费)" json:"shipping_fee"`
	PayName        string    `orm:"column(pay_name)" json:"pay_name"`
	PayId          int       `orm:"column(pay_id)" json:"pay_id"`
	ActualPrice    float64   `orm:"column(actual_price);digits(10);decimals(2);description(实际需要支付的金额)" json:"actual_price"`
	Integral       int       `orm:"column(integral)" json:"integral"`
	IntegralMoney  float64   `orm:"column(integral_money);digits(10);decimals(2)" json:"integral_money"`
	OrderPrice     float64   `orm:"column(order_price);digits(10);decimals(2);description(订单总价)" json:"order_price"`
	GoodsPrice     float64   `orm:"column(goods_price);digits(10);decimals(2);description(商品总价)" json:"goods_price"`
	AddTime        time.Time `orm:"column(add_time)" json:"add_time"`
	ConfirmTime    time.Time `orm:"column(confirm_time)" json:"confirm_time"`
	PayTime        time.Time `orm:"column(pay_time)" json:"pay_time"`
	FreightPrice   float64   `orm:"column(freight_price);digits(10);decimals(2);description(配送费用)" json:"freight_price"`
	CouponId       int       `orm:"column(coupon_id);description(使用的优惠券id)" json:"coupon_id"`
	ParentId       int       `orm:"column(parent_id)" json:"parent_id"`
	CouponPrice    float64   `orm:"column(coupon_price);digits(10);decimals(2)" json:"coupon_price"`
	CallbackStatus int       `orm:"column(callback_status)" json:"callback_status"`
}

type OrderHandleOption struct {
	Cancel   bool `json:"cancel"`
	Delete   bool `json:"delete"`
	Pay      bool `json:"pay"`
	Comment  bool `json:"comment"`
	Delivery bool `json:"delivery"`
	Confirm  bool `json:"confirm"`
	Return   bool `json:"return"`
	Buy      bool `json:"buy"`
}

func (*NideshopOrder) TableName() string {
	return "nideshop_order"
}

func (o *NideshopOrder) PaySuccess() error {
	obj := orm.NewOrm()
	o.PayStatus = 1
	o.PayTime = time.Now()
	_, err := obj.Update(o, "pay_status", "pay_time")
	if err != nil {
		return err
	}
	return nil
}

//获取订单状态信息
func (o *NideshopOrder) GetOrderStatusText() string {
	var text string
	switch o.OrderStatus {
	case 0:
		text = "未付款"
	case 201:
		text = "已付款"
	}
	return text
}

//获取订单状态
func (o *NideshopOrder) GetOrderHandleOptions() OrderHandleOption {
	// 订单流程：下单成功－》支付订单－》发货－》收货－》评论
	// 订单相关状态字段设计，采用单个字段表示全部的订单状态
	// 1xx表示订单取消和删除等状态 0订单创建成功等待付款，101订单已取消，102订单已删除
	// 2xx表示订单支付状态,201订单已付款，等待发货
	// 3xx表示订单物流相关状态,300订单已发货，301用户确认收货
	// 4xx表示订单退换货相关的状态,401没有发货，退款402,已收货，退款退货
	// 如果订单已经取消或是已完成，则可删除和再次购买
	handleOption := OrderHandleOption{
		false,
		false,
		false,
		false,
		false,
		false,
		false,
		false,
	}
	switch o.OrderStatus {
	case 0:
		handleOption.Cancel = true
		handleOption.Pay = true
	case 101:
		handleOption.Delete = true
		handleOption.Buy = true
	case 201:
		handleOption.Return = true
	case 300:
		handleOption.Cancel = true
		handleOption.Buy = true
		handleOption.Return = true
	case 301:
		handleOption.Delete = true
		handleOption.Buy = true
		handleOption.Comment = true
	}
	return handleOption
}

func init() {
	orm.RegisterModel(new(NideshopOrder))
}
