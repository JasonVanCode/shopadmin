package models

import (
	"github.com/beego/beego/v2/client/orm"
)

//商品评论图片
type NideshopCommentPicture struct {
	Id int `orm:"column(id)" json:"id"`
	//CommentId int    `orm:"column(comment_id)" json:"comment_id"`
	PicUrl    string           `orm:"column(pic_url)" json:"pic_url"`
	SortOrder int              `orm:"column(sort_order)" json:"sort_order"`
	Comment   *NideshopComment `orm:"rel(fk)"`
}

func (*NideshopCommentPicture) TableName() string {
	return "nideshop_comment_picture"
}

func init() {
	orm.RegisterModel(new(NideshopCommentPicture))
}
