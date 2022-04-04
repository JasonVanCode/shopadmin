package services

import (
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"shopadmin/models"
	"strings"
)

type CartService struct {
	BaseService
}

//下面是post请求数据
//添加购物车请求数据
type CartAddBody struct {
	GoodsId   int `json:"goodsId"`
	Number    int `json:"number"`
	ProductId int `json:"productId"`
}

//删除购物车请求的数据
type DelCartBody struct {
	ProductIds string `json:"productIds"`
}

//更新购物车请求数据
type UpdateCartBody struct {
	ProductId int `json:"productId"`
	GoodsId   int `json:"goodsId"`
	Number    int `json:"number"`
	Id        int `json:"id"`
}

//购物车是否选中状态请求数据
type CheckedCartBody struct {
	Id         int
	IsChecked  int
	ProductIds int
}

//下面是返回数据格式定义
//商品结算返回的数据结构
type CartAddress struct {
	*models.NideshopAddress
	ProvinceName string `json:"province_name"`
	CityName     string `json:"city_name"`
	DistrictName string `json:"district_name"`
	FullRegion   string `json:"full_region"`
}

type CheckoutRtnJson struct {
	Address          CartAddress                  `json:"checkedAddress"`
	FreightPrice     float64                      `json:"freightPrice"`
	CheckedCoupon    []*models.NideshopUserCoupon `json:"checkedCoupon"`
	CouponList       []*models.NideshopUserCoupon `json:"couponList"`
	CouponPrice      float64                      `json:"couponPrice"`
	CheckedGoodsList []*models.NideshopCart       `json:"checkedGoodsList"`
	GoodsTotalPrice  float64                      `json:"goodsTotalPrice"`
	OrderTotalPrice  float64                      `json:"orderTotalPrice"`
	ActualPrice      float64                      `json:"actualPrice"`
}

//根据id获取当前购物车
func (*CartService) GetCartById(id int) *models.NideshopCart {
	var cart models.NideshopCart
	o := orm.NewOrm()
	err := o.QueryTable(new(models.NideshopCart)).Filter("id", id).One(&cart)
	if err != nil {
		return nil
	}
	return &cart
}

//获取购物车信息
func (c *CartService) GetCartList(userId int, session string) ([]*models.NideshopCart, int, error) {
	var cartList []*models.NideshopCart
	o := orm.NewOrm()
	num, err := o.QueryTable(new(models.NideshopCart)).Filter("user_id", userId).Filter("session_id", session).All(&cartList)

	if err != nil {
		return nil, 0, err
	}
	return cartList, int(num), nil
}

//查看该用户购物车里面是否存在该商品
func (c *CartService) IsAddToCart(goodsId, productId, userId int) *models.NideshopCart {
	var cart models.NideshopCart
	o := orm.NewOrm()
	err := o.QueryTable(new(models.NideshopCart)).Filter("user_id", userId).Filter("goods_id", goodsId).Filter("product_id", productId).One(&cart)
	if err != nil || cart.GoodsId != goodsId {
		return nil
	}
	return &cart
}

//新建购物车数
func (*CartService) AddCart(cart *models.NideshopCart) error {
	o := orm.NewOrm()
	num, err := o.Insert(cart)
	if err != nil || num == 0 {
		return errors.New("数据插入失败")
	}
	return nil
}

//修改购物车商品数量
//isUpdate 0 新增  1 更新
func (*CartService) UpdateCartNumber(cart *models.NideshopCart, number int, isUpdate int) error {
	o := orm.NewOrm()
	if isUpdate == 1 {
		cart.Number = number
	} else {
		cart.Number = cart.Number + number
	}
	num, err := o.Update(cart)
	if err != nil || num == 0 {
		return errors.New("数据更新失败")
	}
	return nil
}

//删除购物车数据
func (*CartService) DelCart(productIds string, userId int) bool {
	pIds := strings.Split(productIds, ",")
	o := orm.NewOrm()
	num, err := o.QueryTable(new(models.NideshopCart)).Filter("product_id__in", pIds).Filter("user_id", userId).Delete()
	if err != nil || num == 0 {
		return false
	}
	return true
}

//更改购物车选中状态
func (c *CartService) UpdateCartChecked(id, isChecked int) error {
	nowCart := c.GetCartById(id)
	if nowCart == nil {
		return errors.New("该条数据不存在")
	}
	o := orm.NewOrm()
	nowCart.Checked = isChecked
	num, err := o.Update(nowCart, "checked")
	if err != nil || num == 0 {
		return errors.New("更新失败")
	}
	return nil
}

//获取选中的数据
func (c *CartService) GetCheckedCart(userId, sessionId int) []*models.NideshopCart {
	var cartList []*models.NideshopCart
	o := orm.NewOrm()
	num, err := o.QueryTable(new(models.NideshopCart)).Filter("checked", 1).Filter("user_id", userId).Filter("session_id", sessionId).All(&cartList)
	if err != nil || num == 0 {
		return nil
	}
	return cartList

}
