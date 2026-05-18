package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rachit1713/inventory-mgmt-backend/data"
	"github.com/rachit1713/inventory-mgmt-backend/models"
)

var inventories = data.MockInventories // start with mock data

func GetInventories(c *gin.Context) {
	search := c.Query("search")
	var result []models.Inventory

	if search == "" {
		result = inventories
	} else {
		for _, inv := range inventories {
			if strings.Contains(strings.ToLower(inv.Name), strings.ToLower(search)) {
				result = append(result, inv)
			}
		}
	}

	c.JSON(http.StatusOK, result)
}

type UpdateRequest struct {
	ID        string `json:"id"`
	UnitsUsed int    `json:"unitsUsed"`
}

func GetEndingInventories(c *gin.Context) {
	ending := []models.Inventory{}
	for _, inv := range inventories {
		if inv.Quantity <= inv.MinQuantity {
			ending = append(ending, inv)
		}
	}
	c.JSON(http.StatusOK, ending)
}

func BulkUpdateInventories(c *gin.Context) {
	var updates []UpdateRequest
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated := []models.Inventory{}
	for _, req := range updates {
		for i := range inventories {
			if inventories[i].ID == req.ID {
				inventories[i].Quantity -= float64(req.UnitsUsed) * inventories[i].UnitConsumption
				if inventories[i].Quantity < 0 {
					inventories[i].Quantity = 0
				}
				inventories[i].UnitsUsed = 0
				updated = append(updated, inventories[i])
			}
		}
	}

	c.JSON(http.StatusOK, updated)
}

func CreateInventory(c *gin.Context) {
	var newItem models.Inventory
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	inventories = append(inventories, newItem)
	c.JSON(http.StatusCreated, newItem)
}

func UpdateInventory(c *gin.Context) {
	id := c.Param("id")
	var updated models.Inventory
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, inv := range inventories {
		if inv.ID == id {
			inventories[i] = updated
			c.JSON(http.StatusOK, updated)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}

func DeleteInventory(c *gin.Context) {
	id := c.Param("id")
	for i, inv := range inventories {
		if inv.ID == id {
			inventories = append(inventories[:i], inventories[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}
