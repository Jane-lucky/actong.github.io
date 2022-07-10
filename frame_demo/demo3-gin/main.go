package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)
type Todo struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}
func main() {
	//创建数据库
	//sql:CREATE DATABASE dubble;
	//连接数据库

	//创建默认的路由引擎
	r := gin.Default()
	//告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("/static", "./frame_demo/demo3-gin/static")
	//注册路由
	//告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("./frame_demo/demo3-gin/templates/*")
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	//v1
	v1Group := r.Group("v1")
	{
		v1Group.POST("/todo",func(ctx *gin.Context) {
			//添加
		})
		//查看所有的代办事项
		v1Group.GET("/todo",func(ctx *gin.Context) {

		})
		// 查看单独的代办事项
		v1Group.GET("/todo/:id",func(ctx *gin.Context) {

		})
		//修改某一项代办事项
		v1Group.PUT("/todo/:id",func(ctx *gin.Context) {

		})

		//删除某一代办事项
		v1Group.DELETE("/todo/:id",func(ctx *gin.Context) {

		})
	}

	//启动服务
	r.Run()
}
