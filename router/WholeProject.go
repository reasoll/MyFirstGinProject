package router

import "MyFirstGinProject/controller"

func WholeProjectRouter(base string) {

	r := Router.Group("/" + base)

	r.GET("/getbyid/:id", controller.GetWholeProjectById)

}
