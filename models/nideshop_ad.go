package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

type NideshopAd struct {
	Id           int       `orm:"column(id)" json:"id"`
	AdPositionId int       `orm:"column(ad_position_id)" json:"ad_position_id"`
	MediaType    int       `orm:"column(media_type)" json:"media_type"`
	Name         string    `orm:"column(name)" json:"name"`
	Link         string    `orm:"column(link)" json:"link"`
	ImageUrl     string    `orm:"column(image_url)" json:"image_url"`
	Content      string    `orm:"column(content)" json:"content"`
	EndTime      time.Time `orm:"column(end_time)" json:"end_time"`
	Enabled      int       `orm:"column(enabled)" json:"enabled"`
}

func (*NideshopAd) TableName() string {
	return "nideshop_ad"
}

func init() {
	orm.RegisterModel(new(NideshopAd))
}
