package uc

import (
	"errors"
	"fortuna-express-web/pkg/domain/entities"
	"fortuna-express-web/pkg/interfaces"
	"log"
)

type LiquidationUseCase interface {
	New(user *entities.User, liquidation *entities.Liquidation) error
	NewDescription(user *entities.User, description *entities.Description) error
	NewAdition(user *entities.User, adition *entities.Adition) error
	List(user *entities.User) ([]*entities.Liquidation, error)
	Get(user *entities.User, id int) (*entities.Liquidation, error)
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

func (l *liquidationUC) New(user *entities.User, liquidation *entities.Liquidation) error {
	if user.Role != "admin" {
		return ErrUnauthorized
	}
	if liquidation == nil {
		return errors.New("liquidation is nil")
	}

	err := l.repo.New(liquidation)
	if err != nil {

		return err

	}
	return nil
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

func (l *liquidationUC) Get(user *entities.User, id int) (*entities.Liquidation, error) {
	if user.Role != "admin" {
		return nil, ErrUnauthorized
	}

	log.Printf("-------------------ENTROOOOOO----------------")
	liquidation, err := l.repo.Get(id)
	if err != nil {
		return nil, err
	}

	return liquidation, nil
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
