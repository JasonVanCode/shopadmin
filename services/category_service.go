package services

import (
	"errors"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/validation"
	"shopadmin/models"
	"shopadmin/utils"
)

type CategoryService struct {
	BaseService
	CategoryRtnJson
}

//文件存放文件夹名称
const CategoryDirName = "category"

//---------小程序短业务处理----------------

//商品分类数据返回
type CategoryRtnJson struct {
	CurCategory     *models.NideshopCategory   `json:"currentCategory"`
	ParentCategory  *models.NideshopCategory   `json:"parentCategory"`
	BrotherCategory []*models.NideshopCategory `json:"brotherCategory"`
}

//根据id获取分类信息
func (c *CategoryService) GetCategoryById(id int) *models.NideshopCategory {
	var category models.NideshopCategory
	o := orm.NewOrm()
	err := o.QueryTable(new(models.NideshopCategory)).Filter("id", id).One(&category)
	if err != nil {
		return nil
	}
	return &category
}

//根据parent_id 获取分类信息
func (c *CategoryService) GetCategoryByParentId(pid int) []*models.NideshopCategory {
	var categoryList []*models.NideshopCategory
	o := orm.NewOrm()
	_, err := o.QueryTable(new(models.NideshopCategory)).Filter("parent_id", pid).All(&categoryList)
	if err != nil {
		return nil
	}
	return categoryList
}

//
func (*CategoryService) GetParentCategoryIds(category_ids []int) []int {
	var parentIds []orm.Params
	o := orm.NewOrm()
	_, err := o.QueryTable(new(models.NideshopCategory)).Filter("id__in", category_ids).Values(&parentIds, "parent_id")
	if err != nil {
		return nil
	}
	return utils.TransMapValueToSliceIntWithKey(parentIds, "ParentId")
}

func (*CategoryService) GetCategoryByIds(category_ids []int) []orm.Params {
	var categoryList []orm.Params
	o := orm.NewOrm()
	_, err := o.QueryTable(new(models.NideshopCategory)).Filter("id__in", category_ids).Values(&categoryList, "id", "name")
	if err != nil {
		return nil
	}
	return categoryList
}

func (*CategoryService) GetCategoryIds(list []*models.NideshopCategory) (ids []int) {
	if len(list) == 0 {
		return nil
	}
	for _, v := range list {
		ids = append(ids, v.Id)
	}
	return
}

//---------后台管理系统业务处理----------------

type CategoryRequestBody struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	FrontName        string `json:"front_name"`
	FrontDesc        string `json:"front_desc"`
	IsShow           int    `json:"is_show"`
	SortOrder        int    `json:"sort_order"`
	ParentId         int    `json:"parent_id"`
	WapBannerUrl     string `json:"wap_banner_url"`
	RealWapBannerUrl string `json:"real_wap_banner_url"`
	IconUrl          string `json:"icon_url"`
	RealIconUrl      string `json:"real_icon_url"`
}

func (c *CategoryRequestBody) Valid(v *validation.Validation) {
	if c.Name == "" {
		v.SetError("分类名称", "不能为空")
	}
	if c.IconUrl == "" {
		v.SetError("LOGO", "不能为空")

	}
	if c.WapBannerUrl == "" {
		v.SetError("分类专区图", "不能为空")
	}
}

//获取分类后台管理系统数据
func (category *CategoryService) GetAdminCategoryLists(pagination map[string]int, otherQueryData map[string]string) (*PageData, error) {
	page := pagination["page"]
	size := pagination["size"]
	var categoryList []*models.NideshopCategory
	var cond = orm.NewCondition()
	cond = cond.And("parent_id", otherQueryData["pid"])
	if v, ok := otherQueryData["name"]; ok {
		cond = cond.And("name__icontains", v)
	}
	o := orm.NewOrm()
	qs := o.QueryTable(new(models.NideshopCategory)).SetCond(cond)
	_, err := qs.Limit(size, (page-1)*size).All(&categoryList)
	if err != nil {
		return nil, err
	}
	allCount, err := category.GetCounts(qs, cond)
	if err != nil {
		return nil, err
	}
	pagedata := PageData{
		size,
		page,
		allCount,
		category.GetTotalPage(allCount, size),
		categoryList,
	}
	return &pagedata, nil
}

func (category *CategoryService) GetAttributeByCateId(cateId int, name string) (interface{}, error) {
	cateInfo := category.GetCategoryById(cateId)
	if cateInfo.Id != cateId {
		return nil, errors.New("查不到该分类数据")
	}
	data, err := cateInfo.GetAttributeLists(name)
	if err != nil {
		return nil, errors.New("查不到该分类数据")
	}
	return struct {
		CategoryInfo   *models.NideshopCategory   `json:"category_info"`
		AttributeLists []models.NideshopAttribute `json:"attribute_lists"`
	}{
		cateInfo,
		data,
	}, nil
}

//获取顶级分类的数据
func (category *CategoryService) GetTopCategoryIdName() []map[string]interface{} {
	var lists []*models.NideshopCategory
	lists = category.GetCategoryByParentId(0)
	var data []map[string]interface{}
	if len(lists) > 0 {
		for _, v := range lists {
			mapData := map[string]interface{}{
				"id":   v.Id,
				"name": v.Name,
			}
			data = append(data, mapData)
		}
	}
	return data
}

