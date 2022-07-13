# GORM
## 连接数据库
连接表的时候，gorm默认会给查询的表名添加上s，可以在连接数据库的时候，进行单表模式的设置
```GO
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
```

设置字段是。字段名称需要大写，否则不会接受
```go
type Stu struct {
	Id   int
	Name string
}
```

## 表查询、删除等操作
**查询**
```go
	stu := Stu{Id: 3, Name: "Lily"}
	result := db.Create(&stu)
	fmt.Println(result.RowsAffected)
```
