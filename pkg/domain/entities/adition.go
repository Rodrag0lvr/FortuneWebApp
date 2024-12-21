package entities

type Adition struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`

	LiquidationId int `json:"liquidation_id"`
	Liquidation   *Liquidation
}

func NewAdition(id int, description string, price float64, liquidationId int) *Adition {
	return &Adition{
		ID:            id,
		Description:   description,
		Price:         price,
		LiquidationId: liquidationId,
	}
}

func NewFakeAdition() Adition {
	return Adition{
		ID:            1,
		Description:   "description",
		Price:         1,
		LiquidationId: 1,
	}
}
