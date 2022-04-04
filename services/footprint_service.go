package services

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"shopadmin/models"
	"time"
)

type FootPrintService struct {
	BaseService
}

//获取用户浏览商品数据
func (*FootPrintService) GetFootprintList(userId int) []*models.NideshopFootprint {
	var footPrintList []*models.NideshopFootprint
	o := orm.NewOrm()
	_, err := o.QueryTable(new(models.NideshopFootprint)).Filter("user_id", userId).RelatedSel().OrderBy("add_time").All(&footPrintList)
	if err != nil {
		return nil
	}
	return footPrintList
}

//添加足迹
func (*FootPrintService) AddFootprint(userId, goodsId int) {
	insertData := models.NideshopFootprint{
		UserId:  userId,
		AddTime: time.Now(),
		Goods:   &models.NideshopGoods{Id: goodsId},
	}
	o := orm.NewOrm()
	_, err := o.Insert(&insertData)
	if err != nil {
		logs.Error("浏览足迹数据添加失败")
	}
}
