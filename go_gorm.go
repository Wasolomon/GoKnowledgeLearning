package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Product struct {
	gorm.Model
	Code string
	Price uint
}

func main()  {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("连接数据库失败")
	}
	defer db.Close()
	//自动迁移模式
	db.AutoMigrate(&Product{})

	db.Create(&Product{Code: "L1212", Price: 1000})

	var product Product
	//查询id为1的product
	db.First(&product, 1)
	db.First(&product, "code = ?", "L1212")

	//更新 - 更新product的price为2000
	db.Model(&product).Update("Price", 2000)

	db.Delete(&product)
}
