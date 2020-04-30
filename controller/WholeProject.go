package controller

import (
	"MyFirstGinProject/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

//获取整个
func GetWholeProjectById(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	project := service.GetWholeProject(id)

	c.JSON(200, project)
}
