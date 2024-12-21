package handlers

import (
	"fmt"
	"fortuna-express-web/pkg/domain/entities"
	uc "fortuna-express-web/pkg/domain/usecases"
	"fortuna-express-web/pkg/web"
	"log"
	"log/slog"
	"net/http"
	"path/filepath"
	"strconv"
	"text/template"
	"time"
)

var templateBasePath = "public/pages/"

type liquidationsHandler struct {
	logger *slog.Logger
	uc     uc.LiquidationUseCase
}

func Render(w http.ResponseWriter, templateFile string, data interface{}) error {
	// Construye la ruta completa al archivo
	fullPath := filepath.Join(web.TemplateBasePath, templateFile)
	log.Println("Attempting to load template from:", fullPath)

	// Cargar la plantilla
	tmpl, err := template.ParseFiles(fullPath)
	if err != nil {
		log.Printf("Error loading template: %v\n", err)
		return err
	}

	// Renderizar la plantilla con los datos proporcionados
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("Error executing template: %v\n", err)
		return err
	}

	return nil
}
func NewLiquidationsHandler(logger *slog.Logger, uc uc.LiquidationUseCase) web.LiquidationsHandler {
	return &liquidationsHandler{
		logger: logger,
		uc:     uc,
	}
}
func (l liquidationsHandler) Get(user *entities.User, w http.ResponseWriter, r *http.Request) (map[string]interface{}, error) {
	// Obtener el ID de la liquidación desde la query string
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		return nil, fmt.Errorf("id is required")
	}

	// Convertir el ID a entero
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, fmt.Errorf("invalid id")
	}

	// Obtener la liquidación
	liquidation, err := l.uc.Get(user, id)
	if err != nil {
		return nil, err
	}

	// Estructurar los datos para la vista
	data := map[string]interface{}{
		"user":        user,
		"liquidation": liquidation,
	}
	log.Printf("--------------------------Liquidation: %+v\n", data)

	return data, nil
}

func (l liquidationsHandler) HomeView(user *entities.User, w http.ResponseWriter, r *http.Request) (map[string]interface{}, error) {
	// Obtener la lista de liquidaciones asociadas al usuario
	liquidations, err := l.uc.List(user)
	if err != nil {

		http.Error(w, "Failed to fetch liquidations: "+err.Error(), http.StatusInternalServerError)
		return nil, err
	}

	// Estructurar los datos para la vista
	data := map[string]interface{}{
		"user":         user,
		"liquidations": liquidations,
	}

	return data, nil
}

