package app

import (
	"http-server/internal/controllers"
	"http-server/internal/router"
	"http-server/internal/usecases"
)

var httpRouter router.Router = router.NewMuxRouter()
var port = ":8001"

func Run() {
	substring := usecases.NewSubstring()
	substringController := controllers.NewSubstringController(substring)
	httpRouter.POST("/api/substring", substringController.GetString)
	httpRouter.SERVE(port)
}
