package dbrepo

import (
	"database/sql"

	"github.com/akshay/bookings/internal/config"
	"github.com/akshay/bookings/internal/repository"
)

//well this would save our time if we are suign a separate db in future and we can easily change that then

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}
