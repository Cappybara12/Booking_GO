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

type testDBRepo struct {
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

func NewTestingRepo(a *config.AppConfig) repository.DatabaseRepo {
	return &testDBRepo{
		App: a,
	}
}

// AllUsers returns true for now
func (m *postgresDBRepo) AllUsers() bool {
	return true
}

// AllUsers for testDBRepo - implement the method for testing
func (m *testDBRepo) AllUsers() bool {
	// For testing purposes, you can return a predetermined value
	return true
}
