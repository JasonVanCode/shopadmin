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
	middleware.AdminAuth()
	//小程序接口管理
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

		//搜索
		beego.NSNamespace("/search",
			beego.NSRouter("/index", &controllers.SearchController{}, "get:GetIndex"),
			beego.NSRouter("/helper", &controllers.SearchController{}, "get:SearchHelper"),
		),
	)

	//商城后台管理系统接口

	ns2 := beego.NewNamespace("/admin",
		//登录验证
		beego.NSNamespace("/auth",
			beego.NSRouter("/login", &controllers.AdminAuthController{}, "post:Login"),
		),
		//首页数据
		beego.NSNamespace("/index",
			beego.NSRouter("/index", &controllers.AdminIndexController{}, "get:GetIndex"),
		),
		//用户列表
		beego.NSNamespace("/user",
			beego.NSRouter("/index", &controllers.AdminUserController{}, "get:GetIndex"),
			beego.NSRouter("/del", &controllers.AdminUserController{}, "post:DelUser"),
			beego.NSRouter("/createOrUpdate", &controllers.AdminUserController{}, "post:CreateOrUpdate"),
			beego.NSRouter("/changeStatus", &controllers.AdminUserController{}, "post:ChangeStatus"),
			beego.NSRouter("/changeRole", &controllers.AdminUserController{}, "post:ChangeRole"),
		),
		//角色列表
		beego.NSNamespace("/role",
			beego.NSRouter("/index", &controllers.AdminRoleController{}, "get:GetIndex"),
			beego.NSRouter("/listIdName", &controllers.AdminRoleController{}, "get:GetRoleListIdName"),
			beego.NSRouter("/createOrUpdate", &controllers.AdminRoleController{}, "post:CreateOrUpdate"),
			beego.NSRouter("/changeStatus", &controllers.AdminRoleController{}, "post:ChangeStatus"),
			beego.NSRouter("/del", &controllers.AdminRoleController{}, "post:DelRole"),
			beego.NSRouter("/getRole", &controllers.AdminRoleController{}, "get:GetRoleInfo"),
			//角色菜单分配
			beego.NSRouter("/allocMenus", &controllers.AdminRoleController{}, "post:AllocMenus"),
		),
		//菜单列表
		beego.NSNamespace("/menu",
			beego.NSRouter("/index", &controllers.AdminMenuController{}, "get:GetIndex"),
			beego.NSRouter("/listIdName", &controllers.AdminMenuController{}, "get:GetMenuListIdName"),
			beego.NSRouter("/createOrUpdate", &controllers.AdminMenuController{}, "post:CreateOrUpdate"),
			//beego.NSRouter("/changeStatus", &controllers.AdminMenuController{}, "post:ChangeStatus"),
			beego.NSRouter("/del", &controllers.AdminMenuController{}, "post:DelMenu"),
			beego.NSRouter("/getAllocMenuLists", &controllers.AdminMenuController{}, "get:GetAllocMenuLists"),
		),
		//商品模块
		beego.NSNamespace("/goods",
			beego.NSRouter("/index", &controllers.AdminGoodsController{}, "get:GetIndex"),
			//修改商品状态
			beego.NSRouter("/changeStatus", &controllers.AdminGoodsController{}, "post:ChangeGoodsStatus"),

			beego.NSNamespace("/brand",
				//商品品牌
				beego.NSRouter("/index", &controllers.AdminBrandController{}, "get:GetIndex"),
				//添加品牌
				beego.NSRouter("/createOrUpdate", &controllers.AdminBrandController{}, "post:CreateOrUpdate"),
				beego.NSRouter("/getBrand", &controllers.AdminBrandController{}, "get:GetEdit"),
				beego.NSRouter("/listIdName", &controllers.AdminBrandController{}, "get:GetBrandIdNames"),
			),
			beego.NSNamespace("/category",
				beego.NSRouter("/index", &controllers.AdminCategoryController{}, "get:GetIndex"),
				//分类属性
				beego.NSRouter("/attribute", &controllers.AdminCategoryController{}, "get:GetCategoryAttributeLists"),
				//获取顶级分类
				beego.NSRouter("/listIdName", &controllers.AdminCategoryController{}, "get:GetCategoryIdName"),
				beego.NSRouter("/createOrUpdate", &controllers.AdminCategoryController{}, "post:CreateOrUpdate"),
				beego.NSRouter("/getCategory", &controllers.AdminCategoryController{}, "get:GetEdit"),
				//删除分类
				beego.NSRouter("/del", &controllers.AdminCategoryController{}, "post:DelCategory"),
				//获取分类以及子类的数据
				beego.NSRouter("/withChildren", &controllers.AdminCategoryController{}, "get:GetCategoryWithChildren"),
			),
			//属性
			beego.NSNamespace("/attribute",
				//添加品牌
				beego.NSRouter("/createOrUpdate", &controllers.AdminCategoryController{}, "post:AttributeCreateOrUpdate"),
				//删除
				beego.NSRouter("/del", &controllers.AdminCategoryController{}, "post:DelAttribute"),
			),
		),
		//文件上传处理
		beego.NSNamespace("/upload",
			beego.NSRouter("/uploadSinglePic", &controllers.UploadController{}, "post:UploadSinglePic"),
		),
	)
	beego.AddNamespace(ns, ns2)
}
