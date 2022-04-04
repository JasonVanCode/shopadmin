package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type NideshopTopic struct {
	Id              int     `orm:"column(id)" json:"id"`
	Title           string  `orm:"column(title)" json:"title"`
	Content         string  `orm:"column(content)" json:"content"`
	Avatar          string  `orm:"column(avatar)" json:"avatar"`
	ItemPicUrl      string  `orm:"column(item_pic_url)" json:"item_pic_url"`
	Subtitle        string  `orm:"column(subtitle)" json:"subtitle"`
	TopicCategoryId int     `orm:"column(topic_category_id)" json:"topic_category_id"`
	PriceInfo       float64 `orm:"column(price_info);digits(10);decimals(2)" json:"price_info"`
	ReadCount       string  `orm:"column(read_count)" json:"read_count"`
	ScenePicUrl     string  `orm:"column(scene_pic_url)" json:"scene_pic_url"`
	TopicTemplateId int     `orm:"column(topic_template_id)" json:"topic_template_id"`
	TopicTagId      int     `orm:"column(topic_tag_id)" json:"topic_tag_id"`
	SortOrder       int     `orm:"column(sort_order)" json:"sort_order"`
	IsShow          int     `orm:"column(is_show)" json:"is_show"`
}

func (*NideshopTopic) TableName() string {
	return "nideshop_topic"
}

func init() {
	orm.RegisterModel(new(NideshopTopic))
}
