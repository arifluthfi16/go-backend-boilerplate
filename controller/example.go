package controller

import (
	"github.com/arifluthfi16/gomvcboilerplate/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ExampleController struct {}

func (controller ExampleController) ExampleHandler(c *gin.Context) {
	var responseString = services.ExampleService.SayHi()
	var response = SuccessResponse{
		Status: true,
		Msg:    responseString,
	}
	c.AbortWithStatusJSON(http.StatusOK, response)
}


