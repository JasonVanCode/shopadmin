package controllers

import (
	"encoding/json"
	"fmt"
	"shopadmin/models"
	"shopadmin/services"
	"shopadmin/utils"
	"strings"
)

type CartController struct {
	BaseController
}

type CartTotal struct {
	CartNum            int     `json:"cartNum"`
	GoodsCount         int     `json:"goodsCount"`
	GoodsAmount        float64 `json:"goodsAmount"`
	CheckedGoodsCount  int     `json:"checkedGoodsCount"`
	CheckedGoodsAmount float64 `json:"checkedGoodsAmount"`
	CheckedCartNum     int     `json:"checkedCartNum"`
}

type IndexData struct {
	CartTotal `json:"cartTotal"`
	CartList  []*models.NideshopCart `json:"cartList"`
}

//列表数据
func (c *CartController) CartIndex() {
	utils.HttpSuccess(c.GetCartList(), c.Ctx)
}

//获取购物车
func (c *CartController) GetCartList() IndexData {
	userId := c.getUserId()
	var cartService services.CartService
	cartList, num, _ := cartService.GetCartList(userId, "1")
	var goodsCount int
	var goodsAmount float64
	var checkedGoodsCount int
	var CheckedGoodsAmount float64
	var checkedCartNum int
	for _, v := range cartList {
		goodsCount += v.Number
		goodsAmount += float64(v.Number) * v.RetailPrice
		if v.Checked == 1 {
			checkedCartNum += 1
			checkedGoodsCount += v.Number
			CheckedGoodsAmount += float64(v.Number) * v.RetailPrice
		}
	}
	return IndexData{
		CartTotal{
			num,
			goodsCount,
			goodsAmount,
			checkedGoodsCount,
			CheckedGoodsAmount,
			checkedCartNum,
		},
		cartList,
	}
}

//添加购物车
func (c *CartController) AddCart() {
	var err error
	var body services.CartAddBody
	var userId = 2
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &body)
	if err != nil || body.GoodsId == 0 || body.ProductId == 0 || body.Number == 0 {
		utils.HttpFail(400, "请求数据错误", c.Ctx)
	}
	//判断商品是否存在
	goodsService := new(services.GoodsService)
	goods := goodsService.GetGoodsById(body.GoodsId)
	if goods == nil || goods.IsDelete == 1 {
		utils.HttpFail(400, "请求数据错误", c.Ctx)
	}
	//商品下的单独产品
	productService := new(services.ProductService)
	product := productService.GetProductById(body.ProductId)
	if product == nil || product.GoodsNumber < body.Number {
		utils.HttpFail(400, "库存不足", c.Ctx)
	}
	//判断该用户是否添加该产品
	var cart *models.NideshopCart
	cartService := new(services.CartService)
	cart = cartService.IsAddToCart(body.GoodsId, body.ProductId, userId)

	//产品规格信息
	var values string
	//购物车没该产品
	if cart == nil {
		SpecificationIds := product.GoodsSpecificationIds
		if SpecificationIds != "" {
			//var intgoodsspecificationids []int
			goodsspecificationids := strings.Split(SpecificationIds, "_")
			var specificationService = new(services.GoodsSpecificationService)
			specificationData := specificationService.GetGoodsSpecification(body.GoodsId, goodsspecificationids)
			//获取产品规格拼接的数据
			values = specificationService.HandleSpecificationValues(specificationData)
		}
		var cartData = models.NideshopCart{
			UserId:                    userId,
			SessionId:                 "1",
			GoodsId:                   goods.Id,
			GoodsSn:                   product.GoodsSn,
			ProductId:                 product.Id,
			GoodsName:                 goods.Name,
			RetailPrice:               product.RetailPrice,
			MarketPrice:               product.RetailPrice,
			Number:                    body.Number,
			GoodsSpecifitionIds:       product.GoodsSpecificationIds,
			GoodsSpecifitionNameValue: values,
			Checked:                   1,
			ListPicUrl:                goods.ListPicUrl,
		}
		if err := cartService.AddCart(&cartData); err != nil {
			utils.HttpFail(400, "添加购物车失败", c.Ctx)
		}

	} else {
		//判断添加的数量是否超
		if cart.Number+body.Number > product.GoodsNumber {
			utils.HttpFail(400, "库存不足", c.Ctx)
		}

		if err := cartService.UpdateCartNumber(cart, body.Number, 0); err != nil {
			utils.HttpFail(400, "添加购物车失败", c.Ctx)
		}

	}
	utils.HttpSuccess(c.GetCartList(), c.Ctx)
}

