package services

import (
	"github.com/beego/beego/v2/client/orm"
	"shopadmin/models"
)

type TopicService struct {
	BaseService
}

func (t *TopicService) GetPageList(page, size int) (*PageData, error) {
	var topic = new(models.NideshopTopic)
	var data []orm.Params
	o := orm.NewOrm()
	_, err := o.QueryTable(topic).Limit(size, (page-1)*size).Values(&data, "id", "scene_pic_url", "title", "subtitle", "price_info")
	if err != nil {
		return nil, err
	}
	//处理map数据
	t.updateJsonKey(data)

	count, err := o.QueryTable(topic).Count()
	if err != nil {
		return nil, err
	}

	allCount := int(count)

	return &PageData{
		size,
		page,
		allCount,
		(allCount + size - 1) / size,
		data,
	}, nil
}
