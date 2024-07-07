package handlers

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"html/template"

	"github.com/akshay/bookings/internal/config"
	"github.com/akshay/bookings/internal/models"
	"github.com/akshay/bookings/internal/render"
	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/justinas/nosurf"
)

var functions = template.FuncMap{}

var app config.AppConfig
var session *scs.SessionManager
var pathToTemplates = "./../../templates"

func getRoutes() http.Handler {
	//what am i going to put in the session
	gob.Register(models.Reservation{})
	// change this to true when in production
	app.InProduction = false

	// set up the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := CreateTestTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	//we chanegs it to treue else it will create the temaptles on every request even though they are created in the render.go
	app.UseCache = true

	repo := NewRepo(&app)
	NewHandlers(repo)

	render.NewTemplates(&app)

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	//comment the no surf out because we dont need to check it as already done in my middleware
	// mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", Repo.Home)
	mux.Get("/generals", Repo.Generals)
	mux.Get("/search", Repo.Availability)
	//we create the get request forst because we want that to test it via get first
	// we alos crteated the new route taht jsut go to handlers and difplsays our struct
	mux.Post("/search-json", Repo.AvailabilityJSON)

	//to handle the psot repsone coming for the search route
	mux.Post("/search", Repo.PostAvailability)

	mux.Get("/contact", Repo.Contact)

	mux.Get("/makeReservation", Repo.Reservation)
	mux.Post("/makeReservation", Repo.PostReservation)
	mux.Get("/reservationSummary", Repo.ReservationSummary)

	mux.Get("/major", Repo.Majors)
	mux.Get("/about", Repo.About)
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux

}

// NoSurf is the csrf protection middleware
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad loads and saves session data for current request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
func CreateTestTemplateCache() (map[string]*template.Template, error) {

	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))
	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
