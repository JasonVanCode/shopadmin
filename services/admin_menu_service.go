package services

import (
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/validation"
	"shopadmin/models"
	"time"
)

type AdminMenuService struct {
	BaseService
}

//菜单编辑保存的请求数据
type MenuSaveBody struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Status    int    `json:"status"`
	Sort      int    `json:"sort"`
	Parent_id int    `json:"parent_Id"`
	Icon      string `json:"icon"`
}

func (body *MenuSaveBody) Valid(v *validation.Validation) {
	if body.Name == "" {
		v.SetError("菜单名", "不能为空")
	}
	if body.Parent_id == 0 && body.Icon == "" {
		v.SetError("图标", "不能为空")
	}
}

//删除菜单的请求数据
type MenuDelBody struct {
	Id       int `json:"id"`
	ParentId int `json:"parent_id"`
}

//菜单首页数据
func (menu *AdminMenuService) GetMenuLists(pagination map[string]int, otherQueryData map[string]string) (*PageData, error) {
	page := pagination["page"]
	size := pagination["size"]
	var menuList []*models.AdminMenu
	var cond = orm.NewCondition()
	if v, ok := otherQueryData["name"]; ok {
		cond = cond.And("name__icontains", v)
	}
	if v, ok := otherQueryData["parent_id"]; ok {
		cond = cond.And("parent_id", v)
	}
	o := orm.NewOrm()
	qs := o.QueryTable(new(models.AdminMenu)).SetCond(cond)
	_, err := qs.Limit(size, (page-1)*size).All(&menuList)
	if err != nil {
		return nil, err
	}
	allCount, err := menu.GetCounts(qs, cond)
	if err != nil {
		return nil, err
	}
	pagedata := PageData{
		size,
		page,
		allCount,
		menu.GetTotalPage(allCount, size),
		menuList,
	}
	return &pagedata, nil
}

//获取顶级菜单数据
func (menu *AdminMenuService) GetMenuIdNames() []orm.Params {
	var menuList []orm.Params
	o := orm.NewOrm()
	_, err := o.QueryTable(new(models.AdminMenu)).Filter("parent_id", 0).Values(&menuList, "id", "name")
	if err != nil {
		return nil
	}
	return menu.TransMapKeyToLower(menuList)
}

//菜单名是否存在
func (*AdminMenuService) IsMenuNameExists(id int, name string) bool {
	var menuInfo models.AdminMenu
	qs := orm.NewOrm().QueryTable(new(models.AdminMenu)).Filter("name", name)
	if id > 0 {
		qs.Exclude("id", id)
	}
	qs.One(&menuInfo)
	if menuInfo.Id > 0 {
		return true
	}
	return false
}

//新建菜单
func (menu *AdminMenuService) AddMenu(body MenuSaveBody) (*models.AdminMenu, error) {
	var menuInfo models.AdminMenu
	//判断账号是否存在
	if menu.IsMenuNameExists(0, body.Name) {
		return nil, errors.New("该菜单名称已经存在")
	}
	var level int
	if body.Parent_id == 0 {
		level = 1
	} else {
		level = 2
	}
	menuInfo = models.AdminMenu{
		Name:        body.Name,
		ParentId:    body.Parent_id,
		SortId:      body.Sort,
		Icon:        body.Icon,
		Level:       level,
		Status:      body.Status,
		CreatedTime: time.Now(),
		UpdateTime:  time.Now(),
	}
	o := orm.NewOrm()
	num, err := o.Insert(&menuInfo)
	if err != nil || num == 0 {
		return nil, err
	}
	return &menuInfo, nil
}

//编辑菜单
func (menu *AdminMenuService) UpdateMenu(body MenuSaveBody) (*models.AdminMenu, error) {
	var menuInfo models.AdminMenu
	//判断账号是否存在
	if menu.IsMenuNameExists(body.Id, body.Name) {
		return nil, errors.New("该菜单名称已经存在")
	}
	var level int
	if body.Parent_id == 0 {
		level = 1
	} else {
		level = 2
	}
	menuInfo = models.AdminMenu{
		Id:         body.Id,
		Name:       body.Name,
		ParentId:   body.Parent_id,
		SortId:     body.Sort,
		Icon:       body.Icon,
		Level:      level,
		Status:     body.Status,
		UpdateTime: time.Now(),
	}
	o := orm.NewOrm()
	num, err := o.Update(&menuInfo, "name", "parent_id", "sort_id", "icon", "level", "status", "update_time")
	if err != nil || num == 0 {
		return nil, err
	}
	return &menuInfo, nil
}

//获取子菜单数据
func (*AdminMenuService) GetMenuListsByParentId(pid int) ([]models.AdminMenu, int) {
	var menuLists []models.AdminMenu
	o := orm.NewOrm()
	num, err := o.QueryTable(new(models.AdminMenu)).Filter("parent_id", pid).All(&menuLists)
	if err != nil {
		return nil, 0
	}
	return menuLists, int(num)
}

//删除菜单
func (menu *AdminMenuService) DelMenu(body MenuDelBody) (bool, error) {
	//判断如果是顶级菜单，如果下面有子菜单就不允许删除
	if body.ParentId == 0 {
		_, num := menu.GetMenuListsByParentId(body.Id)
		if num > 0 {
			return false, errors.New("存在子菜单，不允许删除")
		}
	}
	o := orm.NewOrm()
	num, err := o.QueryTable(new(models.AdminMenu)).Filter("id", body.Id).Delete()
	if err != nil || num == 0 {
		return false, errors.New("删除失败")
	}
	return true, nil
}

//获取所有能用显示的菜单
func (menu *AdminMenuService) GetAllShowMenulis() []models.AdminMenu {
	var menuLists []models.AdminMenu
	o := orm.NewOrm()
	_, err := o.QueryTable(new(models.AdminMenu)).Filter("status", 1).OrderBy("sort_id").All(&menuLists)
	if err != nil {
		return nil
	}
	return menuLists
}

//树形数据格式
func (menu *AdminMenuService) GetTreeMenuLists(pid int, lists []models.AdminMenu) []models.AdminMenu {
	var data []models.AdminMenu
	for _, v := range lists {
		if v.ParentId == pid {
			children := menu.GetTreeMenuLists(v.Id, lists)
			if len(children) > 0 {
				v.Children = children
			}
			data = append(data, v)
		}
	}
	return data
}
