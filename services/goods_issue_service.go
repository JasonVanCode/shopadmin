package services

import (
	"github.com/beego/beego/v2/client/orm"
	"shopadmin/models"
)

type GoodsIssueService struct {
	BaseService
}

func (*GoodsIssueService) GetIssueList(goodsId int) []*models.NideshopGoodsIssue {
	var issueList []*models.NideshopGoodsIssue
	o := orm.NewOrm()
	cond := orm.NewCondition()
	if goodsId > 0 {
		cond = cond.And("goods_id__icontains", goodsId)
	}
	_, err := o.QueryTable(new(models.NideshopGoodsIssue)).SetCond(cond).All(&issueList)
	if err != nil {
		return nil
	}
	return issueList
}