//获取所有2级分类的数据
func (category *CategoryService) GetAllCategoryIdNames() []orm.Params {
	var lists []orm.Params
	o := orm.NewOrm()
	_, err := o.QueryTable(new(models.NideshopCategory)).Filter("parent_id__gt", 0).Values(&lists, "id", "name")
	if err != nil {
		return nil
	}
	return category.TransMapKeyToLower(lists)
}

//判断分类名字是否存在
func (*CategoryService) IsCategoryNameExists(id int, name string) bool {
	var cInfo models.NideshopCategory
	o := orm.NewOrm()
	var seter orm.QuerySeter
	seter = o.QueryTable(new(models.NideshopCategory)).Filter("name", name)
	if id > 0 {
		seter = seter.Exclude("id", id)
	}
	seter.One(&cInfo, "id")
	if cInfo.Id > 0 {
		return true
	}
	return false
}

//新建分类
func (category *CategoryService) CreateCategory(body CategoryRequestBody, userId int) (*models.NideshopCategory, error) {
	if category.IsCategoryNameExists(0, body.Name) {
		return nil, errors.New("该分类名存在")
	}
	var attach AttachmentService
	iconInfo, err := attach.CopyFileToRealPath(CategoryDirName, body.RealIconUrl, userId)
	if err != nil {
		return nil, err
	}
	wapInfo, err := attach.CopyFileToRealPath(CategoryDirName, body.RealWapBannerUrl, userId)
	if err != nil {
		return nil, err
	}
	var categoryInfo = models.NideshopCategory{
		Name:         body.Name,
		ParentId:     body.ParentId,
		FrontDesc:    body.FrontDesc,
		FrontName:    body.FrontName,
		IsShow:       body.IsShow,
		SortOrder:    body.SortOrder,
		IconUrl:      iconInfo.Url,
		WapBannerUrl: wapInfo.Url,
	}
	o := orm.NewOrm()
	_, err = o.Insert(&categoryInfo)
	if err != nil {
		return nil, err
	}
	return &categoryInfo, nil
}

//编辑分类
func (category *CategoryService) UpdateCategory(body CategoryRequestBody, userId int) (*models.NideshopCategory, error) {
	if category.IsCategoryNameExists(body.Id, body.Name) {
		return nil, errors.New("该分类名存在")
	}
	var iconUrl = body.IconUrl
	var wapUrl = body.WapBannerUrl
	var attach AttachmentService
	if body.RealIconUrl != "" {
		iconInfo, err := attach.CopyFileToRealPath(CategoryDirName, body.RealIconUrl, userId)
		if err != nil {
			return nil, err
		}
		iconUrl = iconInfo.Url
	}
	if body.RealWapBannerUrl != "" {
		wapInfo, err := attach.CopyFileToRealPath(CategoryDirName, body.RealWapBannerUrl, userId)
		if err != nil {
			return nil, err
		}
		wapUrl = wapInfo.Url
	}
	var categoryInfo = models.NideshopCategory{
		Id:           body.Id,
		Name:         body.Name,
		ParentId:     body.ParentId,
		FrontDesc:    body.FrontDesc,
		FrontName:    body.FrontName,
		IsShow:       body.IsShow,
		SortOrder:    body.SortOrder,
		IconUrl:      iconUrl,
		WapBannerUrl: wapUrl,
	}
	o := orm.NewOrm()
	_, err := o.Update(&categoryInfo, "id", "name", "parent_id", "front_desc", "front_name", "is_show", "sort_order", "icon_url", "wap_banner_url")
	if err != nil {
		return nil, err
	}
	return &categoryInfo, nil
}

//删除分类
func (category *CategoryService) DelCategory(id int) (bool, error) {
	var cInfo = models.NideshopCategory{
		Id: id,
	}
	o := orm.NewOrm()
	//删除分类
	to, err := o.Begin()
	if err != nil {
		return false, err
	}
	num, err := to.Delete(&cInfo)
	if err != nil {
		to.Rollback()
		return false, err
	}
	//删除下面的属性
	if num > 0 {
		_, err = to.QueryTable(new(models.NideshopAttribute)).Filter("attribute_category_id", id).Delete()
		if err != nil {
			to.Rollback()
			return false, err
		}
	}
	to.Commit()
	return true, nil
}

//获取所有分类并查询子类数据
func (category *CategoryService) GetCategoryWithChildren() ([]*models.NideshopCategory, error) {
	//获取所有顶级的分类
	topLists := category.GetCategoryByParentId(0)
	var ids []int
	for _, v := range topLists {
		ids = append(ids, v.Id)
	}
	fmt.Println("ids--->", ids)
	var childLists []*models.NideshopCategory
	o := orm.NewOrm()
	num, err := o.QueryTable(new(models.NideshopCategory)).Filter("parent_id__in", ids).All(&childLists)
	if err != nil {
		return nil, err
	}
	fmt.Println(len(childLists))
	if num > 0 {
		for _, v := range topLists {
			for _, v2 := range childLists {
				if v.Id == v2.ParentId {
					v.Children = append(v.Children, v2)
				}
			}
		}
	}
	return topLists, nil
}
