package main

import (
	"database/sql"
	"fmt"
	"log"
	"log/slog"

	uc "fortuna-express-web/pkg/domain/usecases"
	repository "fortuna-express-web/pkg/repositories"
	"fortuna-express-web/pkg/web"
	handlers "fortuna-express-web/pkg/web/handlers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // Importar el driver de MySQL
)

// NewDatabase crea y retorna una nueva conexión a la base de datos MySQL
func NewDatabase() (*sql.DB, error) {
	// Obtiene las credenciales de la base de datos desde las variables de entorno

	// Define la cadena de conexión a la base de datos MySQL
	connStr := "root:password@tcp(localhost:3306)/fortuna"

	// Conectar a la base de datos MySQL
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	// Verificar si la conexión es exitosa
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	log.Println("Successfully connected to the MySQL database")

	// Crear las tablas si no existen
	err = createTables(db)
	if err != nil {
		return nil, fmt.Errorf("failed to create tables: %w", err)
	}

	return db, nil
}

// createTables crea las tablas necesarias en la base de datos MySQL
func createTables(db *sql.DB) error {
	// SQL para crear las tablas
	liquidationsQuery := `CREATE TABLE IF NOT EXISTS liquidations (
		id INT AUTO_INCREMENT PRIMARY KEY,
		departure VARCHAR(255) NOT NULL,
		arrival VARCHAR(255) NOT NULL,
		laundry DECIMAL(10, 2) NOT NULL,
		garage DECIMAL(10, 2) NOT NULL,
		guardianship DECIMAL(10, 2) NOT NULL,
		cover DECIMAL(10, 2) NOT NULL,
		sweeper DECIMAL(10, 2) NOT NULL,
		driver VARCHAR(255) NOT NULL,
		fuel DECIMAL(10, 2) NOT NULL,
		date DATETIME NOT NULL,
		freight DECIMAL(10, 2) NOT NULL,
		freight_liquid DECIMAL(10, 2) NOT NULL,
		detraction DECIMAL(10, 2) NOT NULL,
		gremission VARCHAR(255) NOT NULL,
		gtransport VARCHAR(255) NOT NULL,
		gtransport2 VARCHAR(255) NOT NULL,
		invoice VARCHAR(255) NOT NULL,
		driver_pay DECIMAL(10, 2) NOT NULL,
		drive_description TEXT NOT NULL,
		fuel_description TEXT NOT NULL,
		liquid_trip DECIMAL(10, 2) NOT NULL,
		truck VARCHAR(255) NOT NULL,
		expense_total DECIMAL(10, 2) NOT NULL,
		toll DECIMAL(10, 2) NOT NULL,
		gast_adition BOOLEAN NOT NULL
	);`

	additionsQuery := `CREATE TABLE IF NOT EXISTS aditions (
		id INT AUTO_INCREMENT PRIMARY KEY,
		description TEXT NOT NULL,
		price DECIMAL(10, 2) NOT NULL,
		liquidation_id INT NOT NULL,
		FOREIGN KEY (liquidation_id) REFERENCES liquidations(id)
		ON DELETE CASCADE
	);`

	// Ejecutar el script SQL para crear las tablas
	_, err := db.Exec(liquidationsQuery)
	if err != nil {
		return fmt.Errorf("failed to execute create liquidations table query: %w", err)
	}

	_, err = db.Exec(additionsQuery)
	if err != nil {
		return fmt.Errorf("failed to execute create additions table query: %w", err)
	}

	log.Println("Tables created or already exist")
	return nil
}

func main() {
	// Crear una instancia de Gin
	r := gin.Default()

	// Inicializar el logger
	logger := slog.Default()

	// Inicializar la base de datos MySQL
	database, err := NewDatabase()
	if err != nil {
		log.Fatal(err)
	}

	// Crear repositorios y pasar el logger y la base de datos
	liquidationRepo := repository.NewLiquidationRepository(logger, database)
	aditionRepo := repository.NewAditionRepository(logger, database)

	// Crear el caso de uso
	uc := uc.NewLiquidationUseCase(liquidationRepo, aditionRepo)

	// Crear el handler
	handler := handlers.NewLiquidationsHandler(logger, uc)

	// Configurar el router desde el paquete 'web'
	web.SetupRouter(r, handler)

	// Iniciar el servidor en el puerto 8080
	if err := r.Run(":8080"); err != nil {
		logger.Error("Could not start server", err)
	}
}
