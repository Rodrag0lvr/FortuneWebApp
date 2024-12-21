package handlers

import (
	"fortuna-express-web/pkg/domain/entities"
	uc "fortuna-express-web/pkg/domain/usecases"
	"fortuna-express-web/pkg/web"
	"log"
	"log/slog"
	"net/http"
	"path/filepath"
	"strconv"
	"text/template"
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

func (l liquidationsHandler) New(user *entities.User, w http.ResponseWriter, r *http.Request) {
	// Obtén los valores del formulario
	departure := r.FormValue("departure")
	arrival := r.FormValue("arrival")
	weights := r.FormValue("weights")
	laundry := r.FormValue("laundry")
	garage := r.FormValue("garage")
	guardianship := r.FormValue("guardianship")
	cover := r.FormValue("cover")
	sweeper := r.FormValue("sweeper")
	driver := r.FormValue("driver")
	fuel := r.FormValue("fuel")
	expense := r.FormValue("expense")
	price := r.FormValue("price")

	// Convierte las entradas de string a float64
	departureFloat, err := strconv.ParseFloat(departure, 64)
	if err != nil {
		http.Error(w, "Invalid departure value", http.StatusBadRequest)
		return
	}

	arrivalFloat, err := strconv.ParseFloat(arrival, 64)
	if err != nil {
		http.Error(w, "Invalid arrival value", http.StatusBadRequest)
		return
	}

	weightsFloat, err := strconv.ParseFloat(weights, 64)
	if err != nil {
		http.Error(w, "Invalid weights value", http.StatusBadRequest)
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

	driverFloat, err := strconv.ParseFloat(driver, 64)
	if err != nil {
		http.Error(w, "Invalid driver value", http.StatusBadRequest)
		return
	}

	fuelFloat, err := strconv.ParseFloat(fuel, 64)
	if err != nil {
		http.Error(w, "Invalid fuel value", http.StatusBadRequest)
		return
	}

	expenseFloat, err := strconv.ParseFloat(expense, 64)
	if err != nil {
		http.Error(w, "Invalid expense value", http.StatusBadRequest)
		return
	}

	priceFloat, err := strconv.ParseFloat(price, 64)
	if err != nil {
		http.Error(w, "Invalid price value", http.StatusBadRequest)
		return
	}

	// Crea la instancia de Liquidation
	liquidation := entities.Liquidation{
		Departure:    departureFloat,
		Arrival:      arrivalFloat,
		Weights:      weightsFloat,
		Laundry:      laundryFloat,
		Garage:       garageFloat,
		Guardianship: guardianshipFloat,
		Cover:        coverFloat,
		Sweeper:      sweeperFloat,
		Driver:       driverFloat,
		Fuel:         fuelFloat,
		Expense:      expenseFloat,
		Price:        priceFloat,
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
