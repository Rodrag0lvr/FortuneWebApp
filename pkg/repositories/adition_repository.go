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

type aditionRepository struct {
	log *slog.Logger
	db  *sql.DB
}

func NewAditionRepository(log *slog.Logger, database *sql.DB) interfaces.AditionRepository {
	return &aditionRepository{log, database}
}

// New inserta una nueva adición y retorna el ID generado
func (r *aditionRepository) New(adition *entities.Adition) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Consulta de inserción para MySQL (usando ? como placeholder)
	query := `
		INSERT INTO aditions (description, price, liquidation_id)
		VALUES (?, ?, ?)`

	// Ejecutar la inserción
	result, err := r.db.ExecContext(ctx, query,
		adition.Description, adition.Price, adition.LiquidationId,
	)
	if err != nil {
		return 0, fmt.Errorf("failed to add adition: %w", err)
	}

	// Obtener el ID del último registro insertado
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert id: %w", err)
	}

	return int(id), nil
}

func (r *aditionRepository) List() ([]*entities.Adition, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `SELECT id, description, price, liquidation_id FROM aditions`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch aditions: %w", err)
	}
	defer rows.Close()

	var aditions []*entities.Adition
	for rows.Next() {
		adition := &entities.Adition{}
		err := rows.Scan(&adition.ID, &adition.Description, &adition.Price, &adition.LiquidationId)
		if err != nil {
			return nil, fmt.Errorf("failed to scan adition: %w", err)
		}
		aditions = append(aditions, adition)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return aditions, nil
}

func (r *aditionRepository) Get(id int) ([]*entities.Adition, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `SELECT id, description, price, liquidation_id FROM aditions WHERE id = ?`
	rows, err := r.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch aditions: %w", err)
	}
	defer rows.Close()

	var aditions []*entities.Adition
	for rows.Next() {
		adition := &entities.Adition{}
		if err := rows.Scan(&adition.ID, &adition.Description, &adition.Price, &adition.LiquidationId); err != nil {
			return nil, fmt.Errorf("failed to scan adition: %w", err)
		}
		aditions = append(aditions, adition)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	if len(aditions) == 0 {
		return nil, fmt.Errorf("no aditions found with id: %d", id)
	}

	return aditions, nil
}

func (r *aditionRepository) Update(adition *entities.Adition) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `
		UPDATE aditions 
		SET description = ?, price = ?, liquidation_id = ?
		WHERE id = ?`
	result, err := r.db.ExecContext(ctx, query,
		adition.Description, adition.Price, adition.LiquidationId, adition.ID,
	)

	if err != nil {
		return fmt.Errorf("failed to update adition: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no adition updated with id: %d", adition.ID)
	}

	return nil
}

func (r *aditionRepository) Delete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	query := `DELETE FROM aditions WHERE id = ?`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete adition: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to check rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no adition deleted with id: %d", id)
	}

	return nil
}
