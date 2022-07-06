package main

import (
	"fmt"
	"go-mongo-one/controllers"

	"github.com/kamva/mgm/v3"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	_ = mgm.SetDefaultConfig(
		nil, "oneD", options.Client().ApplyURI("mongodb://localhost:27017/?readPreference=primary&appname=MongoDB%20Compass&directConnection=true&ssl=false"),
	)
}

func main() {
	server := echo.New()
	server.Debug = true

	server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	// server.GET("/", func(ctx echo.Context) error {
	// 	ctx.JSON(http.StatusOK, { "running": true, "dev": "efen" })
	// })
	server.POST("/book", controllers.Create)
	server.GET("/book/:id", controllers.Get)
	server.GET("/book", controllers.GetAll)

	if err := server.Start(":1234"); err != nil {
		fmt.Println("Error running server")
	}

}
