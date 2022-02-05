package controllers

import (
	"testing"

	"github.com/LOO2/business-remote-management-api/database"
	models "github.com/LOO2/business-remote-management-api/domain"
)

func TestShowAllRevenues(t *testing.T) {
	database.StartDB()
	db := database.GetDatabase()

	var p []models.Revenue
	err := db.Find(&p).Error

	if err != nil {
		t.Errorf("cannot find revenue: " + err.Error())
	}
}
