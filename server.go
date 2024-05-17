package main

import (
	"elemento-api/config"
	"elemento-api/routes"
	"fmt"
)

func main() {
	db := config.InitDB()
	config.AutoMigration(db)

	// Route
	route, e := routes.Init()

	routes.RouteAuth(route, db)

	// Start server
	port := 8080
	address := fmt.Sprintf("0.0.0.0:%d", port)
	e.Logger.Fatal(e.Start(address))
}
