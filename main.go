package main

import (
	"fmt"
	"landtick-be/database"
	"landtick-be/pkg/mysql"
	"landtick-be/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	e := echo.New()

	mysql.DatabaseInit()
	database.RunMigration()

	routes.RouteInit(e.Group("/api/v1"))

	e.Static("/uploads", "./uploads")

	fmt.Println("Server Running in localhost:5000")
	e.Logger.Fatal(e.Start("localhost:5000"))
}
