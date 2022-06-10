package controllers

import (
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"

	"shopadmin/middleware"
	"shopadmin/models"
	"shopadmin/services"
	"shopadmin/utils"
)

type AdminBaseController struct {
	beego.Controller
	userInfo       *models.AdminUser
	pagination     map[string]int    //分页的数据
	otherQueryData map[string]string //其他get请求数据
}

func (base *AdminBaseController) Prepare() {
	//用户id
	userId := middleware.AdminUserId
	if userId > 0 {
		base.userInfo, _ = base.GetUserInfo(middleware.AdminUserId)
	}
	//get请求
	if base.Ctx.Input.IsGet() {
		requestUrl := base.Ctx.Request.URL.Query()
		base.pagination = make(map[string]int)
		base.otherQueryData = make(map[string]string)
		for k, v := range requestUrl {
			if k == "page" {
				base.pagination["page"] = utils.TransStringToInt(v[0])
			} else if k == "size" {
				base.pagination["size"] = utils.TransStringToInt(v[0])
			} else {
				if v[0] != "" {
					base.otherQueryData[k] = v[0]
				}
			}
		}
	}
}

func (base *AdminBaseController) GetUserId() int {
	return base.userInfo.Id
}

func (*AdminBaseController) GetUserInfo(id int) (*models.AdminUser, error) {
	var adminUserService services.AdminUserService
	return adminUserService.GetUserById(id)
}

//错误日志处理
func (*AdminBaseController) ErrLog(err error) {
	logs.Error(err.Error() + "\n")
}
