package router

import "MyFirstGinProject/controller/AnswerMasterController"

func AnswerMasterRouter(base string) {

	r := Router.Group("/" + base)

	r.GET("/getbyid/:id", AnswerMasterController.GetAnswerListByProjectID)

}
