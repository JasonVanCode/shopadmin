package controllers

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"net/url"
	"shopadmin/middleware"
	"shopadmin/models"
	"shopadmin/services"
	"strconv"
)

const TimeFormat = "2006-01-02 15:04:05"

const TimeDateFormat = "2006-01-02"

type BaseController struct {
	beego.Controller
	user     *models.Nideshop_user
	pageData map[string]int
	getData  map[string]string
}

//list 返回的map数据
type ListData map[string]interface{}

//每次控制器处理预先处理的方法
func (b *BaseController) Prepare() {
	//查看是否有id
	if middleware.UserId > 0 {
		b.user, _ = services.GetNideshop_userById(middleware.UserId)
		fmt.Println("获取到用户的信息是：", b.user)
	}

	//分页请求数据处理
	if b.Ctx.Input.IsGet() {
		if value, err := url.ParseQuery(b.Ctx.Request.URL.RawQuery); err == nil {
			b.pageData = make(map[string]int)
			b.getData = make(map[string]string)
			for queryK, querV := range value {
				if queryK == "page" {
					b.pageData["page"], _ = strconv.Atoi(querV[0])
				}
				if queryK == "size" {
					b.pageData["size"], _ = strconv.Atoi(querV[0])
				}
				if queryK != "page" && queryK != "size" {
					b.getData[queryK] = querV[0]
				}
			}

		}
	}

}

//获取登录用户的id
func (b *BaseController) getUserId() int {
	if b.user == nil || b.user.Id == 0 {
		return 0
	}
	return b.user.Id
}

//获取用户的WeixinOpenid
func (b *BaseController) getUserOpenId() string {
	return b.user.WeixinOpenid
}
