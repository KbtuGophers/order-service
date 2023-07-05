package repository

import (
	"github.com/KbtuGophers/order-service/internal/repository/postgres"
	"github.com/KbtuGophers/order-service/pkg/database"
)

type Configuration func(r *Repository) error

type Repository struct {
	postgres *database.Database
	Order    order.Repository
}

func New(configs ...Configuration) (r *Repository, err error) {
	// Create the repository
	r = &Repository{}

	// Apply all Configurations passed in
	for _, cfg := range configs {
		// Pass the repository into the configuration function
		if err = cfg(r); err != nil {
			return
		}
	}

	return
}

func (r *Repository) Close() {
	if r.postgres != nil {
		r.postgres.Client.Close()
	}
}

func WithPostgresStore(schema, dataSourceName string) Configuration {
	return func(r *Repository) (err error) {
		r.postgres, err = database.NewDatabase(schema, dataSourceName)
		if err != nil {
			return
		}

		if err = r.postgres.Migrate(); err != nil && err.Error() != "no change" {
			return
		}
		err = nil

		r.Order = postgres.NewOrderRepository(r.postgres.Client)

		return
	}
}
