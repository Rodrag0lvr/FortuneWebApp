package interfaces

import "fortuna-express-web/pkg/domain/entities"

type LiquidationRepository interface {
	New(liquidation *entities.Liquidation) (int, error)
	List() ([]*entities.Liquidation, error)
	Get(id int) (*entities.Liquidation, error)
	Update(liquidation *entities.Liquidation) error
	Delete(id int) error
}

type DescriptionRepository interface {
	New(description *entities.Description) (int, error)
	List() ([]*entities.Description, error)
	Get(id int) (*entities.Description, error)
	Update(adition *entities.Description) error
	Delete(id int) error
}

type AditionRepository interface {
	New(adition *entities.Adition) (int, error)
	List() ([]*entities.Adition, error)
	Get(id int) ([]*entities.Adition, error)
	Update(adition *entities.Adition) error
	Delete(id int) error
}
