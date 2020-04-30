package QuestionController

import (
	"MyFirstGinProject/service/QuestionMasterService"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetQuestionMasterList(c *gin.Context) {

	qmList := QuestionMasterService.GetQuestionMasterList()

	c.JSON(200, qmList)
}

func GetQuestionMasterById(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	qmList := QuestionMasterService.GetQuestionMasterById(id)

	c.JSON(200, qmList)

}
