package helpers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/akshay/bookings/internal/config"
)

var app *config.AppConfig

// sets up add config for helpers
func NewHelpers(a *config.AppConfig) {
	app = a
}

// we can  have two kind of errors hta are either the client error or the server error cvering both
func ClientError(w http.ResponseWriter, status int) {
	app.InfoLog.Println("client error with status", status)
	http.Error(w, http.StatusText(status), status)
}
func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w,http.StatusText(http.StatusInternalServerError),http.StatusInternalServerError)
}
