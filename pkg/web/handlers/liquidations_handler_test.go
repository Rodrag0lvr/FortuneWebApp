package handlers

/*
var handler web.LiquidationsHandler

func setup() {
	logger := slog.Default()
	liquidationRepo := repository.NewInMemoryLiquidationRepository()
	aditionRepo := repository.NewInMemoryAditionRepository()
	uc := uc.NewLiquidationUseCase(liquidationRepo, aditionRepo)
	handler = NewLiquidationsHandler(logger, uc)

}

func TestLiquidationHandlerNew(t *testing.T) {

	setup()
	actor := entities.NewFakeUser()
	form := url.Values{}

	form.Add("departure", "100.5")
	form.Add("arrival", "200.5")
	form.Add("weights", "50.3")
	form.Add("laundry", "30.0")
	form.Add("garage", "20.0")
	form.Add("guardianship", "40.0")
	form.Add("cover", "10.0")
	form.Add("sweeper", "15.0")
	form.Add("driver", "25.0")
	form.Add("fuel", "60.0")
	form.Add("expense", "500.0")
	form.Add("price", "1000.0")

	req, err := http.NewRequest("POST", "/liquidations/new", bytes.NewBufferString(form.Encode()))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	rr := httptest.NewRecorder()

	// Ejecutar el controlador
	handler.New(&actor, rr, req)
	if status := rr.Code; status != http.StatusSeeOther {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusSeeOther)
	}
}
*/
