package services

import (
	"encoding/base64"
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/validation"
	"shopadmin/models"
	"shopadmin/utils"
)

type AdminAuthService struct {
	BaseService
}

//登录请求的数据
type LoginBody struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

func (u *LoginBody) Valid(v *validation.Validation) {
	if u.UserName == "" {
		v.SetError("用户名", "不能为空")
	}

	if u.Password == "" {
		v.SetError("密码", "不能为空")
	}
}

//登录验证
func (*AdminAuthService) Login(data LoginBody) (user *models.AdminUser, err error) {
	var userInfo models.AdminUser
	userName := data.UserName
	o := orm.NewOrm()
	err = o.QueryTable(new(models.AdminUser)).Filter("username", userName).One(&userInfo)
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	//判断账号是否禁用
	if userInfo.Status == 0 {
		return nil, errors.New("账号被禁用")
	}
	newPass := data.Password
	pByte, err := base64.StdEncoding.DecodeString(userInfo.Password)
	if isVerfy := utils.PasswordVerify(pByte, []byte(newPass)); !isVerfy || err != nil {
		return nil, errors.New("密码错误")
	}
	return &userInfo, nil
}
