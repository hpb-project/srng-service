package db

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dbhost := beego.AppConfig.String("mysql::dbhost")
	dbuser := beego.AppConfig.String("mysql::dbuser")
	dbpassword := beego.AppConfig.String("mysql::dbpassword")
	dbname := beego.AppConfig.String("mysql::dbname")

	orm.RegisterDriver("mysql", orm.DRMySQL)
	maxIdle := beego.AppConfig.DefaultInt("maxIdle", 50)  //连接池空闲
	maxConn := beego.AppConfig.DefaultInt("maxConn", 512) //连接池最大连接数  数据库默认链接数一般为512
	orm.RegisterDataBase("default", "mysql",
		dbuser+":"+dbpassword+"@tcp("+dbhost+")/"+dbname+"?charset=utf8&parseTime=True&loc=Asia%2FShanghai",
		maxIdle, maxConn)
	// 不同步数据库
	if "dev" == beego.AppConfig.String("runmode") {
		orm.Debug = true
	}
	beego.Info("orm init success")
	tableinit()
	orm.RunSyncdb("default", false, true)
}

func tableinit() {
	orm.RegisterModel(new(TbCommitHash))
	orm.RegisterModel(new(TbRandomConsumed))
	orm.RegisterModel(new(TbReveal))
	orm.RegisterModel(new(TbSubscribe))
	orm.RegisterModel(new(TbUnSubscribe))
}
