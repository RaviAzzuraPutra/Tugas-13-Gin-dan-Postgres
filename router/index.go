package router

import (
	"formative-13/controller"

	"github.com/gin-gonic/gin"
)

func BioskopRouter(app *gin.Engine, BioskopController *controller.Controller) {

	app.GET("/", BioskopController.Get)
	app.GET("/:id", BioskopController.GetById)
	app.POST("/bioskop", BioskopController.Create)
	app.PUT("/update/:id", BioskopController.Update)
	app.DELETE("/delete/:id", BioskopController.Delete)
}
