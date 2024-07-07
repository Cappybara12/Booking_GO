package main

import (
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {

	os.Exit(m.Run())
}

type Myhandler struct{}

func (mh *Myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
