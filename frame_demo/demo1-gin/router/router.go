package router

import (
	"golang-learning/package_demo/demo1-gin/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Load(g *gin.Engine) *gin.Engine {
	g.Use(gin.Recovery())
	// 404
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 not found")
	})

	g.GET("/", controller.Index)

	return g
}
