package delivery_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/LOO2/business-remote-management-api/delivery"
	"github.com/gin-gonic/gin"
)

func Test_ShowAllRevenues(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		router := gin.Default()

		router.GET("api/revenue/", delivery.ShowAllRevenues)

		ts := httptest.NewServer(router)

		ts.URL = "http://localhost:8080"

		defer ts.Close()

		// Make a request to our server with the {base url}/api/revenue
		resp, err := http.Get(fmt.Sprintf("%s/api/revenue", ts.URL))

		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if resp.StatusCode != 200 {
			t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
		}

		val, ok := resp.Header["Content-Type"]

		// Assert that the "content-type" header is actually set
		if !ok {
			t.Fatalf("Expected Content-Type header to be set")
		}

		// Assert that it was set as expected
		if val[0] != "application/json; charset=utf-8" {
			t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", val[0])
		}
	})
}
