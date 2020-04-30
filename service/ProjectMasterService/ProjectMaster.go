package ProjectMasterService

import "MyFirstGinProject/model"

func GetProjectMaster(id int) model.ProjectMaster {
	return model.ProjectMaster{}.GetProjectMasterById(id)
}
