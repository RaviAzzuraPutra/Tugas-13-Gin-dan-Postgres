package controller

import (
	"formative-13/helper"
	"formative-13/request"
	service_interface "formative-13/service/interface"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	service service_interface.Service_Interface
}

func NewControllerRegistry(service service_interface.Service_Interface) *Controller {
	return &Controller{
		service: service,
	}
}

func (c *Controller) Create(ctx *gin.Context) {
	request := new(request.Bioskop_Request)

	errRequest := ctx.ShouldBind(request)

	if errRequest != nil {
		ctx.JSON(400, gin.H{
			"Error": "Bad Request :" + errRequest.Error(),
		})
	}

	Bioskop, err := c.service.Create(request)

	if err != nil {
		if appErr, ok := err.(*helper.AppError); ok {
			ctx.JSON(appErr.Code, gin.H{
				"Error": appErr.Message,
			})
			return
		}

		ctx.JSON(500, gin.H{
			"Error": "Internal Server Error",
		})
		return
	}

	ctx.JSON(201, gin.H{
		"Message": "Success Create Cinema",
		"Data":    Bioskop,
	})
}

func (c *Controller) Get(ctx *gin.Context) {

	Bioskop, err := c.service.Get()

	if err != nil {
		if appErr, ok := err.(*helper.AppError); ok {
			ctx.JSON(appErr.Code, gin.H{
				"Error": appErr.Message,
			})
			return
		}

		ctx.JSON(500, gin.H{
			"Error": "Internal Server Error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"Message": "Success Get Cinema",
		"Data":    Bioskop,
	})
}

func (c *Controller) GetById(ctx *gin.Context) {

	id := ctx.Param("id")

	Bioskop, err := c.service.GetById(id)

	if err != nil {
		if appErr, ok := err.(*helper.AppError); ok {
			ctx.JSON(appErr.Code, gin.H{
				"Error": appErr.Message,
			})
			return
		}

		ctx.JSON(500, gin.H{
			"Error": "Internal Server Error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"Message": "Success Get Cinema By Id",
		"Data":    Bioskop,
	})
}

func (c *Controller) Update(ctx *gin.Context) {

	id := ctx.Param("id")

	request := new(request.Bioskop_Request)

	errRequest := ctx.ShouldBind(request)

	if errRequest != nil {
		ctx.JSON(400, gin.H{
			"Error": "Bad Request :" + errRequest.Error(),
		})
	}

	Bioskop, err := c.service.Update(id, request)

	if err != nil {
		if appErr, ok := err.(*helper.AppError); ok {
			ctx.JSON(appErr.Code, gin.H{
				"Error": appErr.Message,
			})
			return
		}

		ctx.JSON(500, gin.H{
			"Error": "Internal Server Error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"Message": "Success Updated Cinema",
		"Data":    Bioskop,
	})
}

func (c *Controller) Delete(ctx *gin.Context) {
	id := ctx.Param("id")

	err := c.service.Delete(id)

	if err != nil {
		if appErr, ok := err.(*helper.AppError); ok {
			ctx.JSON(appErr.Code, gin.H{
				"Error": appErr.Message,
			})
			return
		}

		ctx.JSON(500, gin.H{
			"Error": "Internal Server Error",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"Message": "Success Delete Cinema",
	})
}
