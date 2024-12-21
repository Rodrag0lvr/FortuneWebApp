package entities

import "time"

type Liquidation struct {
	ID           int       `json:"id"`
	Departure    string    `json:"departure"`
	Arrival      string    `json:"arrival"`
	Weights      float64   `json:"weights"`
	Laundry      float64   `json:"laundry"`
	Garage       float64   `json:"garage"`
	Guardianship float64   `json:"guardianship"`
	Cover        float64   `json:"cover"`
	Sweeper      float64   `json:"sweeper"`
	Driver       string    `json:"driver"`
	Fuel         float64   `json:"fuel"`
	Expense      float64   `json:"expense"`
	Price        float64   `json:"price"`
	Date         time.Time `json:"date"`
}

func NewLiquidation(id int, weights, laundry, garage, guardianship, cover, sweeper, fuel, expense, price float64, driver, departure, arrival string, date time.Time) *Liquidation {
	return &Liquidation{
		ID:           id,
		Departure:    departure,
		Arrival:      arrival,
		Weights:      weights,
		Laundry:      laundry,
		Garage:       garage,
		Guardianship: guardianship,
		Cover:        cover,
		Sweeper:      sweeper,
		Driver:       driver,
		Fuel:         fuel,
		Expense:      expense,
		Price:        price,
		Date:         date,
	}
}

func NewFakeLiquidation() *Liquidation {
	return &Liquidation{
		ID:           1,
		Departure:    "lima",
		Arrival:      "chimbote",
		Weights:      1,
		Laundry:      1,
		Garage:       1,
		Guardianship: 1,
		Cover:        1,
		Sweeper:      1,
		Driver:       "aguilar",
		Fuel:         1,
		Expense:      1,
		Price:        1,
		Date:         time.Now(),
	}
}
