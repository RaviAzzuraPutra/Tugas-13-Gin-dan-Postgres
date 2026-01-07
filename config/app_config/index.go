package app_config

import "os"

var PORT string

func App_Connfig() {
	PORT = os.Getenv("PORT")
}
