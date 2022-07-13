package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Stu struct {
	Id   int
	Name string
}

func main() {
	dsn := "root:123456@tcp(124.70.56.141:32250)/test?charset=utf8mb4"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println("数据库连接失败", err.Error())
	}
	//创建数据
	// stu := Stu{Id: 3, Name: "Lily"}
	// result := db.Create(&stu)
	// fmt.Println(result.RowsAffected)

	// var stu Stu
	//检索单个对象
	// result := db.First(&stu)
	//检索所有对象
	var stus []Stu
	result := db.Find(&stus)
	fmt.Println(result.RowsAffected, result.Error)
}
