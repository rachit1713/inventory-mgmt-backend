package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rachit1713/inventory-mgmt-backend/routes"
)

func main() {
	r := gin.Default()

	// Register routes
	routes.RegisterInventoryRoutes(r)

	// Start server
	r.Run(":8080")
}
