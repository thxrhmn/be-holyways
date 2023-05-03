package main

import (
	"fmt"
	"holyways/database"
	"holyways/pkg/mysql"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	// env
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	mysql.DatabaseInit()
	database.RunMigration()

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		c.Response().WriteHeader(http.StatusOK)
		return c.String(http.StatusOK, "Hello World")
	})

	// fmt.Println("Server running on localhost:5000")
	// e.Logger.Fatal(e.Start("localhost:5000"))

	var PORT = os.Getenv("PORT")

	fmt.Println("server running localhost:" + PORT)
	e.Logger.Fatal(e.Start(":" + PORT)) // delete localhost
}
