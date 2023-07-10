package postgres

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type ProcessRepository struct {
	db *sqlx.DB
}

func (p *ProcessRepository) GetStatus(ctx context.Context, id string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (p *ProcessRepository) Checkout(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (p *ProcessRepository) Cancel(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func NewProcessRepository(db *sqlx.DB) *ProcessRepository {
	return &ProcessRepository{db: db}
}
