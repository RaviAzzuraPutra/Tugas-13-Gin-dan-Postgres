package router

import (
	"formative-13/controller"

	"github.com/gin-gonic/gin"
)

func BioskopRouter(app *gin.Engine, BioskopController *controller.Controller) {

	bioskop := app.Group("/bioskop")

	bioskop.GET("/", BioskopController.Get)
	bioskop.GET("/:id", BioskopController.GetById)
	bioskop.POST("/", BioskopController.Create)
	bioskop.PUT("/:id", BioskopController.Update)
	bioskop.DELETE("/:id", BioskopController.Delete)
}
