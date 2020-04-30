package router

import "MyFirstGinProject/controller/ProjectController"

func ProjectMasterRouter(base string) {
	r := Router.Group("/" + base)

	r.GET("/get/:id", ProjectController.GetProjectMaster)
}
