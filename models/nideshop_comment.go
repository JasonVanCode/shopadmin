package models

import (
	"github.com/beego/beego/v2/client/orm"
	"time"
)

//商品评论
type NideshopComment struct {
	Id      int       `orm:"column(id)" json:"id"`
	TypeId  int       `orm:"column(type_id)" json:"type_id"`
	ValueId int       `orm:"column(value_id)" json:"value_id"` //存的是goods_id
	Content string    `orm:"column(content)" json:"content"`
	AddTime time.Time `orm:"column(add_time)" json:"add_time"`
	Status  int       `orm:"column(status)" json:"status"`
	//UserId     int            `orm:"column(user_id)" json:"user_id"`
	User            *Nideshop_user            `orm:"rel(fk)"`
	NewContent      string                    `orm:"column(new_content)" json:"new_content"`
	CommentPictures []*NideshopCommentPicture `orm:"reverse(many)"`
}

func (*NideshopComment) TableName() string {
	return "nideshop_comment"
}

func init() {
	orm.RegisterModel(new(NideshopComment))
}
