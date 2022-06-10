package services

import (
	"encoding/base64"
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/validation"
	"shopadmin/models"
	"shopadmin/utils"
	"strings"
	"time"
)

type AdminUserService struct {
	BaseService
}

type UserRequestBody struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	Nickname   string `json:"nickname"`
	Password   string `json:"password"`
	RePassword string `json:"rePassword"`
}

func (u *UserRequestBody) Valid(v *validation.Validation) {
	if u.Username == "" {
		v.SetError("用户名", "不能为空")
	}
	if u.Nickname == "" {
		v.SetError("昵称", "不能为空")
	}
	if u.Password == "" {
		v.SetError("密码", "不能为空")
	}
	if u.RePassword == "" {
		v.SetError("确认密码", "不能为空")
	}
}

type UserDelBody struct {
	Id string `json:"id"`
}

type UserEditBody struct {
	Id int `json:"id"`
}

type UserStatusBody struct {
	Id     int `json:"id"`
	Status int `json:"status"`
}

type UserRoleBody struct {
	Id      int    `json:"id"`
	RoleIds string `json:"roleIds"`
}

//根据id获取用户信息
func (*AdminUserService) GetUserById(id int) (*models.AdminUser, error) {
	var user models.AdminUser
	o := orm.NewOrm()
	err := o.QueryTable(new(models.AdminUser)).Filter("id", id).One(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

//判断是否存在重复账号
func (*AdminUserService) IsAcountExists(id int, name string) bool {
	var user models.AdminUser
	qs := orm.NewOrm().QueryTable(new(models.AdminUser)).Filter("username", name)
	if id > 0 {
		qs.Exclude("id", id)
	}
	qs.One(&user)
	if user.Id > 0 {
		return true
	}
	return false
}

func (user *AdminUserService) GetUserLists(pagination map[string]int, otherQueryData map[string]string) (*PageData, error) {
	page := pagination["page"]
	size := pagination["size"]
	var userList []*models.AdminUser
	var cond = orm.NewCondition()
	if v, ok := otherQueryData["name"]; ok {
		cond = cond.And("nickname__icontains", v)
	}
	o := orm.NewOrm()
	qs := o.QueryTable(new(models.AdminUser)).SetCond(cond)
	_, err := qs.Limit(size, (page-1)*size).All(&userList)
	if err != nil {
		return nil, err
	}
	allCount, err := user.GetUserCounts(qs, cond)
	if err != nil {
		return nil, err
	}
	pagedata := PageData{
		size,
		page,
		allCount,
		user.GetTotalPage(allCount, size),
		userList,
	}
	return &pagedata, nil
}

func (a *AdminUserService) GetUserCounts(qs orm.QuerySeter, cond *orm.Condition) (int, error) {
	count, err := qs.SetCond(cond).Count()
	return int(count), err
}

//新建用户
func (user *AdminUserService) AddUser(body UserRequestBody) (*models.AdminUser, error) {
	var userInfo models.AdminUser
	//判断账号是否存在
	if user.IsAcountExists(0, body.Username) {
		return nil, errors.New("该账号已经存在")
	}
	userInfo = models.AdminUser{
		Username:    body.Username,
		Nickname:    body.Nickname,
		Status:      1,
		CreatedTime: time.Now(),
		UpdateTime:  time.Now(),
	}
	password, err := utils.PasswordHash(body.Password)
	if err != nil {
		return nil, err
	}
	userInfo.Password = base64.StdEncoding.EncodeToString([]byte(password))
	o := orm.NewOrm()
	num, err := o.Insert(&userInfo)
	if err != nil || num == 0 {
		return nil, err
	}
	return &userInfo, nil
}

//更新用户
func (user *AdminUserService) UpdateUser(body UserRequestBody) (*models.AdminUser, error) {
	var userInfo models.AdminUser
	if user.IsAcountExists(body.Id, body.Username) {
		return nil, errors.New("该账号已经存在")
	}
	userInfo = models.AdminUser{
		Id:         body.Id,
		Username:   body.Username,
		Nickname:   body.Nickname,
		UpdateTime: time.Now(),
	}
	password, err := utils.PasswordHash(body.Password)
	if err != nil {
		return nil, err
	}
	userInfo.Password = base64.StdEncoding.EncodeToString([]byte(password))
	o := orm.NewOrm()
	num, err := o.Update(&userInfo, "username", "password", "nickname", "update_time")
	if err != nil || num == 0 {
		return nil, err
	}
	return &userInfo, nil

}

//删除用户
func (*AdminUserService) DelUser(id string) (bool, error) {
	var ids = strings.Split(id, ",")
	o := orm.NewOrm()
	num, err := o.QueryTable(new(models.AdminUser)).Filter("id__in", ids).Delete()
	if err != nil || num == 0 {
		return false, errors.New("删除失败")
	}
	return true, nil
}

//修改用户状态
func (user *AdminUserService) ChangeStatus(body UserStatusBody) (bool, error) {
	userInfo, err := user.GetUserById(body.Id)
	if err != nil || userInfo.Id != body.Id {
		return false, errors.New("用户不存在")
	}
	o := orm.NewOrm()
	updateData := models.AdminUser{
		Id:         body.Id,
		Status:     body.Status,
		UpdateTime: time.Now(),
	}
	num, err := o.Update(&updateData, "status", "update_time")
	if err != nil || num == 0 {
		return false, errors.New("修改失败")
	}
	return true, nil
}

//修改用户的角色
func (user *AdminUserService) ChangeUserRole(body UserRoleBody) (bool, error) {
	if body.RoleIds == "" {
		return false, errors.New("请选择要添加的角色")
	}
	userInfo, err := user.GetUserById(body.Id)
	if err != nil || userInfo.Id != body.Id {
		return false, errors.New("用户不存在")
	}
	o := orm.NewOrm()
	updateData := models.AdminUser{
		Id:         body.Id,
		Role:       body.RoleIds,
		UpdateTime: time.Now(),
	}
	num, err := o.Update(&updateData, "role", "update_time")
	if err != nil || num == 0 {
		return false, errors.New("修改失败")
	}
	return true, nil

}
