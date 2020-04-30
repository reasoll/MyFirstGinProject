package ProjectController

import (
	"MyFirstGinProject/service/ProjectMasterService"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetProjectMaster(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	pm := ProjectMasterService.GetProjectMaster(id)

	c.JSON(200, pm)

}

func GetProjectMasterList(c *gin.Context) {

	pmList := ProjectMasterService.GetProjectMasterList()
	c.JSON(200, pmList)
}