func (l liquidationsHandler) NewView(user *entities.User, w http.ResponseWriter, r *http.Request) {
	// Cargar y renderizar las plantillas
	err := Render(w, "new.html", nil)
	if err != nil {
		log.Println("Error rendering new.html:", err)
		http.Error(w, "Failed to render new.html", http.StatusInternalServerError)
	}
}
func (l liquidationsHandler) New(user *entities.User, w http.ResponseWriter, r *http.Request) {
	// Obtén los valores del formulario
	departure := r.FormValue("departure")
	arrival := r.FormValue("arrival")
	laundry := r.FormValue("laundry")
	garage := r.FormValue("garage")
	guardianship := r.FormValue("guardianship")
	cover := r.FormValue("cover")
	sweeper := r.FormValue("sweeper")
	driver := r.FormValue("driver")
	fuel := r.FormValue("fuel")
	freight := r.FormValue("freight")
	freightLiquid := r.FormValue("freight_liquid")
	detraction := r.FormValue("detraction")
	gremission := r.FormValue("gremission")
	gtransport := r.FormValue("gtransport")
	gtransport2 := r.FormValue("gtransport2")
	invoice := r.FormValue("invoice")
	driverPay := r.FormValue("driver_pay")
	driveDescription := r.FormValue("drive_description")
	fuelDescription := r.FormValue("fuel_description")
	liquidTrip := r.FormValue("liquid_trip")
	expenseTotal := r.FormValue("expense_total")
	truck := r.FormValue("truck")
	toll := r.FormValue("peaje")

	tollFloat, err := strconv.ParseFloat(toll, 64)
	if err != nil {
		http.Error(w, "Invalid toll value", http.StatusBadRequest)
		return
	}

	log.Printf("---------------------%s", expenseTotal)
	expenseTotalFloat, err := strconv.ParseFloat(expenseTotal, 64)
	if err != nil {
		http.Error(w, "Invalid expense total value", http.StatusBadRequest)
		return
	}

	dateStr := r.FormValue("date")
	if dateStr == "" {
		http.Error(w, "date is required", http.StatusBadRequest)
		return
	}

	// Parsear la fecha al formato deseado
	layout := "2006-01-02" // Cambia según el formato esperado del input
	parsedDate, err := time.Parse(layout, dateStr)
	if err != nil {
		http.Error(w, "invalid date format", http.StatusBadRequest)
		return
	}
	freightFloat, err := strconv.ParseFloat(freight, 64)
	if err != nil {
		http.Error(w, "Invalid freight value", http.StatusBadRequest)
		return
	}

	laundryFloat, err := strconv.ParseFloat(laundry, 64)
	if err != nil {
		http.Error(w, "Invalid laundry value", http.StatusBadRequest)
		return
	}

	garageFloat, err := strconv.ParseFloat(garage, 64)
	if err != nil {
		http.Error(w, "Invalid garage value", http.StatusBadRequest)
		return
	}

	guardianshipFloat, err := strconv.ParseFloat(guardianship, 64)
	if err != nil {
		http.Error(w, "Invalid guardianship value", http.StatusBadRequest)
		return
	}

	coverFloat, err := strconv.ParseFloat(cover, 64)
	if err != nil {
		http.Error(w, "Invalid cover value", http.StatusBadRequest)
		return
	}

	sweeperFloat, err := strconv.ParseFloat(sweeper, 64)
	if err != nil {
		http.Error(w, "Invalid sweeper value", http.StatusBadRequest)
		return
	}

	fuelFloat, err := strconv.ParseFloat(fuel, 64)
	if err != nil {
		http.Error(w, "Invalid fuel value", http.StatusBadRequest)
		return
	}

	freightLiquidFloat, err := strconv.ParseFloat(freightLiquid, 64)
	if err != nil {
		http.Error(w, "Invalid freight liquid value", http.StatusBadRequest)
		return
	}
	detractionFloat, err := strconv.ParseFloat(detraction, 64)
	if err != nil {
		http.Error(w, "Invalid detraction value", http.StatusBadRequest)
		return
	}
	driverPayFloat, err := strconv.ParseFloat(driverPay, 64)
	if err != nil {
		http.Error(w, "Invalid driver pay value", http.StatusBadRequest)
		return
	}
	liquidTripFloat, err := strconv.ParseFloat(liquidTrip, 64)
	if err != nil {
		http.Error(w, "Invalid liquid trip value", http.StatusBadRequest)
		return
	}

	calculoTotal := 0.0

	if laundryFloat > 1 {
		calculoTotal += laundryFloat
	}
	if sweeperFloat > 1 {
		calculoTotal += sweeperFloat
	}
	if coverFloat > 1 {
		calculoTotal += coverFloat
	}

	if fuelFloat > 1 {
		calculoTotal += fuelFloat
	}
	calculoTotal += garageFloat + guardianshipFloat + driverPayFloat

	// Crea la instancia de Liquidation
	liquidation := entities.Liquidation{
		Departure:        departure,
		Arrival:          arrival,
		Laundry:          laundryFloat, //
		Garage:           garageFloat,
		Guardianship:     guardianshipFloat,
		Cover:            coverFloat,   //
		Sweeper:          sweeperFloat, //
		Driver:           driver,
		Fuel:             fuelFloat,          //
		Freight:          freightFloat,       //?
		FreightLiquid:    freightLiquidFloat, //?
		Detraction:       detractionFloat,    //?
		Gremission:       gremission,
		Gtransport:       gtransport,
		Gtransport2:      gtransport2,
		Invoice:          invoice,
		DriverPay:        driverPayFloat,
		DriveDescription: driveDescription,
		FuelDescription:  fuelDescription,
		LiquidTrip:       liquidTripFloat,
		Date:             parsedDate,
		Truck:            truck,
		ExpenseTotal:     expenseTotalFloat,
		Toll:             tollFloat,
	}

	err = l.uc.New(user, &liquidation)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func (l liquidationsHandler) Update(user *entities.User, w http.ResponseWriter, r *http.Request) {

}
func (l liquidationsHandler) Delete(user *entities.User, w http.ResponseWriter, r *http.Request) {

}
