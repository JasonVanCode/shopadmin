package services

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"

	"shopadmin/models"
)

type KeywordsService struct {
	BaseService
}

//获取默认关键词
func (*KeywordsService) GetDefalutKeyWorld() *models.NideshopKeywords {
	var keyData models.NideshopKeywords
	o := orm.NewOrm()
	err := o.QueryTable(new(models.NideshopKeywords)).Filter("is_default", 1).One(&keyData)
	fmt.Println(err)
	if err != nil {
		return nil
	}
	return &keyData
}

//获取多条关键词
func (*KeywordsService) GetHotKeyWorldList() []*models.NideshopKeywords {
	var keyData []*models.NideshopKeywords
	o := orm.NewOrm()
	num, err := o.QueryTable(new(models.NideshopKeywords)).Filter("is_hot", 1).Distinct().Limit(10).All(&keyData)
	if err != nil || num == 0 {
		return nil
	}
	return keyData
}

//模糊查询关键字
func (*KeywordsService) FuzzyQueryKeywordList(keyword string) []*models.NideshopKeywords {
	var keyData []*models.NideshopKeywords
	o := orm.NewOrm()
	num, err := o.QueryTable(new(models.NideshopKeywords)).Filter("keyword__icontains", keyword).Distinct().Limit(10).All(&keyData)
	if err != nil || num == 0 {
		return nil
	}
	return keyData
}
