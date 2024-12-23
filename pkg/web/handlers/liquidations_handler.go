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
	liquidation, adition, err := l.uc.Get(user, id)
	if err != nil {
		return nil, err
	}

	// Estructurar los datos para la vista
	data := map[string]interface{}{
		"user":        user,
		"liquidation": liquidation,
		"adition":     adition,
	}

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
	// Procesa el formulario
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	// Procesar otros valores del formulario
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

	// Conversión de valores numéricos
	tollFloat, err := strconv.ParseFloat(toll, 64)
	if err != nil {
		http.Error(w, "Invalid toll value", http.StatusBadRequest)
		return
	}

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

	layout := "2006-01-02"
	parsedDate, err := time.Parse(layout, dateStr)
	if err != nil {
		http.Error(w, "invalid date format", http.StatusBadRequest)
		return
	}

	freightFloat, _ := strconv.ParseFloat(freight, 64)
	laundryFloat, _ := strconv.ParseFloat(laundry, 64)
	garageFloat, _ := strconv.ParseFloat(garage, 64)
	guardianshipFloat, _ := strconv.ParseFloat(guardianship, 64)
	coverFloat, _ := strconv.ParseFloat(cover, 64)
	sweeperFloat, _ := strconv.ParseFloat(sweeper, 64)
	fuelFloat, _ := strconv.ParseFloat(fuel, 64)
	freightLiquidFloat, _ := strconv.ParseFloat(freightLiquid, 64)
	detractionFloat, _ := strconv.ParseFloat(detraction, 64)
	driverPayFloat, _ := strconv.ParseFloat(driverPay, 64)
	liquidTripFloat, _ := strconv.ParseFloat(liquidTrip, 64)
	gastosDescripcion := r.Form["gasto-descripcion[]"]
	gastosMonto := r.Form["gasto-monto[]"]

	bandera := false
	if len(gastosDescripcion) != 0 {
		bandera = true
	}

	GastoTotal := 0.0

	if laundryFloat > 1 {
		GastoTotal += laundryFloat
	}
	if sweeperFloat > 1 {
		GastoTotal += sweeperFloat
	}
	if coverFloat > 1 {
		GastoTotal += coverFloat
	}

	if fuelFloat > 1 {
		GastoTotal += fuelFloat
	}

	GastoTotal += tollFloat + garageFloat + guardianshipFloat + driverPayFloat

	// Crear instancia de Liquidation
	liquidation := entities.Liquidation{
		Departure:        departure,
		Arrival:          arrival,
		Laundry:          laundryFloat,
		Garage:           garageFloat,
		Guardianship:     guardianshipFloat,
		Cover:            coverFloat,
		Sweeper:          sweeperFloat,
		Driver:           driver,
		Fuel:             fuelFloat,
		Freight:          freightFloat,
		FreightLiquid:    freightLiquidFloat,
		Detraction:       detractionFloat,
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
		GastAdition:      bandera,
	}

	// Guardar la liquidación
	id, err := l.uc.New(user, &liquidation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Obtener los valores de la lista dinámica de gastos adicionales

	// Validar que ambas listas tengan la misma longitud
	if len(gastosDescripcion) != len(gastosMonto) {
		http.Error(w, "Mismatch between gasto-descripcion and gasto-monto", http.StatusBadRequest)
		return
	}

	// Procesar los gastos adicionales
	for i := range gastosDescripcion {
		monto, err := strconv.ParseFloat(gastosMonto[i], 64)
		if err != nil {
			http.Error(w, "Invalid monto value for additional expense", http.StatusBadRequest)
			return
		}

		adition := entities.Adition{
			Description:   gastosDescripcion[i],
			Price:         monto,
			LiquidationId: id,
		}
		err = l.uc.NewAdition(user, &adition)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	// Redirigir al home después de guardar
	http.Redirect(w, r, "/home", http.StatusSeeOther)
}

func (l liquidationsHandler) Update(user *entities.User, w http.ResponseWriter, r *http.Request) {

}
func (l liquidationsHandler) Delete(user *entities.User, w http.ResponseWriter, r *http.Request) {

}
