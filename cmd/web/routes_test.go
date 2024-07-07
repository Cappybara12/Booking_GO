package main

import (
	"fmt"
	"testing"

	"github.com/akshay/bookings/internal/config"
	"github.com/go-chi/chi"
)
//so while creating tests for a particluar file we will checkout what kind of inputs it is taking 
//then depedning on that inouts we will test them 
func TestRoutes(t *testing.T) {
	//decalred the inout that our fucntion in routes si getting 
	var app config.AppConfig
	//we will test the mux for the routes 
	mux := routes(&app)
	switch v := mux.(type) {
	case *chi.Mux:
		//do nothing test passed

	default:
		t.Error(fmt.Sprintf("type is not *chi.Mux,type is %T", v))
	}
}
