package main

import (
	"fmt"
	"moonlay-test/migrations"
	"moonlay-test/pkg/database"
	"moonlay-test/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	database.DatabaseInit()
	migrations.RunMigration()

	routes.RouteInit(e.Group("/api/v1"))

	e.Static("/storages", "./storages")

	fmt.Println("Server running on localhost:5000")
	e.Logger.Fatal(e.Start("localhost:5000"))
}
