package persistence

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
)

// signed file to generate mock.
//go:generate mockgen -source ./persistence.go -destination mocks/persistence.mock.go -package mocks

// Repository persistence methods.
type Repository interface {
	GetTime(ctx context.Context) (*time.Time, error)
}

// RepositoryPg is a postgres implementation of Repository.
type RepositoryPg struct {
	dbInstance *sqlx.DB
}

// This line ensures that the RepositoryPg struct implements the Repository interface.
var _ Repository = &RepositoryPg{}

func NewRepository(db *sql.DB) (*RepositoryPg, error) {
	pgDB := sqlx.NewDb(db, "postgres")

	return &RepositoryPg{
		dbInstance: pgDB,
	}, nil
}

// GetTime implements Repository.
func (r *RepositoryPg) GetTime(ctx context.Context) (*time.Time, error) {
	timeStruct := time.Time{}

	err := r.dbInstance.GetContext(ctx, &timeStruct, "SELECT NOW() as t")
	if err != nil {
		return nil, err
	}

	return &timeStruct, nil
}
