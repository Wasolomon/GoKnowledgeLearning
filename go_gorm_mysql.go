package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main()  {
	db, err := gorm.Open("mysql", "root:147258@go?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("连接MySQL数据库失败")
	}
	defer db.Close()
}
