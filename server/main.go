package main

import (
	"echo-rest-api/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(200, &echo.Map{"data": "Yo Bitch"})
	})

	//routes
	routes.ContainerRoute(e) //add this

	e.Logger.Fatal(e.Start(":8080"))
}
