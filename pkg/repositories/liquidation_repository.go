package repository

import (
	"context"
	"database/sql"
	"fmt"
	"fortuna-express-web/pkg/domain/entities"
	"fortuna-express-web/pkg/interfaces"
	"log/slog"
	"time"
)

type liquidationRepository struct {
	log *slog.Logger
	db  *sql.DB
}

func NewLiquidationRepository(log *slog.Logger, database *sql.DB) interfaces.LiquidationRepository {
	return &liquidationRepository{log, database}
}

func (r *liquidationRepository) New(liquidation *entities.Liquidation) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Primero, insertamos los datos
	query := `
		INSERT INTO liquidations (
			departure, arrival, laundry, garage, guardianship, cover, sweeper, driver, fuel, 
			date, freight, freight_liquid, detraction, gremission, gtransport, gtransport2, 
			invoice, driver_pay, drive_description, fuel_description, liquid_trip, truck, 
			expense_total, toll, gast_adition
		) VALUES (
			?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
		)`
	_, err := r.db.ExecContext(ctx, query,
		liquidation.Departure, liquidation.Arrival, liquidation.Laundry, liquidation.Garage,
		liquidation.Guardianship, liquidation.Cover, liquidation.Sweeper, liquidation.Driver,
		liquidation.Fuel, liquidation.Date, liquidation.Freight, liquidation.FreightLiquid,
		liquidation.Detraction, liquidation.Gremission, liquidation.Gtransport, liquidation.Gtransport2,
		liquidation.Invoice, liquidation.DriverPay, liquidation.DriveDescription, liquidation.FuelDescription,
		liquidation.LiquidTrip, liquidation.Truck, liquidation.ExpenseTotal, liquidation.Toll,
		liquidation.GastAdition,
	)

	if err != nil {
		return 0, fmt.Errorf("failed to add liquidation: %w", err)
	}

	// Luego, obtenemos el último ID insertado
	var id int
	err = r.db.QueryRowContext(ctx, "SELECT LAST_INSERT_ID()").Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert id: %w", err)
	}

	return id, nil
}

func (r *liquidationRepository) List() ([]*entities.Liquidation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "SELECT * FROM liquidations"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to list liquidations: %w", err)
	}
	defer rows.Close()

	var liquidations []*entities.Liquidation
	for rows.Next() {
		var l entities.Liquidation
		var dateRaw []byte // Cambiamos a []byte para manejar el valor crudo

		err := rows.Scan(
			&l.ID, &l.Departure, &l.Arrival, &l.Laundry, &l.Garage, &l.Guardianship,
			&l.Cover, &l.Sweeper, &l.Driver, &l.Fuel, &dateRaw, &l.Freight, &l.FreightLiquid,
			&l.Detraction, &l.Gremission, &l.Gtransport, &l.Gtransport2, &l.Invoice,
			&l.DriverPay, &l.DriveDescription, &l.FuelDescription, &l.LiquidTrip,
			&l.Truck, &l.ExpenseTotal, &l.Toll, &l.GastAdition,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan liquidation: %w", err)
		}

		// Si el valor de dateRaw no es nulo, lo convertimos a time.Time
		if len(dateRaw) > 0 {
			parsedTime, err := time.Parse("2006-01-02 15:04:05", string(dateRaw))
			if err != nil {
				return nil, fmt.Errorf("failed to parse date: %w", err)
			}
			l.Date = &parsedTime
		}

		liquidations = append(liquidations, &l)
	}

	return liquidations, nil
}

func (r *liquidationRepository) Get(id int) (*entities.Liquidation, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "SELECT * FROM liquidations WHERE id = ?"
	var l entities.Liquidation
	var dateRaw []byte // Cambiamos a []byte para manejar el valor crudo
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&l.ID, &l.Departure, &l.Arrival, &l.Laundry, &l.Garage, &l.Guardianship,
		&l.Cover, &l.Sweeper, &l.Driver, &l.Fuel, &dateRaw, &l.Freight, &l.FreightLiquid,
		&l.Detraction, &l.Gremission, &l.Gtransport, &l.Gtransport2, &l.Invoice,
		&l.DriverPay, &l.DriveDescription, &l.FuelDescription, &l.LiquidTrip,
		&l.Truck, &l.ExpenseTotal, &l.Toll, &l.GastAdition,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to get liquidation with id %d: %w", id, err)
	}

	// Si el valor de dateRaw no es nulo, lo convertimos a time.Time
	if len(dateRaw) > 0 {
		parsedTime, err := time.Parse("2006-01-02 15:04:05", string(dateRaw))
		if err != nil {
			return nil, fmt.Errorf("failed to parse date: %w", err)
		}
		l.Date = &parsedTime
	}

	return &l, nil
}

func (r *liquidationRepository) Update(liquidation *entities.Liquidation) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		UPDATE liquidations
		SET departure = ?, arrival = ?, laundry = ?, garage = ?, guardianship = ?, cover = ?, 
			sweeper = ?, driver = ?, fuel = ?, date = ?, freight = ?, freight_liquid = ?, 
			detraction = ?, gremission = ?, gtransport = ?, gtransport2 = ?, invoice = ?, 
			driver_pay = ?, drive_description = ?, fuel_description = ?, liquid_trip = ?, 
			truck = ?, expense_total = ?, toll = ?, gast_adition = ?
		WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query,
		liquidation.Departure, liquidation.Arrival, liquidation.Laundry, liquidation.Garage,
		liquidation.Guardianship, liquidation.Cover, liquidation.Sweeper, liquidation.Driver,
		liquidation.Fuel, liquidation.Date, liquidation.Freight, liquidation.FreightLiquid,
		liquidation.Detraction, liquidation.Gremission, liquidation.Gtransport, liquidation.Gtransport2,
		liquidation.Invoice, liquidation.DriverPay, liquidation.DriveDescription, liquidation.FuelDescription,
		liquidation.LiquidTrip, liquidation.Truck, liquidation.ExpenseTotal, liquidation.Toll,
		liquidation.GastAdition, liquidation.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update liquidation with id %d: %w", liquidation.ID, err)
	}

	return nil
}

func (r *liquidationRepository) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := "DELETE FROM liquidations WHERE id = ?"
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete liquidation with id %d: %w", id, err)
	}

	return nil
}
