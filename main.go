package main

import (
	"fmt"
	"holyways/database"
	"holyways/pkg/mysql"
	"holyways/pkg/routes"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	// CORS agar bisa akses backend
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},                                                 // mengijinkan akses semuanya
		AllowMethods: []string{echo.GET, echo.POST, echo.PATCH, echo.DELETE},        // mengijinkan method apa aja yg bisa digunakan
		AllowHeaders: []string{"X-Requested-With", "Content-Type", "Authorization"}, // mengijinkan headers apa aja yg bisa digunakan
	}))

	routes.RouteInit(e.Group("/api/v1"))

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
