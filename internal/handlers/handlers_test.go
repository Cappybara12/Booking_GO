package handlers

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/akshay/bookings/internal/models"
)

type postData struct {
	key   string
	value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	// {"home", "/", "GET", []postData{}, http.StatusOK},
	// {"about", "/about", "GET", []postData{}, http.StatusOK},
	// {"gq", "/generals", "GET", []postData{}, http.StatusOK},
	// {"ms", "/major", "GET", []postData{}, http.StatusOK},
	// {"sa", "/search", "GET", []postData{}, http.StatusOK},
	// {"contact", "/contact", "GET", []postData{}, http.StatusOK},
	// {"mr", "/makeReservation", "GET", []postData{}, http.StatusOK},
	// {"post-search-avail", "/search", "POST", []postData{
	// 	{key: "start", value: "2020-01-01"},
	// 	{key: "end", value: "2020-01-02"},
	// }, http.StatusOK},
	// {"post-search-avail-json", "/search", "POST", []postData{
	// 	{key: "start", value: "2020-01-01"},
	// 	{key: "end", value: "2020-01-02"},
	// }, http.StatusOK},
	// {"mrp", "/makeReservation", "POST", []postData{
	// 	{key: "first_name", value: "John"},
	// 	{key: "last_name", value: "Smith"},
	// 	{key: "email", value: "me@here.com"},
	// 	{key: "phone", value: "9212122"},
	// }, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		t.Run(e.name, func(t *testing.T) {
			if e.method == "GET" {
				resp, err := ts.Client().Get(ts.URL + e.url)
				if err != nil {
					t.Errorf("Error making GET request to %s: %v", e.url, err)
					return
				}
				defer resp.Body.Close()

				if resp.StatusCode != e.expectedStatusCode {
					t.Errorf("For %s, expected status %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
				}
			} else if e.method == "POST" {
				values := url.Values{}
				for _, p := range e.params {
					values.Add(p.key, p.value)
				}

				resp, err := ts.Client().PostForm(ts.URL+e.url, values)
				if err != nil {
					t.Errorf("Error making POST request to %s: %v", e.url, err)
					return
				}
				defer resp.Body.Close()

				if resp.StatusCode != e.expectedStatusCode {
					t.Errorf("For %s, expected status %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
				}
			}
		})
	}
}

func TestRepository_Reservation(t *testing.T) {
	reservation := models.Reservation{
		RoomID: 1,
		Room: models.Room{
			ID:       1,
			RoomName: "General's Quarters",
		},
	}

	req, _ := http.NewRequest("GET", "/makeReservation", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)
	rr := httptest.NewRecorder()
	session.Put(ctx, "reservation", reservation)
	handler := http.HandlerFunc(Repo.Reservation)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Reservation handler returned wrong response code: got %v want %v", rr.Code, http.StatusOK)
	}
}

func getCtx(req *http.Request) context.Context {
	ctx, err := session.Load(req.Context(), req.Header.Get("X-Session"))
	if err != nil {
		log.Println(err)
	}
	return ctx
}
