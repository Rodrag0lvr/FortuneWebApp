package entities

import "time"

type Liquidation struct {
	ID               int        `json:"id"`
	Departure        string     `json:"departure"`
	Arrival          string     `json:"arrival"`
	Laundry          float64    `json:"laundry"`
	Garage           float64    `json:"garage"`
	Guardianship     float64    `json:"guardianship"`
	Cover            float64    `json:"cover"`
	Sweeper          float64    `json:"sweeper"`
	Driver           string     `json:"driver"`
	Fuel             float64    `json:"fuel"`
	Date             *time.Time `json:"date"`
	Freight          float64    `json:"freight"`
	FreightLiquid    float64    `json:"freight_liquid"`
	Detraction       float64    `json:"detraction"`
	Gremission       string     `json:"gremission"`
	Gtransport       string     `json:"gtransport"`
	Gtransport2      string     `json:"gtransport2"`
	Invoice          string     `json:"invoice"`
	DriverPay        float64    `json:"driver_pay"`
	DriveDescription string     `json:"drive_description"`
	FuelDescription  string     `json:"fuel_description"`
	LiquidTrip       float64    `json:"liquid_trip"`
	Truck            string     `json:"truck"`
	ExpenseTotal     float64    `json:"expense_total"`
	Toll             float64    `json:"toll"`
	GastAdition      bool       `json:"gast_adition"`
}

func NewLiquidation(id int, laundry, garage, guardianship, cover, sweeper, fuel, freight, FreightLiquid, detraction, driverPay, LiquidTrip, expenseTotal, toll float64, driver, departure, arrival, gremission, gtransport, gtransport2, invoice, DriveDescription, fuelDescription, truck string, date time.Time, gastAdition bool) *Liquidation {
	return &Liquidation{
		ID:               id,
		Departure:        departure,
		Arrival:          arrival,
		Laundry:          laundry,
		Garage:           garage,
		Guardianship:     guardianship,
		Cover:            cover,
		Sweeper:          sweeper,
		Driver:           driver,
		Fuel:             fuel,
		Date:             &date,
		Freight:          freight,
		FreightLiquid:    FreightLiquid,
		Detraction:       detraction,
		Gremission:       gremission,
		Gtransport:       gtransport,
		Gtransport2:      gtransport2,
		Invoice:          invoice,
		DriverPay:        driverPay,
		DriveDescription: DriveDescription,
		FuelDescription:  fuelDescription,
		LiquidTrip:       LiquidTrip,
		Truck:            truck,
		ExpenseTotal:     expenseTotal,
		Toll:             toll,
		GastAdition:      gastAdition,
	}
}

func NewFakeLiquidation() *Liquidation {
	time := time.Now()
	return &Liquidation{
		ID:               1,
		Departure:        "lima",
		Arrival:          "chimbote",
		Laundry:          1,
		Garage:           1,
		Guardianship:     1,
		Cover:            1,
		Sweeper:          1,
		Driver:           "aguilar",
		Fuel:             1,
		Date:             &time,
		Freight:          1,
		FreightLiquid:    1,
		Detraction:       1,
		Gremission:       "1",
		Gtransport:       "1",
		Gtransport2:      "1",
		Invoice:          "1",
		DriverPay:        1,
		DriveDescription: "1",
		FuelDescription:  "1",
		LiquidTrip:       1,
		Truck:            "1",
		ExpenseTotal:     1,
		Toll:             1,
		GastAdition:      true,
	}
}
