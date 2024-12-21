package entities

type Description struct {
	ID            int    `json:"id"`
	Input         string `json:"input"`
	Description   string `json:"description"`
	LiquidationId int    `json:"liquidation_id"`
	Liquidation   *Liquidation
}

func NewDescription(id int, input, description string, liquidationId int) *Description {
	return &Description{
		ID:            id,
		Input:         input,
		Description:   description,
		LiquidationId: liquidationId,
	}
}

func NewFakeDescription() Description {
	return Description{
		ID:            1,
		Input:         "Driver",
		Description:   "description",
		LiquidationId: 1,
	}
}
