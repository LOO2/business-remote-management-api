package controllers_test

import (
	"net/http"
	"testing"

	"github.com/steinfletcher/apitest"
)

func TestShowAllRevenues(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		msg := `{
				"debit"  : 1.0,
				"credit" : 2.0
			}`
		_, _ = w.Write([]byte(msg))
		w.WriteHeader(http.StatusOK)
	}

	apitest.New().
		HandlerFunc(handler).
		Get("/revenue").
		Expect(t).
		Body(`{
			"debit"  : 1.0,
			"credit" : 2.0
		}`).
		Status(http.StatusOK).
		End()
}
