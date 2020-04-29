package model

import (
	"MyFirstGinProject/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func Setup(){
	var mySqlConfig = config.Conf.MysqlConfig
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true&loc=Local",
		mySqlConfig.UserNmae,mySqlConfig.Password,mySqlConfig.Addr,mySqlConfig.Port,mySqlConfig.DataBaseName)

	DB,err := gorm.Open("mysql",connStr)

	if err != nil {
		panic(err)
	}

	//不加s建表
	DB.SingularTable(true)

	db = DB


}
