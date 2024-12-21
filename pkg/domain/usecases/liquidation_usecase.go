package uc

import (
	"errors"
	"fortuna-express-web/pkg/domain/entities"
	"fortuna-express-web/pkg/interfaces"
	"log"
)

type LiquidationUseCase interface {
	New(user *entities.User, liquidation *entities.Liquidation) (int, error)
	NewDescription(user *entities.User, description *entities.Description) error
	NewAdition(user *entities.User, adition *entities.Adition) error
	List(user *entities.User) ([]*entities.Liquidation, error)
	Get(user *entities.User, id int) (*entities.Liquidation, []*entities.Adition, error)
	Update(user *entities.User, liquidation *entities.Liquidation) error
	Delete(user *entities.User, id int) error
}

type liquidationUC struct {
	repo            interfaces.LiquidationRepository
	repoDescription interfaces.DescriptionRepository
	repoAdition     interfaces.AditionRepository
}

var (
	ErrUnauthorized = errors.New("unauthorized")
)

func NewLiquidationUseCase(repo interfaces.LiquidationRepository, repoDescription interfaces.DescriptionRepository, repoAdition interfaces.AditionRepository) LiquidationUseCase {
	return &liquidationUC{
		repo:            repo,
		repoDescription: repoDescription,
		repoAdition:     repoAdition,
	}
}

func (l *liquidationUC) NewAdition(user *entities.User, adition *entities.Adition) error {
	if user.Role != "admin" {
		return ErrUnauthorized
	}
	if adition == nil {
		return errors.New("adition is nil")
	}

	_, err := l.repoAdition.New(adition)
	if err != nil {

		return err

	}
	return nil
}

func (l *liquidationUC) NewDescription(user *entities.User, description *entities.Description) error {
	if user.Role != "admin" {
		return ErrUnauthorized
	}
	if description == nil {
		return errors.New("description is nil")
	}

	_, err := l.repoDescription.New(description)
	if err != nil {

	}
	return nil
}

func (l *liquidationUC) New(user *entities.User, liquidation *entities.Liquidation) (int, error) {
	if user.Role != "admin" {
		return 0, ErrUnauthorized
	}
	if liquidation == nil {
		return 0, errors.New("liquidation is nil")
	}

	id, err := l.repo.New(liquidation)
	if err != nil {

		return 0, err

	}
	return id, nil
}

func (l *liquidationUC) List(user *entities.User) ([]*entities.Liquidation, error) {
	if user.Role != "admin" {
		return nil, ErrUnauthorized
	}

	liquidations, err := l.repo.List()

	if err != nil {
		return nil, err
	}

	return liquidations, nil
}
func (l *liquidationUC) Get(user *entities.User, id int) (*entities.Liquidation, []*entities.Adition, error) {
	// Validar permisos del usuario
	if user.Role != "admin" {
		return nil, nil, ErrUnauthorized
	}

	// Obtener la liquidación
	liquidation, err := l.repo.Get(id)
	if err != nil {
		return nil, nil, err
	}

	// Si la liquidación tiene gastos adicionales, obtenerlos
	var aditions []*entities.Adition
	if liquidation.GastAdition {
		aditions, err = l.repoAdition.Get(id)
		if err != nil {
			log.Printf("Error obteniendo gastos adicionales: %v", err)
			// Continuar con la liquidación incluso si los gastos adicionales no se encuentran
		} else {

			log.Printf("0000000Gastos adicionales obtenidos: %v", len(aditions))
		}
	}

	// Retornar la liquidación y los gastos adicionales
	return liquidation, aditions, nil
}

func (l *liquidationUC) Update(user *entities.User, liquidation *entities.Liquidation) error {
	if user.Role != "admin" {
		return ErrUnauthorized
	}
	if liquidation == nil {
		return errors.New("liquidation is nil")
	}
	err := l.repo.Update(liquidation)
	if err != nil {
		return err
	}

	return nil
}

func (l *liquidationUC) Delete(user *entities.User, id int) error {
	if user.Role != "admin" {
		return ErrUnauthorized
	}

	if id == 0 {
		return errors.New("id is 0")
	}

	err := l.repo.Delete(id)
	if err != nil {

		return err
	}

	return nil
}
