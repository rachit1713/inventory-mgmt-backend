package data

import "github.com/rachit1713/inventory-mgmt-backend/models"

var MockInventories = []models.Inventory{
	{
		ID:              "1",
		Name:            "Rice",
		Quantity:        10.5,
		QuantityType:    "kg",
		Location:        "Kitchen",
		MinQuantity:     2.0,
		UnitConsumption: 0.5,
		UnitsUsed:       0,
	},
	{
		ID:              "2",
		Name:            "Milk",
		Quantity:        5.0,
		QuantityType:    "liters",
		Location:        "Fridge",
		MinQuantity:     1.0,
		UnitConsumption: 0.25,
		UnitsUsed:       0,
	},
}
