package main

import (
	"net/http"

	"github.com/akshay/bookings/internal/config"
	"github.com/akshay/bookings/internal/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/generals", handlers.Repo.Generals)
	mux.Get("/search", handlers.Repo.Availability)
	//we create the get request forst because we want that to test it via get first
	// we alos crteated the new route taht jsut go to handlers and difplsays our struct
	mux.Post("/search-availability-json", handlers.Repo.AvailabilityJSON)
	mux.Get("/choose-room/{id}", handlers.Repo.ChooseRoom)
	mux.Get("/book-room", handlers.Repo.BookRoom)

	//to handle the psot repsone coming for the search route
	mux.Post("/search", handlers.Repo.PostAvailability)

	mux.Get("/contact", handlers.Repo.Contact)

	mux.Get("/makeReservation", handlers.Repo.Reservation)
	mux.Post("/makeReservation", handlers.Repo.PostReservation)
	mux.Get("/reservationSummary", handlers.Repo.ReservationSummary)

	mux.Get("/major", handlers.Repo.Majors)
	mux.Get("/about", handlers.Repo.About)
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}
