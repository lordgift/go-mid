package main

import (
	"github.com/stretchr/testify/assert"
	"go-mid/services"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClientPing(t *testing.T) {
	s := services.CreateService()
	r := services.CreateRouter(s)
	//r.Run()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, w.Body)
}
