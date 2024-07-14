package dbrepo

import (
	"database/sql"

	"github.com/akshay/bookings/internal/config"
	"github.com/akshay/bookings/internal/repository"
)

// postgresDBRepo holds the DB and AppConfig
type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

// NewPostgresRepo creates a new repository
func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}

// AllUsers returns true for now
func (m *postgresDBRepo) AllUsers() bool {
	return true
}
