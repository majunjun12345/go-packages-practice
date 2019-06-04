package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// 在数据库中的表名变成了 users，首字母小写，加上 s，字段也会变成小写
// 必须包含 gorm.Model
type Product struct {
	gorm.Model
	Code  string
	Price string
	Color string
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@/test_gorm?charset=utf8&parseTime=True&loc=Local")
	defer db.Close()
	if err != nil {
		panic("连接数据库失败")
	}

	db.DropTableIfExists("products")

	// 自动迁移仅仅会创建表，缺少列和索引，并且不会改变现有列的类型或删除未使用的列以保护数据
	// 除上述字段外，还会额外创建四个字段：id、created_at、updated_at、deleted_at
	db.AutoMigrate(&Product{})

	// 检查表是否存在
	if !db.HasTable(&Product{}) { // db.HasTable("products")
		fmt.Println("创建表")
		// db.CreateTable(&Product{})

		// 创建表时添加表后缀
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Product{})
	}

	// 添加唯一索引
	db.Model(&Product{}).AddUniqueIndex("idx_gorm_code", "code")

	db.Create(&Product{Code: "M", Price: "15"})
	db.Create(&Product{Code: "L", Price: "19"})

	var product, p Product
	db.First(&product, 1)
	fmt.Println(product)
	db.First(&p, "code = ?", "L")
	fmt.Println(p)

	db.Model(&p).Update("price", "2000")
	fmt.Println(p)

	db.Model(&Product{}).DropColumn("price")
}
