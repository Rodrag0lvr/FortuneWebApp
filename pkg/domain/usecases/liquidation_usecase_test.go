package uc

import (
	repository "fortuna-express-web/pkg/repositories"
)

func setupTest() LiquidationUseCase {
	liquidationRepo := repository.NewInMemoryLiquidationRepository()
	descriptionRepo := repository.NewInMemoryDescriptionRepository()
	aditionRepo := repository.NewInMemoryAditionRepository()
	return NewLiquidationUseCase(liquidationRepo, descriptionRepo, aditionRepo)
}

/*
func TestNewLiquidation(t *testing.T) {
	uc := setupTest()
	user := entities.NewFakeUser()
	liquidation := entities.NewFakeLiquidation()
	err := uc.New(&user, liquidation)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	liquidations, err := uc.List(&user)

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	assert.Equal(t, 2, len(liquidations))

}

func TestListLiquidation(t *testing.T) {
	uc := setupTest()
	user := entities.NewFakeUser()
	liquidations, err := uc.List(&user)

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	assert.Equal(t, 1, len(liquidations))
}

func TestGetLiquidation(t *testing.T) {
	uc := setupTest()
	user := entities.NewFakeUser()
	liquidation, err := uc.Get(&user, 1)

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	assert.Equal(t, 1, liquidation.ID)
}
func TestUpdateLiquidation(t *testing.T) {
	uc := setupTest()
	user := entities.NewFakeUser()
	liquidation := entities.NewFakeLiquidation()
	liquidation.Cover = 1800
	err := uc.Update(&user, liquidation)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	liquidation, err = uc.Get(&user, 1)

	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
	assert.Equal(t, float64(1800), liquidation.Cover)
}

func TestDeleteLiquidation(t *testing.T) {
	uc := setupTest()
	user := entities.NewFakeUser()
	err := uc.Delete(&user, 1)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}

	liquidations, _ := uc.List(&user)

	assert.Equal(t, 0, len(liquidations))
}
*/
