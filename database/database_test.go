package database_test

import (
	"testing"

	"github.com/LOO2/business-remote-management-api/database"
	"github.com/gin-gonic/gin"
)

func TestStartDB(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Database Connecting", func(t *testing.T) {
		db_error := database.StartDB()
		if db_error != nil {
			t.Fatalf("error database connect, got :" + db_error.Error())
		}
	})

	t.Run("Instantiating Database", func(t *testing.T) {
		db_error := database.GetDatabase().Error
		if db_error != nil {
			t.Fatalf("error database instante, got :" + db_error.Error())
		}
	})

	t.Run("Closing Database Connection", func(t *testing.T) {
		db_error := database.CloseConn()

		if db_error != nil {
			t.Fatalf("error database clone connection, got :" + db_error.Error())
		}
	})
}
