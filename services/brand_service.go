package services

import (
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/validation"
	"shopadmin/models"
)

type BrandService struct {
	BaseService
}

//文件存放文件夹名称
const BrandDirName = "brand"

//品牌新增\编辑请求数据
type BrandSaveBody struct {
	Id            int     `json:"id"`
	Name          string  `json:"name"`
	SimpleDesc    string  `json:"simple_desc"`
	IsShow        int     `json:"is_show"`
	FloorPrice    float64 `json:"floor_price"`
	Logo          string  `json:"logo"`
	BigPic        string  `json:"bigPic"`
	RealLogoUrl   string  `json:"realLogoUrl"`
	RealBigPicUrl string  `json:"realBigPicUrl"`
	SortOrder     int     `json:"sort_Order"`
}

func (brand *BrandSaveBody) Valid(v *validation.Validation) {
	if brand.Name == "" {
		v.SetError("品牌名称", "不能为空")
	}
	if brand.Logo == "" {
		v.SetError("品牌LOGO", "不能为空")
	}
	if brand.BigPic == "" {
		v.SetError("品牌专区大图", "不能为空")
	}
}

//获取所有的品牌 返回id 和 name
func (brand *BrandService) GetAllBrandIdNames() []orm.Params {
	var data []orm.Params
	o := orm.NewOrm()
	_, err := o.QueryTable(new(models.NideshopBrand)).Filter("is_show", 1).Values(&data, "id", "name")
	if err != nil {
		return nil
	}
	return brand.TransMapKeyToLower(data)
}

//根据id获取品牌信息
func (brand *BrandService) GetBrandById(id int) (*models.NideshopBrand, error) {
	var brandInfo models.NideshopBrand
	o := orm.NewOrm()
	err := o.QueryTable(new(models.NideshopBrand)).Filter("id", id).One(&brandInfo)
	if err != nil {
		return nil, err
	}
	return &brandInfo, nil
}

func (brand *BrandService) GetBrandLists(pagination map[string]int, otherQueryData map[string]string) (*PageData, error) {
	page := pagination["page"]
	size := pagination["size"]
	var brandList []*models.NideshopBrand
	var cond = orm.NewCondition()
	if v, ok := otherQueryData["name"]; ok {
		cond = cond.And("name__icontains", v)
	}
	o := orm.NewOrm()
	qs := o.QueryTable(new(models.NideshopBrand)).SetCond(cond)
	_, err := qs.Limit(size, (page-1)*size).All(&brandList)
	if err != nil {
		return nil, err
	}
	allCount, err := brand.GetCounts(qs, cond)
	if err != nil {
		return nil, err
	}
	//查询品牌数量以及评论数量
	for _, v := range brandList {
		v.GoodsCounts = v.GetGoodsCounts()
		v.CommentsCounts = v.GetCommentsCount()
	}
	pagedata := PageData{
		size,
		page,
		allCount,
		brand.GetTotalPage(allCount, size),
		brandList,
	}
	return &pagedata, nil
}

//判断品牌是否存在
func (role *BrandService) IsBrandNameExists(id int, name string) bool {
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

//新增菜单
func (brand *BrandService) CreateBrand(body BrandSaveBody, userId int) (*models.NideshopBrand, error) {
	var brandInfo models.NideshopBrand
	if brand.IsBrandNameExists(body.Id, body.Name) {
		return nil, errors.New("该品牌名称已经存在")
	}
	//将图片正式保存到相应的目录中
	var att AttachmentService
	iconFileInfo, err := att.CopyFileToRealPath(BrandDirName, body.RealLogoUrl, userId)
	if err != nil {
		return nil, err
	}
	bigPicFileInfo, err := att.CopyFileToRealPath(BrandDirName, body.RealBigPicUrl, userId)
	if err != nil {
		return nil, err
	}
	brandInfo = models.NideshopBrand{
		Name:          body.Name,
		FloorPrice:    body.FloorPrice,
		PicUrl:        iconFileInfo.Url,
		ListPicUrl:    bigPicFileInfo.Url,
		AppListPicurl: bigPicFileInfo.Url,
		SimpleDesc:    body.SimpleDesc,
		IsShow:        body.IsShow,
		IsNew:         0,
		NewSortOrder:  10,
		SortOrder:     body.SortOrder,
	}
	o := orm.NewOrm()
	num, err := o.Insert(&brandInfo)
	if err != nil || num == 0 {
		return nil, err
	}
	return &brandInfo, nil

}

////更新品牌
func (brand *BrandService) UpdateBrand(body BrandSaveBody, userId int) (*models.NideshopBrand, error) {
	var brandInfo models.NideshopBrand
	if brand.IsBrandNameExists(body.Id, body.Name) {
		return nil, errors.New("该品牌名称已经存在")
	}
	var att AttachmentService
	var iconUrl = body.Logo
	var bigPicUrl = body.BigPic
	if body.RealBigPicUrl != "" {
		bigPicFileInfo, err := att.CopyFileToRealPath(BrandDirName, body.RealBigPicUrl, userId)
		if err != nil {
			return nil, err
		}
		bigPicUrl = bigPicFileInfo.Url
	}
	if body.RealLogoUrl != "" {
		iconFileInfo, err := att.CopyFileToRealPath(BrandDirName, body.RealLogoUrl, 1)
		if err != nil {
			return nil, err
		}
		iconUrl = iconFileInfo.Url
	}

	brandInfo = models.NideshopBrand{
		Id:            body.Id,
		FloorPrice:    body.FloorPrice,
		PicUrl:        iconUrl,
		ListPicUrl:    bigPicUrl,
		AppListPicurl: bigPicUrl,
		SimpleDesc:    body.SimpleDesc,
		SortOrder:     body.SortOrder,
		IsShow:        body.IsShow,
	}
	o := orm.NewOrm()
	num, err := o.Update(&brandInfo, "floor_price", "pic_url", "list_pic_url", "app_list_pic_url", "simple_desc", "sort_order", "is_show")
	if err != nil || num == 0 {
		return nil, err
	}
	return &brandInfo, nil
}
