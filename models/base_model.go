package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	databaseResources "github.com/doctorJE/users-go/classes/database-resources"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dbDriver := beego.AppConfig.String("DB_Driver")
	dbHost := beego.AppConfig.String("DB_HOST")
	dbPort := beego.AppConfig.String("DB_PORT")
	dbDatabase := beego.AppConfig.String("DB_DATABASE")
	dbUsername := beego.AppConfig.String("DB_USERNAME")
	dbPassword := beego.AppConfig.String("DB_PASSWORD")
	dsn := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbDatabase + "?charset=utf8&parseTime=true&loc=Asia%2FTaipei"

	//設定預設資料庫
	err := orm.RegisterDataBase(
		"default",
		dbDriver,
		dsn,
	)
	if err != nil {
		fmt.Println(err.Error())
	}

	//註冊定義的 model
	orm.RegisterModel(new(databaseResources.User))

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}
