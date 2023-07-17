package postgres

import (
	"context"
	"github.com/KbtuGophers/order-service/internal/domain/processes"
	"github.com/jmoiron/sqlx"
)

type ProcessRepository struct {
	db *sqlx.DB
}

func (p *ProcessRepository) GetStatus(ctx context.Context, orderID string) (data processes.Entity, err error) {
	query := `SELECT * FROM processes WHERE order_id = $1`
	args := []any{orderID}

	err = p.db.GetContext(ctx, &data, query, args...)
	if err != nil {
		return
	}

	return
}

func (p *ProcessRepository) PendingToProcessing(ctx context.Context, orderID, stage, status string) error {
	query := `UPDATE processes SET state='processing', stage=$1, order_status=$2 WHERE order_id = $3`
	args := []any{stage, status, orderID}

	_, err := p.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProcessRepository) ProcessingToCompleted(ctx context.Context, orderID, stage, status string) error {
	query := `UPDATE processes SET state='completed', stage=$1, order_status=$2 WHERE order_id = $3`
	args := []any{stage, status, orderID}

	_, err := p.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}
func (p *ProcessRepository) ProcessingToCanceled(ctx context.Context, orderID, stage, status string) error {
	query := `UPDATE processes SET state='canceled', stage=$1, order_status=$2 WHERE order_id = $3`
	args := []any{stage, status, orderID}

	_, err := p.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil

}

func NewProcessRepository(db *sqlx.DB) *ProcessRepository {
	return &ProcessRepository{db: db}
}
