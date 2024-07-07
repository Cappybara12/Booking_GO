package render

import (
	"net/http"
	"testing"

	"github.com/akshay/bookings/internal/models"
	"github.com/alexedwards/scs/v2"
)

func TestAddDefaultData(t *testing.T) {
	var td models.TemplateData
	r, err := getSession()
	if err != nil {
		t.Error(err)
	}
	session.Put(r.Context(), "flash", "123")
	result := AddDefaultData(&td, r)
	if result.Flash != "123" { // Changed condition to check for inequality
		t.Error("flash value 123 not found in session")
	}
}

func getSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/some-url", nil)
	if err != nil {
		return nil, err
	}
	ctx := r.Context()
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))
	r = r.WithContext(ctx)
	return r, nil
}

func init() {
	session = scs.New()
}
