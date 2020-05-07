package controller

import (
	"MyFirstGinProject/service"
	"MyFirstGinProject/view"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
)

//获取整个问卷
func GetWholeProjectById(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	project := service.GetWholeProject(id)

	c.JSON(200, project)
}

//创建整个问卷
func CreateSurvey(c *gin.Context) {
	contentLength := c.Request.ContentLength
	var buf = make([]byte, contentLength)
	n, _ := c.Request.Body.Read(buf)

	//surveyJson := string(buf[0:n])

	var survey view.WholePriject
	json.Unmarshal(buf[0:n], &survey)

	flg := service.CreateSurvey(survey)
	if flg {
		c.JSON(200, gin.H{
			"result": "创建成功",
		})
	} else {
		c.JSON(200, gin.H{
			"result": "创建失败",
		})
	}

}
