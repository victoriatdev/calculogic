package main

import (
	"fmt"

	"fyp-server/cmd/handlers"
	"fyp-server/cmd/logic"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	fmt.Println("Starting server...")

	e := echo.New()
	e.HideBanner = true
	e.Use(handlers.LogRequest)

	// CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{
			echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
		},
		AllowMethods: []string{echo.GET, echo.POST},
	}))

	e.GET("/", handlers.Home)

	e.GET("/natural-deduction", handlers.LogicRoot)
	e.POST("/natural-deduction", logic.NaturalDeductionRoot)

	e.POST("/sequent-calculus", logic.SequentCalculusRoot)

	//storage.InitDatabase()

	e.Logger.Fatal(e.Start(":1323"))
}
