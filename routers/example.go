package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/arifluthfi16/gomvcboilerplate/controller"
)

type ExampleRouter struct {}

func (r *ExampleRouter) Route (route *gin.Engine){
	router := route.Group("/example")
	Controller := controller.ExampleController{}

	router.GET("/", Controller.ExampleHandler)
}

