package bootstrap

import (
	"formative-13/config"
	"formative-13/config/app_config"
	"formative-13/controller"
	"formative-13/database"
	"formative-13/repository"
	"formative-13/router"
	"formative-13/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func InitApp() {
	_ = godotenv.Load()

	config.Init_Config()

	database.Connect()

	app := gin.Default()

	repository := repository.NewRpositoryRegistry()
	service := service.NewServiceRegistry(repository)
	controller := controller.NewControllerRegistry(service)

	router.BioskopRouter(app, controller)

	app.Run(app_config.PORT)
}
