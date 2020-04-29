package ProjectMasterView

import (
	"MyFirstGinProject/model"
	"MyFirstGinProject/view/ProjectMaster"
)

func GetProjectList() ProjectMaster.ProjectMasterList {
	model.ProjectMaster.Add()
}
