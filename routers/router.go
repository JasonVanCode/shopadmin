// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"shopadmin/controllers"
	"shopadmin/middleware"
)

func init() {
	//登录的验证
	middleware.Auth()
	ns := beego.NewNamespace("/api",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSRouter("/getUser", &controllers.UserController{}, "get:Get"),
		),
		//登录验证
		beego.NSNamespace("/auth",
			beego.NSRouter("/loginByWeixin", &controllers.AuthController{}, "post:LoginByWeixin"),
		),
		//首页数据
		beego.NSNamespace("/index",
			beego.NSRouter("/index", &controllers.IndexController{}, "get:Index"),
		),

		//商品
		beego.NSNamespace("/goods",
			beego.NSRouter("/count", &controllers.GoodsController{}, "get:GetCount"),
			//分类数据
			beego.NSRouter("/category", &controllers.GoodsController{}, "get:GetCategory"),
			beego.NSRouter("/list", &controllers.GoodsController{}, "get:GetGoodsList"),
			//商品详情
			beego.NSRouter("/detail", &controllers.GoodsController{}, "get:GetGoodsDetail"),
			//相关商品
			beego.NSRouter("/related", &controllers.GoodsController{}, "get:GetRelatedGoods"),
		),

		//专题
		beego.NSNamespace("/topic",
			beego.NSRouter("/list", &controllers.TopicController{}, "get:GetTopicList"),
		),

		//商品分类
		beego.NSNamespace("/catalog",
			beego.NSRouter("/index", &controllers.CategoryController{}, "get:GetCategoryList"),
		),

		//购物车
		beego.NSNamespace("/cart",
			beego.NSRouter("/index", &controllers.CartController{}, "get:CartIndex"),
			beego.NSRouter("/add", &controllers.CartController{}, "post:AddCart"),
			beego.NSRouter("/goodscount", &controllers.CartController{}, "get:GetCartCount"),
			beego.NSRouter("/update", &controllers.CartController{}, "post:UpdateCart"),
			beego.NSRouter("/delete", &controllers.CartController{}, "post:DeleteCart"),
			beego.NSRouter("/checkout", &controllers.CartController{}, "get:CheckOutCart"),
			beego.NSRouter("/checked", &controllers.CartController{}, "post:ChekcedCart"),
		),

		//订单
		beego.NSNamespace("/order",
			beego.NSRouter("/list", &controllers.OrderController{}, "get:GetOrderList"),
			//提交订单
			beego.NSRouter("/submit", &controllers.OrderController{}, "post:SubmitOrder"),
			//订单详情
			beego.NSRouter("/detail", &controllers.OrderController{}, "get:GetOrderDetail"),
		),

		//地区区域
		beego.NSNamespace("/region",
			beego.NSRouter("/list", &controllers.RegionController{}, "get:GetRegionList"),
		),

		//地址维护
		beego.NSNamespace("address",
			beego.NSRouter("/list", &controllers.AddressController{}, "get:GetAddressList"),
			beego.NSRouter("/save", &controllers.AddressController{}, "post:SaveAddress"),
			beego.NSRouter("/delete", &controllers.AddressController{}, "post:DeleteAddress"),
		),

		//商品浏览足迹
		beego.NSNamespace("/footprint",
			beego.NSRouter("/list", &controllers.FootprintController{}, "get:GetFootprintList"),
		),

		//评价
		beego.NSNamespace("/comment",
			beego.NSRouter("/list", &controllers.CommentController{}, "get:GetCommentList"),
			beego.NSRouter("/count", &controllers.CommentController{}, "get:GetCommentCount"),
		),

		//支付
		beego.NSNamespace("/pay",
			beego.NSRouter("/prepay", &controllers.PayController{}, "get:PrePay"),
		),
	)
	beego.AddNamespace(ns)
}
