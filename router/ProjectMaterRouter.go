package router

import (
	"github.com/gin-gonic/gin"
)

func ProjectMasterRouter(base string) {
	r := Router.Group("/" + base)

	r.GET("/get", func(context *gin.Context) {

	})
}
