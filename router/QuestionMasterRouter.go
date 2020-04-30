package router

import "MyFirstGinProject/controller/QuestionController"

func QuestionMasterRouter(base string) {

	r := Router.Group("/" + base)
	r.GET("/getall", QuestionController.GetQuestionMasterList)

	r.GET("getbyid", QuestionController.GetQuestionMasterById)

}
