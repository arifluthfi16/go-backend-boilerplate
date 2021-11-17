package routers

import (
	"github.com/arifluthfi16/gomvcboilerplate/controller"
	"github.com/gin-gonic/gin"
)

type ExampleRouter struct {}

func (r *ExampleRouter) Route (route *gin.Engine){
	router := route.Group("/example")
	Controller := controller.ExampleController{}

	router.GET("/", Controller.ExampleHandler)
	router.GET("/checkdb", Controller.CheckDBConnection)
}

