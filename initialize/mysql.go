package initialize

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

//初始化连接数据库
func init() {
	host, _ := beego.AppConfig.String("mysql::host")
	name, _ := beego.AppConfig.String("mysql::name")
	password, _ := beego.AppConfig.String("mysql::password")
	port, _ := beego.AppConfig.Int("mysql::port")
	database, _ := beego.AppConfig.String("mysql::database")
	charset, _ := beego.AppConfig.String("mysql::charset")
	connectInfo := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&loc=Local", name, password, host, port, database, charset)
	fmt.Println("mysql 数据库连接的信息是：", connectInfo)
	orm.MaxIdleConnections(20)
	err := orm.RegisterDataBase("default", "mysql", connectInfo)
	if err != nil {
		logs.Error(err)
	}

}
