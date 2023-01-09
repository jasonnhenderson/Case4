package routes

import (
	"echo-rest-api/controllers"

	"github.com/labstack/echo/v4"
)

func ContainerRoute(e *echo.Echo) {
	//All routes related to users comes here
	e.GET("/run-container", controllers.RunContainer)
	e.GET("/list-container", controllers.RunContainer)
	e.GET("/stop-container", controllers.RunContainer)
	e.GET("/pull-container", controllers.RunContainer)
	e.GET("/commit-container", controllers.RunContainer)
}
