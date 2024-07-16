package handlers

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
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
	expectedStatusCode int
}{
	{"home", "/", "GET", http.StatusOK},
	{"about", "/about", "GET", http.StatusOK},
	{"gq", "/generals", "GET", http.StatusOK},
	{"ms", "/major", "GET", http.StatusOK},
	{"sa", "/search", "GET", http.StatusOK},
	{"contact", "/contact", "GET", http.StatusOK},
}

func TestRepository_PostReservation(t *testing.T) {
	tests := []struct {
		name               string
		reqBody            string
		expectedStatusCode int
	}{
		{
			"valid request",
			"start_date=2050-01-01&end_date=2050-01-02&first_name=John&last_name=Smith&email=john@smith.com&phone=123456789&room_id=1",
			http.StatusSeeOther,
		},
		{
			"missing post body",
			"",
			http.StatusTemporaryRedirect,
		},
		{
			"invalid start date",
			"start_date=invalid&end_date=2050-01-02&first_name=John&last_name=Smith&email=john@smith.com&phone=123456789&room_id=1",
			http.StatusTemporaryRedirect,
		},
		{
			"invalid end date",
			"start_date=2050-01-01&end_date=invalid&first_name=John&last_name=Smith&email=john@smith.com&phone=123456789&room_id=1",
			http.StatusTemporaryRedirect,
		},
		{
			"invalid room ID",
			"start_date=2050-01-01&end_date=2050-01-02&first_name=John&last_name=Smith&email=john@smith.com&phone=123456789&room_id=invalid",
			http.StatusTemporaryRedirect,
		},
	}

	for _, test := range tests {
		req, _ := http.NewRequest("POST", "/makeReservation", strings.NewReader(test.reqBody))
		ctx := getCtx(req)
		req = req.WithContext(ctx)
		req.Header.Set("Content-type", "application/x-www-form-urlencoded")
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(Repo.PostReservation)
		handler.ServeHTTP(rr, req)

		if rr.Code != test.expectedStatusCode {
			t.Errorf("%s: Post Reservation handler returned wrong response code: got %v want %v", test.name, rr.Code, test.expectedStatusCode)
		}
	}
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		resp, err := ts.Client().Get(ts.URL + e.url)
		if err != nil {
			t.Errorf("Error making GET request to %s: %v", e.url, err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != e.expectedStatusCode {
			t.Errorf("For %s, expected status %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
		}
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

	// Test case where reservation won't be in session
	req, _ = http.NewRequest("GET", "/makeReservation", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("Reservation handler returned wrong response code for missing reservation: got %v want %v", rr.Code, http.StatusTemporaryRedirect)
	}

	// Test with non-existent room
	reservation.RoomID = 100
	session.Put(ctx, "reservation", reservation)
	req, _ = http.NewRequest("GET", "/makeReservation", nil)
	ctx = getCtx(req)
	req = req.WithContext(ctx)
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusTemporaryRedirect {
		t.Errorf("Reservation handler returned wrong response code for non-existent room: got %v want %v", rr.Code, http.StatusTemporaryRedirect)
	}
}

func getCtx(req *http.Request) context.Context {
	ctx, err := session.Load(req.Context(), req.Header.Get("X-Session"))
	if err != nil {
		log.Println(err)
	}
	return ctx
}
