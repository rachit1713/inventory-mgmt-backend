package models

type Inventory struct {
	ID              string  `json:"id"`
	Name            string  `json:"name"`
	Quantity        float64 `json:"quantity"`
	QuantityType    string  `json:"quantityType"`
	Location        string  `json:"location"`
	MinQuantity     float64 `json:"minQuantity"`
	UnitConsumption float64 `json:"unitConsumption"`
	UnitsUsed       int     `json:"unitsUsed"`
}
