package services

import (
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"shopadmin/models"
	"strings"
	"time"
)

type AdminRoleService struct {
	BaseService
}

//角色请求体
type RoleRequestBody struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

type RoleStatusBody struct {
	Id     int `json:"id"`
	Status int `json:"status"`
}

type RoleDelBody struct {
	Id string `json:"id"`
}

//分配菜单请求数据
type RoleMenusBody struct {
	Ids    string `json:"ids"`
	RoleId int    `json:"roleId"`
}

//根据id获取角色
func (*AdminRoleService) GetRoleById(id int) (*models.AdminRole, error) {
	var roleInfo models.AdminRole
	o := orm.NewOrm()
	err := o.QueryTable(new(models.AdminRole)).Filter("id", id).One(&roleInfo)
	if err != nil {
		return nil, err
	}
	return &roleInfo, nil
}

//获取角色列表 id name值
func (role *AdminRoleService) GetRoleIdNames() []orm.Params {
	var roleLists []orm.Params
	o := orm.NewOrm()
	_, err := o.QueryTable(new(models.AdminRole)).Filter("status", 1).Values(&roleLists, "id", "name")
	if err != nil {
		return nil
	}
	return role.TransMapKeyToLower(roleLists)
}

//获取角色列表
func (role *AdminRoleService) GetRoleLists(pagination map[string]int, otherQueryData map[string]string) (*PageData, error) {
	page := pagination["page"]
	size := pagination["size"]
	role.GetRoleIdNames()
	var userList []*models.AdminRole
	var cond = orm.NewCondition()
	if v, ok := otherQueryData["name"]; ok {
		cond = cond.And("name__icontains", v)
	}
	o := orm.NewOrm()
	qs := o.QueryTable(new(models.AdminRole)).SetCond(cond)
	_, err := qs.Limit(size, (page-1)*size).All(&userList)
	if err != nil {
		return nil, err
	}
	allCount, err := role.GetRoleCounts(qs, cond)
	if err != nil {
		return nil, err
	}
	pagedata := PageData{
		size,
		page,
		allCount,
		role.GetTotalPage(allCount, size),
		userList,
	}
	return &pagedata, nil
}

func (role *AdminRoleService) GetRoleCounts(qs orm.QuerySeter, cond *orm.Condition) (int, error) {
	num, err := qs.SetCond(cond).Count()
	if err != nil {
		return 0, err
	}
	return int(num), nil
}

//判断角色是否存在
func (role *AdminRoleService) IsRoleNameExists(id int, name string) bool {
	var roleInfo models.AdminRole
	qs := orm.NewOrm().QueryTable(new(models.AdminRole)).Filter("name", name)
	if id > 0 {
		qs.Exclude("id", id)
	}
	qs.One(&roleInfo)
	if roleInfo.Id > 0 {
		return true
	}
	return false
}

//新建角色
func (role *AdminRoleService) AddRole(body RoleRequestBody) (*models.AdminRole, error) {
	var roleInfo models.AdminRole
	//判断账号是否存在
	if role.IsRoleNameExists(0, body.Name) {
		return nil, errors.New("该角色名称已经存在")
	}
	roleInfo = models.AdminRole{
		Name:        body.Name,
		Description: body.Description,
		Status:      body.Status,
		CreatedTime: time.Now(),
		UpdateTime:  time.Now(),
	}
	o := orm.NewOrm()
	num, err := o.Insert(&roleInfo)
	if err != nil || num == 0 {
		return nil, err
	}
	return &roleInfo, nil
}

//更新角色
func (role *AdminRoleService) UpdateRole(body RoleRequestBody) (*models.AdminRole, error) {
	var roleInfo models.AdminRole
	if role.IsRoleNameExists(body.Id, body.Name) {
		return nil, errors.New("该角色名称已经存在")
	}
	roleInfo = models.AdminRole{
		Id:          body.Id,
		Name:        body.Name,
		Description: body.Description,
		UpdateTime:  time.Now(),
	}
	o := orm.NewOrm()
	num, err := o.Update(&roleInfo, "name", "description", "update_time")
	if err != nil || num == 0 {
		return nil, err
	}
	return &roleInfo, nil

}

//修改角色状态
func (role *AdminRoleService) ChangeStatus(body RoleStatusBody) (bool, error) {
	roleInfo, err := role.GetRoleById(body.Id)
	if err != nil || roleInfo.Id != body.Id {
		return false, errors.New("角色不存在")
	}
	o := orm.NewOrm()
	updateData := models.AdminRole{
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

//删除角色
func (*AdminRoleService) DelRole(id string) (bool, error) {
	var ids = strings.Split(id, ",")
	o := orm.NewOrm()
	num, err := o.QueryTable(new(models.AdminRole)).Filter("id__in", ids).Delete()
	if err != nil || num == 0 {
		return false, errors.New("删除失败")
	}
	return true, nil
}

//保存菜单
func (role *AdminRoleService) SaveMenus(body RoleMenusBody) (bool, error) {
	//获取角色信息
	roleInfo, err := role.GetRoleById(body.RoleId)
	if err != nil {
		return false, err
	}
	var roleSave models.AdminRole
	roleSave.Id = roleInfo.Id
	roleSave.Url = body.Ids
	_, err = orm.NewOrm().Update(&roleSave, "url")
	if err != nil {
		return false, err
	}
	return true, nil
}