//获取购物车数量
func (c *CartController) GetCartCount() {
	userId := c.getUserId()
	var cartService services.CartService
	_, num, _ := cartService.GetCartList(userId, "1")
	utils.HttpSuccess(struct {
		CartNum int `json:"cartNum"`
	}{
		num,
	}, c.Ctx)
	c.Ctx.WriteString("0")
}

//更新购物车信息
func (c *CartController) UpdateCart() {
	var updateBody services.UpdateCartBody
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &updateBody); err != nil {
		utils.HttpFail(400, "请求参数有误", c.Ctx)
	}
	var productService services.ProductService
	product := productService.GetProductById(updateBody.ProductId)
	if product == nil || product.GoodsNumber < updateBody.Number {
		utils.HttpFail(400, "库存不足", c.Ctx)
	}

	var cartService services.CartService
	cart := cartService.GetCartById(updateBody.Id)
	if err := cartService.UpdateCartNumber(cart, updateBody.Number, 1); err == nil {
		utils.HttpSuccess(c.GetCartList(), c.Ctx)
	}
	utils.HttpFail(400, "修改失败", c.Ctx)
}

//删除购物车信息
func (c *CartController) DeleteCart() {
	userId := c.getUserId()
	var delBody services.DelCartBody
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &delBody); err != nil || delBody.ProductIds == "" {
		utils.HttpFail(400, "请求参数有问题", c.Ctx)
	}
	var cartservice = new(services.CartService)
	if res := cartservice.DelCart(delBody.ProductIds, userId); res {
		utils.HttpSuccess(c.GetCartList(), c.Ctx)
	} else {
		utils.HttpFail(400, "删除失败", c.Ctx)
	}
}

//购物车是否选中状态
func (c *CartController) ChekcedCart() {
	var checkedBody services.CheckedCartBody
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &checkedBody); err != nil || checkedBody.Id == 0 {
		utils.HttpFail(400, "请求参数有问题", c.Ctx)
	}
	var cartservice = new(services.CartService)
	err := cartservice.UpdateCartChecked(checkedBody.Id, checkedBody.IsChecked)
	if err == nil {
		utils.HttpSuccess(c.GetCartList(), c.Ctx)
	} else {
		utils.HttpFail(400, err.Error(), c.Ctx)
	}
}

//结算
func (c *CartController) CheckOutCart() {
	userId := c.getUserId()
	addressId, _ := c.GetInt("addressId")
	var address *models.NideshopAddress
	var addressService = new(services.AddressService)
	//取默认地址
	if addressId == 0 {
		address = addressService.GetDefaultAddress(userId)
	} else {
		address = addressService.GetAddressById(addressId)
	}
	if address == nil {
		utils.HttpFail(400, "请填写购物地址", c.Ctx)
	}
	//获取名字
	PCDMapNames := address.GetPCDNames()
	addressInfo := services.CartAddress{
		address,
		PCDMapNames[address.ProvinceId],
		PCDMapNames[address.CityId],
		PCDMapNames[address.DistrictId],
		PCDMapNames[address.ProvinceId] + PCDMapNames[address.CityId] + PCDMapNames[address.DistrictId],
	}
	fmt.Println(addressInfo)
	//选中的商品
	var cartService = new(services.CartService)
	cartList := cartService.GetCheckedCart(userId, 1)
	if cartList == nil {
		utils.HttpFail(400, "请选择要结算的商品", c.Ctx)
	}
	//商品总价格
	var goodstotalprice float64
	for _, v := range cartList {
		goodstotalprice += v.RetailPrice * float64(v.Number)
	}
	//获取该用户的优惠券
	var userCouponService = new(services.UserCouponService)
	couponList := userCouponService.GetUnusedCoupon(userId)
	//邮费
	var freightPrice float64 = 0.0
	//优惠价格
	var couponPrice float64 = 0.0
	//实际支付价格
	var actualPrice = goodstotalprice + freightPrice - couponPrice
	utils.HttpSuccess(services.CheckoutRtnJson{
		Address:          addressInfo,
		FreightPrice:     freightPrice,
		CouponList:       couponList,
		CouponPrice:      couponPrice,
		CheckedGoodsList: cartList,
		GoodsTotalPrice:  goodstotalprice,
		//OrderTotalPrice:actualPrice,
		ActualPrice: actualPrice,
	}, c.Ctx)

}
