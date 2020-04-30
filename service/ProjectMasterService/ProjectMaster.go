package ProjectMasterService

import "MyFirstGinProject/model"

//get Pm by id
func GetProjectMaster(id int) model.ProjectMaster {
	return model.ProjectMaster{}.GetProjectMasterById(id)
}

//Get PmList
func GetProjectMasterList() []model.ProjectMaster {
	return model.ProjectMaster{}.GetProjectMasterList()
}
