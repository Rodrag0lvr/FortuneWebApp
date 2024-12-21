package repository

import (
	"errors"
	"fortuna-express-web/pkg/domain/entities"
	"fortuna-express-web/pkg/interfaces"
	"time"
)

type InMemoryLiquidationRepository struct {
	liquidations []*entities.Liquidation
}

func NewInMemoryLiquidationRepository() interfaces.LiquidationRepository {
	liquidations := []*entities.Liquidation{
		{
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
			Date:             time.Now(),
			Freight:          1,
			FreightLiquid:    1,
			Detraction:       1,
			DriverPay:        1,
			LiquidTrip:       1,
			Gremission:       "1",
			Gtransport:       "1",
			Gtransport2:      "1",
			Invoice:          "1",
			DriveDescription: "1",
			FuelDescription:  "1",
			Truck:            "1",
			Toll:             1,
		},
		{
			ID:               2,
			Departure:        "chimbote",
			Arrival:          "lima",
			Laundry:          1,
			Garage:           1,
			Guardianship:     1,
			Cover:            1,
			Sweeper:          1,
			Driver:           "damian",
			Fuel:             1,
			Date:             time.Now(),
			Freight:          1,
			FreightLiquid:    1,
			Detraction:       1,
			DriverPay:        1,
			LiquidTrip:       1,
			Gremission:       "1",
			Gtransport:       "1",
			Gtransport2:      "1",
			Invoice:          "1",
			DriveDescription: "1",
			FuelDescription:  "1",
			Truck:            "1",
			Toll:             1,
		},
	}
	return &InMemoryLiquidationRepository{
		liquidations: liquidations,
	}
}

func (l *InMemoryLiquidationRepository) New(liquidation *entities.Liquidation) (int, error) {
	liquidation.ID = len(l.liquidations) + 1
	l.liquidations = append(l.liquidations, liquidation)
	return liquidation.ID, nil
}

func (l *InMemoryLiquidationRepository) List() ([]*entities.Liquidation, error) {
	if len(l.liquidations) == 0 {
		return nil, errors.New("no liquidations available")
	}
	return l.liquidations, nil
}

func (l *InMemoryLiquidationRepository) Get(id int) (*entities.Liquidation, error) {
	for _, liquidation := range l.liquidations {
		if liquidation.ID == id {
			return liquidation, nil
		}
	}
	return nil, errors.New("liquidation not found")
}

func (l *InMemoryLiquidationRepository) Update(liquidation *entities.Liquidation) error {
	for i, existing := range l.liquidations {
		if existing.ID == liquidation.ID {
			l.liquidations[i] = liquidation
			return nil
		}
	}
	return errors.New("liquidation not found")
}

func (l *InMemoryLiquidationRepository) Delete(id int) error {
	for i, liquidation := range l.liquidations {
		if liquidation.ID == id {
			l.liquidations = append(l.liquidations[:i], l.liquidations[i+1:]...)
			return nil
		}
	}
	return errors.New("liquidation not found")
}

type InMemoryAditionRepository struct {
	aditions []*entities.Adition
}

func NewInMemoryAditionRepository() interfaces.AditionRepository {

	aditions := []*entities.Adition{
		{
			ID:            1,
			Description:   "description",
			Price:         1,
			LiquidationId: 1,
		},
	}
	return &InMemoryAditionRepository{
		aditions: aditions,
	}
}
func (a *InMemoryAditionRepository) New(adition *entities.Adition) (int, error) {
	adition.ID = len(a.aditions) + 1
	a.aditions = append(a.aditions, adition)
	return adition.ID, nil
}

func (a *InMemoryAditionRepository) List() ([]*entities.Adition, error) {
	if len(a.aditions) == 0 {
		return nil, errors.New("no aditions available")
	}
	return a.aditions, nil
}

func (a *InMemoryAditionRepository) Get(id int) ([]*entities.Adition, error) {
	var result []*entities.Adition
	for _, adition := range a.aditions {
		if adition.LiquidationId == id {
			result = append(result, adition)
		}
	}

	if len(result) == 0 {
		return nil, errors.New("adition not found")
	}

	return result, nil
}

func (a *InMemoryAditionRepository) Update(adition *entities.Adition) error {
	for i, existing := range a.aditions {
		if existing.ID == adition.ID {
			a.aditions[i] = adition
			return nil
		}
	}
	return errors.New("adition not found")
}

func (a *InMemoryAditionRepository) Delete(id int) error {
	for i, adition := range a.aditions {
		if adition.ID == id {
			a.aditions = append(a.aditions[:i], a.aditions[i+1:]...)
			return nil
		}
	}
	return errors.New("adition not found")
}

func NewInMemoryDescriptionRepository() interfaces.DescriptionRepository {
	return &InMemoryDescriptionRepository{}
}

type InMemoryDescriptionRepository struct {
	descriptions []*entities.Description
}

func (d *InMemoryDescriptionRepository) New(description *entities.Description) (int, error) {
	description.ID = len(d.descriptions) + 1
	d.descriptions = append(d.descriptions, description)
	return description.ID, nil
}

func (d *InMemoryDescriptionRepository) List() ([]*entities.Description, error) {
	if len(d.descriptions) == 0 {
		return nil, errors.New("no descriptions available")
	}
	return d.descriptions, nil
}

func (d *InMemoryDescriptionRepository) Get(id int) (*entities.Description, error) {
	for _, description := range d.descriptions {
		if description.ID == id {
			return description, nil
		}
	}
	return nil, errors.New("description not found")
}

func (d *InMemoryDescriptionRepository) Update(description *entities.Description) error {
	for i, existing := range d.descriptions {
		if existing.ID == description.ID {
			d.descriptions[i] = description
			return nil
		}
	}
	return errors.New("description not found")
}

func (d *InMemoryDescriptionRepository) Delete(id int) error {
	for i, description := range d.descriptions {
		if description.ID == id {
			d.descriptions = append(d.descriptions[:i], d.descriptions[i+1:]...)
			return nil
		}
	}
	return errors.New("description not found")
}
