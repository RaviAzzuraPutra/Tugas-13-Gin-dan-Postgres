package config

import (
	"formative-13/config/app_config"
	"formative-13/config/db_config"
)

func Init_Config() {
	app_config.App_Connfig()
	db_config.DB_Config()
}
