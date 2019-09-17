package utils

import (
	"database/sql"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" //一定要导入，不然报错
)

var db *sql.DB

func InitMysql() {

	driverName := beego.AppConfig.String("driverName")
	orm.RegisterDriver(driverName, orm.DRMySQL)
	user := beego.AppConfig.String("mysqlUser")
	pwd := beego.AppConfig.String(("mysqlPwd"))
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("dbname")
	dbconn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"
	logs.Info(dbconn)
	err := orm.RegisterDataBase("default", driverName, dbconn)
	if err != nil {
		return
	}
	//db1 , err := sql.Open(driverName,dbconn)
	/*if err != nil {
		logs.Error(err.Error())

	}else {
		//db = db1
		logs.Info("conn database succ")
	}*/

}
