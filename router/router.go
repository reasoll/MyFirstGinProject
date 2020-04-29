package router

import (
	"MyFirstGinProject/router/middleware"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func NewRouter() *gin.Engine {

	r := gin.Default()
	gin.SetMode(gin.DebugMode)
	r.Use(middleware.Cors())

	r.GET("/ping", func(context *gin.Context) {
		context.Writer.WriteString("Pong")


	})


	r.NoRoute(func(context *gin.Context) {
		context.Writer.WriteString("对不起，页面不存在！！！")
	})

	Router = r
	ProjectMasterRouter("project");
	return Router
}
