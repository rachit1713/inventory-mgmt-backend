package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rachit1713/inventory-mgmt-backend/handlers"
)

func RegisterInventoryRoutes(r *gin.Engine) {
	r.GET("/inventories", handlers.GetInventories)
	r.POST("/inventories", handlers.CreateInventory)
	r.PUT("/inventories/:id", handlers.UpdateInventory)
	r.DELETE("/inventories/:id", handlers.DeleteInventory)
	r.POST("/inventories/update-used", handlers.BulkUpdateInventories)
	r.GET("/inventories/ending", handlers.GetEndingInventories)
}
