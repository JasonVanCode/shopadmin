package models

import (
	"errors"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Nideshop_user struct {
	Id            int
	Username      string    `orm:"column(username);size(128)"`
	Password      string    `orm:"column(password);size(128)"`
	Gender        int       `orm:"column(gender);size(128)"`
	Birthday      int       `orm:"column(birthday);size(128)"`
	RegisterTime  time.Time `orm:"column(register_time);size(128)"`
	LastLoginTime time.Time `orm:"column(last_login_time);size(128)"`
	LastLoginIp   string    `orm:"column(last_login_ip);size(128)"`
	UserLevelId   int       `orm:"column(user_level_id);size(128)"`
	Nickname      string    `orm:"column(nickname);size(128)"`
	Mobile        string    `orm:"column(mobile);size(128)"`
	RegisterIp    string    `orm:"column(register_ip);size(128)"`
	Avatar        string    `orm:"column(avatar);size(128)"`
	WeixinOpenid  string    `orm:"column(weixin_openid);size(128)"`
}

//自定义表名
func (u *Nideshop_user) TableName() string {
	return "nideshop_user"
}

func init() {
	orm.RegisterModel(new(Nideshop_user))
}

// AddNideshop_user insert a new Nideshop_user into database and returns
// last inserted Id on success.
func AddNideshop_user(m *Nideshop_user) (id int, err error) {
	//o := orm.NewOrm()
	//id, err = o.Insert(m)
	return
}

// GetNideshop_userById retrieves Nideshop_user by Id. Returns error if
// Id doesn't exist
func GetNideshop_userById(id int) (v *Nideshop_user, err error) {
	o := orm.NewOrm()
	v = &Nideshop_user{Id: id}
	if err = o.QueryTable(new(Nideshop_user)).Filter("Id", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllNideshop_user retrieves all Nideshop_user matches certain condition. Returns empty list if
// no records exist
func GetAllNideshop_user(query map[string]string, fields []string, sortby []string, order []string,
	offset int, limit int) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Nideshop_user))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Nideshop_user
	qs = qs.OrderBy(sortFields...).RelatedSel()
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateNideshop_user updates Nideshop_user by Id and returns error if
// the record to be updated doesn't exist
func UpdateNideshop_userById(m *Nideshop_user) (err error) {
	//o := orm.NewOrm()
	//v := Nideshop_user{Id: m.Id}
	//// ascertain id exists in the database
	//if err = o.Read(&v); err == nil {
	//	var num int
	//	if num, err = o.Update(m); err == nil {
	//		fmt.Println("Number of records updated in database:", num)
	//	}
	//}
	return
}

// DeleteNideshop_user deletes Nideshop_user by Id and returns error if
// the record to be deleted doesn't exist
func DeleteNideshop_user(id int) (err error) {
	//o := orm.NewOrm()
	//v := Nideshop_user{Id: id}
	//// ascertain id exists in the database
	//if err = o.Read(&v); err == nil {
	//	var num int
	//	if num, err = o.Delete(&Nideshop_user{Id: id}); err == nil {
	//		fmt.Println("Number of records deleted in database:", num)
	//	}
	//}
	return
}
