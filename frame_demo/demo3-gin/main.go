package main

import "github.com/gin-gonic/gin"

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello",
	})
}

func main() {
	//创建默认的路由引擎
	r := gin.Default()
	r.GET("/hello", sayHello)

	//启动服务
	r.Run(":9090")
}
