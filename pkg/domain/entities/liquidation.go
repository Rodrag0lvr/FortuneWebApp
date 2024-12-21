package entities

type Liquidation struct {
	ID           int     `json:"id"`
	Departure    float64 `json:"departure"`
	Arrival      float64 `json:"arrival"`
	Weights      float64 `json:"weights"`
	Laundry      float64 `json:"laundry"`
	Garage       float64 `json:"garage"`
	Guardianship float64 `json:"guardianship"`
	Cover        float64 `json:"cover"`
	Sweeper      float64 `json:"sweeper"`
	Driver       float64 `json:"driver"`
	Fuel         float64 `json:"fuel"`
	Expense      float64 `json:"expense"`
	Price        float64 `json:"price"`
}

func NewLiquidation(id int, departure, arrival, weights, laundry, garage, guardianship, cover, sweeper, driver, fuel, expense, price float64) *Liquidation {
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
	}
}

func NewFakeLiquidation() *Liquidation {
	return &Liquidation{
		ID:           1,
		Departure:    1,
		Arrival:      1,
		Weights:      1,
		Laundry:      1,
		Garage:       1,
		Guardianship: 1,
		Cover:        1,
		Sweeper:      1,
		Driver:       1,
		Fuel:         1,
		Expense:      1,
		Price:        1,
	}
}
