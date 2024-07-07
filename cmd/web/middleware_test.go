package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSurve(t *testing.T) {
	var myH Myhandler
	h := NoSurf(&myH)
	//we wjsut look for the type and it has to be correct type  to pss our test 
	switch v := h.(type) {
	case http.Handler:
		//do nothing
	default:
		t.Error(fmt.Sprintf("type is not http.handler, but is %T", v))
	}
}
func TestNoSurf(t *testing.T) {
	var myH Myhandler
	h := SessionLoad(&myH)

	switch v := h.(type) {
	case http.Handler:
		//do nothing
	default:
		t.Error(fmt.Sprintf("type is not http.handler, but is %T", v))
	}
}
