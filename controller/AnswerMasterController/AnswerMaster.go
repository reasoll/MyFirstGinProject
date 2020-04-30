package AnswerMasterController

import (
	"MyFirstGinProject/service/AnswerMasterService"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetAnswerListByProjectID(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	amList := AnswerMasterService.GetAnswerListByQuestionID(id)

	c.JSON(200, amList)

}
