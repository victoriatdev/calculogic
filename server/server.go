package main

import (
	"fmt"
	"log"
	"net/http"
//	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
//	setConfig()

	fmt.Println("Starting server...")

	e := echo.New()
	e.HideBanner = true

	//test_env := os.Getenv("TEST_SECRET")
	//fmt.Println(test_env)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/users/:id", getUser)

	e.Logger.Fatal(e.Start(":1323"))
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func setConfig() {
	envError := godotenv.Load(".env")

	if envError != nil {
		log.Fatal("Error loading .env file.")
	}
}
